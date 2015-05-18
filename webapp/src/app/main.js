// just so my syntastic complains less
var React = React || {},
    _ = _ || {},
    d3 = d3 || {};

var data = new pvGraph(JSON.parse(document.getElementById("pipe-graph").innerHTML));

var Viz = React.createClass({
    displayName: "pipeviz-graph",
    getInitialState: function() {
        return {
            force: d3.layout.force()
            .charge(-4000)
            .chargeDistance(250)
            .size([this.props.width, this.props.height])
            .linkStrength(function(link) {
                if (link.source.Typ() == "environment") {
                    return 0.5;
                }
                if (link.source.Typ() === "anchor" || link.target.Typ() === "anchor") {
                    return 1;
                }
                return 0.3;
            })
            .linkDistance(function(link) {
                if (link.source.Typ() === "anchor" || link.target.Typ() === "anchor") {
                    return 25;
                }
                return 250;
            })
        };
    },
    getDefaultProps: function() {
        return {
            vertices: {},
            edges: {},
        };
    },
    render: function() {
        return React.DOM.svg({
            className: "pipeviz",
            viewBox: "0 0 " + this.props.width + " " + this.props.height
        });
    },
    componentDidUpdate: function() {
        this.state.force.nodes(this.props.nodes);
        this.state.force.links(this.props.links);

        return this.graphRender(this.getDOMNode(), this.state, this.props);
    },
    graphRender: function(el, state, props) {
        var link = d3.select(el).selectAll('.link')
            .data(props.links, function(d) { return d.source.objid() + '-' + d.target.objid(); }),
        node = d3.select(el).selectAll('.node')
            .data(props.nodes, function(d) { return d.id; });

        link.enter().append('line')
        .attr('class', function(d) {
            if (d.source.Typ() === "anchor" || d.target.Typ() === "anchor") {
                return 'link anchor';
            }
            if (_.has(d, 'path')) {
                return 'link link-commit';
            }
            return 'link';
        })
        .style('stroke-width', function(d) {
            return (d.path && d.path.length > 0) ? 1.5 * Math.sqrt(d.path.length) : 1;
        });

        var nodeg = node.enter().append('g')
        .attr('class', function(d) {
            return 'node ' + d.vType();
        });

        nodeg.append('circle')
        .attr('x', 0)
        .attr('y', 0)
        .attr('r', function(d) {
            if (d.Typ() === "logic-state") {
                return 45;
            }
            if (d.Typ() === "commit") {
                return 3;
            }
            if (d.Typ() === "anchor") {
                return 0;
            }
        })
        .on('click', props.target);

        nodeg.append('image')
        .attr('class', 'provider-logo')
        .attr('height', 22)
        .attr('width', 22)
        .attr('y', '-37')
        .attr('x', '-10')
        .attr('xlink:href', function(d) {
            if (d.Typ() === "logic-state" && _.has(d.vertex, "provider")) {
                return "assets/" + d.vertex.provider + ".svg";
            }
        });

        var nodetext = nodeg.append('text');
        nodetext.append('tspan').text(function(d) { return d.name(); });
        nodetext.append('tspan').text(function(d) {
            // FIXME omg terrible
            if (d.Typ() !== "logic-state") {
                return '';
            }

            return d.ref().id.commit.slice(0, 7);
        })
        .attr('dy', "1.4em")
        .attr('x', 0)
        .attr('class', function(d) {
            if (d.Typ() !== "logic-state") {
                return;
            }

            var output = 'commit-subtext',
            commit = d.ref().id.commit;

            if (_.has(props.commitMeta, commit) &&
                _.has(props.commitMeta[commit], 'testState')) {
                output += ' commit-' + props.commitMeta[commit].testState;
            }

            return output;
        });

        node.exit().remove();
        link.exit().remove();

        state.force.on('tick', function() {
            link.attr("x1", function(d) { return d.source.x; })
            .attr("y1", function(d) { return d.source.y; })
            .attr("x2", function(d) { return d.target.x; })
            .attr("y2", function(d) { return d.target.y; });

            node.attr('transform', function(d) { return 'translate(' + d.x + ',' + d.y + ')'; });
        });

        state.force.start();
        return false;
    }
});

