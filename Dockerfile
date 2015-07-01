# Pipeviz Dockerfile 0.0.1
#
## To use this Dockerfile, do the following:

## Build a docker container image tagged pipeviz
# docker build -t pipeviz .

## Run a docker container named pipeserv using the image you just built
# docker run -ti -p 8008:8008 --rm --name pipeserv  pipeviz

## In a seperate shell, run a separate container to manually send input to 
## the pipeviz server:
# docker run -ti --link pipeserv:pipeserv --rm pipeviz pvutil fixr fixtures/realistic/ -t http://pipeserv:2309

FROM debian:latest
MAINTAINER Michael Halstead <halstead@happypunch.com>

ENV GO_VERSION 1.4.2
ENV GOPATH /opt/go
ENV PATH /usr/local/go/bin:$GOPATH/bin:$PATH

RUN apt-get update && apt-get install -y --no-install-recommends \
	wget \
	git \
	npm \
	&& rm -rf /var/lib/apt/lists/*
RUN ln -v /usr/bin/nodejs /usr/bin/node
RUN npm install -g bower
RUN wget -q https://storage.googleapis.com/golang/go$GO_VERSION.linux-amd64.tar.gz && tar -C /usr/local -xzf go$GO_VERSION.linux-amd64.tar.gz
RUN mkdir -p $GOPATH/src/github.com/tag1consulting/pipeviz
COPY . $GOPATH/src/github.com/tag1consulting/pipeviz
RUN cd $GOPATH/src/github.com/tag1consulting/pipeviz/webapp && bower install --allow-root --quiet
RUN cd $GOPATH/src/github.com/tag1consulting/pipeviz && go build . && \
	go get -d ./... && \
	go build ./... && \
	go install && \
	go install ./cmd/...  
WORKDIR $GOPATH/src/github.com/tag1consulting/pipeviz
CMD ["pipeviz", "-b"]
