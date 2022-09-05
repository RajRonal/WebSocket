package main

import "net/http"

func WebSocket(response http.ResponseWriter, request *http.Request) {
	serveWs(hub, response, request)
}
