package server

import (
	"github.com/nad279444/csrf_token_api/server/middleware"
	"log"
	"net/http"
)


func StartServer(hostname string, port string) error{
	host :=  hostname + ":" + port
	log.Printf("Starting server at %s " + host)

	handler := middleware.NewHandler()

	http.Handle("/", handler)

	return http.ListenAndServe(host, nil)
}