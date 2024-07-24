package main

import (
	"log"
	"net/http"
)

func main() {

	muxHandler := http.NewServeMux()

	handleRoutes(muxHandler)

	// declarartion of our server we can clean this up maybe...
	server := http.Server{
		Addr:    ":9000",
		Handler: muxHandler,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Println("an error has occured")
		return
	}

}

func handleRoutes(muxHandler *http.ServeMux) {

	// just so we can check if server is up and running or not.....
	muxHandler.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("ok"))
	})

	muxHandler.HandleFunc("/ws", handleWs)

}
