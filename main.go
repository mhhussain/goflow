package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mhhussain/goflow/outbox"
	"github.com/mhhussain/goflow/router"
)

var port int

func main() {

	//port = os.Getenv("SERVICE_PORT")
	port = 110
	servePort := fmt.Sprintf(":%d", port)

	rtr := router.GetRouter()
	outbox.ProcessOutbox()

	fmt.Printf("Listening on port %s", servePort)
	log.Fatal(http.ListenAndServe(servePort, rtr))

}
