package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"net/http"
)

var hub = newHub()

func main() {

	router := chi.NewRouter()
	router.Get("/ws", WebSocket)
	fmt.Printf("Starting server at port 8081\n")
	if err := http.ListenAndServe(":8081", router); err != nil {
		logrus.Fatal(err)
	}
	//srv := server.SetupRoutes()
	//err := srv.Run(":8083")
	//if err != nil {
	//	logrus.Error(err)
	//}

}
