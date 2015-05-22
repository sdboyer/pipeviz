package webapp

import (
	"encoding/json"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
	"time"

	"github.com/gorilla/websocket"
	"github.com/tag1consulting/pipeviz/broker"
	"github.com/tag1consulting/pipeviz/represent"
	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"
)

var (
	assetDir = filepath.Join(defaultBase("github.com/tag1consulting/pipeviz/webapp"), "assets")
	jsDir    = filepath.Join(defaultBase("github.com/tag1consulting/pipeviz/webapp"), "src")
	tmplDir  = filepath.Join(defaultBase("github.com/tag1consulting/pipeviz/webapp"), "tmpl")
)

var (
	// TODO crappily hardcoded, for now
	brokerListen broker.GraphReceiver
	latestGraph  represent.CoreGraph
)

const (
	// Time allowed to write data to the client.
	writeWait = 10 * time.Second
	// Time allowed to read the next pong message from the client.
	pongWait = 60 * time.Second
	// Send pings to client with this period; less than pongWait.
	pingPeriod = (pongWait * 9) / 10
)

var (
	// gorilla websocket upgrader
	upgrader = websocket.Upgrader{ReadBufferSize: 256, WriteBufferSize: 8192}
)

func init() {
	// Subscribe to the master broker and store latest locally as it comes
	brokerListen = broker.Get().Subscribe()
	// FIXME spawning a goroutine in init() used to be crappy, is it still?
	go func() {
		for g := range brokerListen {
			latestGraph = g
		}
	}()

	// Iniitally set the latestGraph to a new, empty one to avoid nil pointer
	latestGraph = represent.NewGraph()
}

// Creates a Goji *web.Mux that can act as the http muxer for the frontend app.
func NewMux() *web.Mux {
	m := web.New()

	m.Use(middleware.Logger)
	m.Get("/assets/*", http.StripPrefix("/assets/", http.FileServer(http.Dir(assetDir))))
	m.Get("/js/*", http.StripPrefix("/js/", http.FileServer(http.Dir(jsDir))))
	m.Get("/", WebRoot)
	m.Get("/sock", OpenSocket)

	return m
}

func graphToJson(g represent.CoreGraph) ([]byte, error) {
	var vertices []interface{}
	for _, v := range g.VerticesWith(represent.Qbv(represent.VTypeNone)) {
		vertices = append(vertices, v.Flat())
	}

	// TODO use something that lets us write to a reusable byte buffer instead
	return json.Marshal(struct {
		Id       int           `json:"id"`
		Vertices []interface{} `json:"vertices"`
	}{
		Id:       g.MsgId(),
		Vertices: vertices,
	})
}

func WebRoot(w http.ResponseWriter, r *http.Request) {
	// TODO first step here is just kitchen sink-ing - send everything.
	j, err := graphToJson(latestGraph)
	vars := struct {
		Title    string
		Vertices string
	}{
		Title:    "pipeviz",
		Vertices: string(j),
	}

	t, err := template.ParseFiles(filepath.Join(tmplDir, "index.html"))
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, vars)
}

func OpenSocket(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		if _, ok := err.(websocket.HandshakeError); !ok {
			log.Println(err)
		}
		return
	}

	go wsWriter(ws)
	wsReader(ws)
}

func wsReader(ws *websocket.Conn) {
	defer ws.Close()
	ws.SetReadLimit(512)
	ws.SetReadDeadline(time.Now().Add(pongWait))
	ws.SetPongHandler(func(string) error { ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	// for now, we ignore all messages from the frontend
	for {
		_, _, err := ws.ReadMessage()
		if err != nil {
			break
		}
	}
}

func wsWriter(ws *websocket.Conn) {
	graphIn := broker.Get().Subscribe()
	pingTicker := time.NewTicker(pingPeriod)
	defer func() {
		pingTicker.Stop()
		broker.Get().Unsubscribe(graphIn)
		ws.Close()
	}()

	// write the current graph state first, before entering loop
	graphToSock(ws, latestGraph)
	var g represent.CoreGraph
	for {
		select {
		case <-pingTicker.C:
			// ensure client connection is healthy
			ws.SetWriteDeadline(time.Now().Add(writeWait))
			if err := ws.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		case g = <-graphIn:
			graphToSock(ws, g)
		}
	}
}

func graphToSock(ws *websocket.Conn, g represent.CoreGraph) {
	j, err := graphToJson(g)
	if err != nil {
		log.Println(err)
	}

	if j != nil {
		ws.SetWriteDeadline(time.Now().Add(writeWait))
		if err := ws.WriteMessage(websocket.TextMessage, j); err != nil {
			return
		}
	}
}