var VizPrep = React.createClass({
    getInitialState: function() {
        return {
            anchorL: new Anchor("L", 0, this.props.height/2),
            anchorR: new Anchor("R", this.props.width, this.props.height/2),
        };
    },
    extractVizGraph: function(repo) {
        var g = this.props.graph.commitGraph(),
        links = [],
        cmp = this,
        members = {},
        nodes = _.filter(_.map(this.props.graph.verticesWithType("logic-state"), function(d) { return _.create(pvVertex.prototype, d); }), function(v) {
            var vedges = _.filter(_.map(v.outEdges, function(edgeId) { return cmp.props.graph.get(edgeId); }), isType("version"));
            if (vedges.length === 0) {
                return false;
            }

            if (cmp.props.graph.get(vedges[0].target).prop("repository").value === repo) {
                if (!_.has(members, vedges[0].target)) {
                    members[vedges[0].target] = [];
                }
                // have to copy objects for d3's prop intrusion
                members[vedges[0].target].push(v);
                return true;
            }
            return false;
        });

        _.each(this.props.graph.verticesWithType("vcs-label"), function(l) {
            var vedges = _.filter(_.map(l.outEdges, function(edgeId) { return cmp.props.graph.get(edgeId); }), isType("version"));

            if (cmp.props.graph.get(vedges[0].target).prop("repository").value === repo) {
                if (!_.has(members, vedges[0].target)) {
                    members[vedges[0].target] = [];
                }
                members[vedges[0].target].push(_.create(pvVertex.prototype, l));
            }
        });

        // walker to discover 'joints' - otherwise uninteresting commits with
        // multiple descendents that will necessitate creating a node later. also
        // identifies the 'interesting' sources and sinks (apps) that should actually
        // be used, rather than sinks/sources of the commit graph on its own.
        var visited = [], // "black" list - vertices that have been visited
        deepestApp, // var to hold the deepest app we've visited (for sink tracking)
        npath = [], // the current path, but only with logic states
        isources = [], // interesting sources
        isinks = [], // interesting sinks
        path = [], // the current path of interstitial commits
        sinkmap = {}, // map that tracks eventual sink target; prevents duplicate traversal
        prepwalk = function(v) {
            // DAG, so skip grey/back-edge

            var has_app = false;
            // check if this commit has an app. don't check if white b/c sink calc necessitates we
            // recheck this every time through.
            if (_.has(members, v) && _.filter(members[v], isType("logic-state").length !== 0)) {
                if (npath.length === 0) {
                    // nothing in the npath, this is a source
                    isources.push(v);
                }
                // so we know to pop npath stack and check if source later
                has_app = true;
                npath.push(v);
                // track that this is now our deepest app
                deepestApp = v;
            }
            // if black, and nothing in members about the commit, it's a joint
            if (visited.indexOf(v) !== -1 && !_.has(members, v)) {
                members[v] = [_.create(pvVertex.prototype, cmp.props.graph.get(v))];
                //return;
            }

            path.push(v);
            g.successors(v).map(function(s) {
                prepwalk(s);
            });
            visited.push(v);
            path.pop();

            if (has_app) {
                npath.pop();
                // if deepestApp is self, it's because we found no apps further down the dag,
                // so this is effectively a sink
                if (deepestApp === v) {
                    isinks.push(v);
                    // write this vertex into the sinkmap for all current paths
                    _.merge(sinkmap, _.zipObject(_.zip(path, _.fill(v, path.length))));
                }
            }
        };

        // now traverse depth-first to figure out the overlaid edge structure
        var labels = [], // the set of non-app-coinciding labels and their relative positions
        lpnodes = {}, // the set of non-app-coinciding labels that have found their origin, but not their target
        v, // vertex (commit) currently being visited
        walk = function(v) { // main depth-first walker
            // git guarantees commit graph is acyclic, thus safe to skip grey/back-edge

            var pop_npath = false;
            // grab head of node path from stack
            var from = npath.length === 0 ? [] : npath[npath.length - 1];

            if (visited.indexOf(v) !== -1) {
                // Commit vertex is black/visited; create links, update lpnodes, and return.
                _.each(members[v], function(tgt) {
                    from.map(function(src) {
                        // slice guarantees path array is copied
                        links.push({ source: src.obj, target: tgt.obj, path: path.slice(0) });
                    });
                });

                // process all label nodes we have waiting around
                _.forOwn(lpnodes, function(id, llen) {
                    // need to push the actual objects on so the viz can cheat and track x/y props
                    labels.push({id: id, l: from[0].obj, r: ls[0].obj, pos: llen / (path.length+1)});
                });

                // zero out lpnodes set
                lpnodes = {};
                // zero out interstitial commit path tracker
                path = [];
                return;
            }

                // see if this is a commit that's interesting (has an app instance or label)
                if (_.has(members, v)) {
                    // different behavior depending on whether we're finding commit or (app and/or label)
                    // app handling is identical to commit handling, so we reuse
                    var ls = _.union(_.filter(members[v], isType("logic-state")),
                                     _.filter(members[v], isType("commit")));
                    var lbls = _.filter(members[v], isType("vcs-label"));
                    if (ls.length !== 0) {
                        // has at least one app, or is a commit joint. create link from last thing to this
                        _.each(ls, function(tgt) {
                            from.map(function(src) {
                                links.push({ source: src, target: tgt, path: path.slice(0) });
                                if (tgt.Typ() === "commit") {
                                    // push the joint onto the node list
                                    nodes.push(tgt);
                                }
                            });
                        });

                        // process all label nodes we have waiting around
                        _.forOwn(lpnodes, function(id, llen) {
                            // need to push the actual objects on so the viz can cheat and track x/y props
                            labels.push({id: id, l: from[0], r: ls[0], pos: llen / (path.length+1)});
                        });

                        // zero out lpnodes set
                        lpnodes = {};

                        if (ls.length !== members[v].length) {
                            // there are also some labels directly on the app commit, add them
                            _.each(lbls, function(lbl) {
                                labels.push({id: lbl.id, l: ls[0], r: ls[0], pos: 0});
                            });
                        }

                        // zero out the interstitial commit path tracker
                        path = [];
                        // push newly-found app(s) onto our npath, it's the new 'from'
                        npath.push(ls);
                        // correspondingly, indicate to pop the npath when exiting
                        pop_npath = true;
                    } else {
                        // for each label found, record the length of the path walked so far,
                        // plus one for the commit we're currently on
                        _.each(lbls, function(lbl) {
                            lpnodes[lbl.id] = path.length+1;
                        });

                        // still not a real node point though, so need to push the commit onto the path
                        path.push(v);
                    }
                }
                else {
                    // not a node point and not self - push commit onto path
                    path.push(v);
                }

            // recursive call, the crux of this depth-first traversal. but we
            // skip it if the first search proved this to be a sink
            if (isinks.indexOf(v) !== -1) {
                g.successors(v).map(function(s) {
                    walk(s);
                });
            }

            // Mark commit black/visited...but only if it's a member-associated
            // commit. This is OK to do because we already traversed once and
            // found all the joints, and those are guaranteed to be members
            if (_.has(members, v)) {
                visited.push(v);
            }

            if (pop_npath) {
                npath.pop();
            }
            // pop current commit off the visit path
            path.pop();

            // decrement lpnodes values, and filter out any that reach 0;
            // they're not between any apps and thus won't be displayed.
            // this SHOULD be unnecessary b/c we preselected good sinks and
            // sources, but just in case.
            lpnodes = _(lpnodes)
                .mapValues(function(v) { return v-1; })
                .filter(function(v) { return v > 0; }) // > 1?
                .value();
        };

        var stack = _.filter(g.sources(), function(d) {
            return cmp.props.graph.get(d).prop("repository").value === repo;
        });
        // DF walk, working from source commit members
        while (stack.length !== 0) {
            v = stack.pop();
            prepwalk(v);
        }

        // traversal pattern almost guarantees duplicate sinks
        isinks = _.uniq(isinks)

        // put the source anchor links in
        _.each(isources, function(commit) {
            _.each(_.filter(members[commit], isType("logic-state")), function(member) {
                links.push({ source: cmp.state.anchorL, target: member });
            });
        });

        // and the sink anchor links
        _.each(isinks, function(commit) {
            _.each(_.filter(members[commit], isType("logic-state")), function(member) {
                links.push({ source: member, target: cmp.state.anchorR });
            });
        });

        // reset search path vars
        npath = [];
        visited = [];

        while (isources.length !== 0) {
            v = isources.pop();
            //npath.push(members[v.commit]);
            walk(v);
        }

        return [nodes, links, labels];
    },
    shouldComponentUpdate: function(nextProps, nextState) {
        // In the graph object, state is invariant with respect to the message id.
        return nextProps.graph.mid !== this.props.graph.mid;
    },
    render: function() {
        var vizdata = this.extractVizGraph("https://github.com/sdboyer/pipeviz");
        return React.createElement(Viz, {width: this.props.width, height: this.props.height, graph: this.props.graph, nodes: vizdata[0], links: vizdata[1], labels: vizdata[2]})
    },
});

var App = React.createClass({
    dispayName: "pipeviz",
    getInitialState: function() {
        return {
            graph: {},
        };
    },
    getDefaultProps: function() {
        return {
            vizWidth: window.innerWidth * 0.83,
            //vizWidth: window.innerWidth,
            vizHeight: window.innerHeight,
        };
    },
    render: function() {
        return React.createElement("div", {id: "pipeviz"},
                   React.createElement(VizPrep, {width: this.props.vizWidth, height: this.props.vizHeight, graph: this.props.graph})
              );
    },
});

React.render(React.createElement(App, {graph: data}), document.body)
