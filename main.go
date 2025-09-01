package main

import (
	"log"

	"github.com/nad279444/csrf_token_api/server"
	"github.com/nad279444/csrf_token_api/server/middleware/myJwt"
)

var host = "localhost"
var port = "9000"


func main() {
	jwtErr := myJwt.InitJWT()
	if jwtErr!= nil {
		log.Println("Error initializing the JWT's!")
		log.Fatal(jwtErr)
	}

	// start the server
	serverErr := server.StartServer(host, port)
	if serverErr != nil {
		log.Println("Error starting server!")
		log.Fatal(serverErr)
	}
}