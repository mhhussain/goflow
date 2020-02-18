package main

import (
	//"os"
	"fmt"
	"log"
	"net/http"

	"github.com/mhhussain/goflow/router"
)

var port int

func main() {

	//port = os.Getenv("SERVICE_PORT")
	port = 110
	servePort := fmt.Sprintf(":%d", port)

	rtr := router.GetRouter()

	fmt.Printf("Listening on port %s", servePort)
	log.Fatal(http.ListenAndServe(servePort, rtr))
}
