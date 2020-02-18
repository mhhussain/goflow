package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mhhussain/goflow/router"
)

func main() {
	fmt.Printf("hi\n")

	rtr := router.GetRouter()

	log.Fatal(http.ListenAndServe(":110", rtr))
}
