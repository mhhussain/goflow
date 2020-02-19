package main

import (
	//"os"
	"fmt"
	//"log"
	//"net/http"

	//"github.com/mhhussain/goflow/router"
	"github.com/mhhussain/goflow/outbox"
)

var port int

/*
outbound = {
	caller: 'service',
	request: {
		serviceLink: '',
		headers: {},
		body: {}
	}
}

*/

func main() {

	//port = os.Getenv("SERVICE_PORT")
	/*port = 110
	servePort := fmt.Sprintf(":%d", port)

	rtr := router.GetRouter()

	fmt.Printf("Listening on port %s", servePort)
	log.Fatal(http.ListenAndServe(servePort, rtr))*/

	outbox.ProcessOutbox()

	fmt.Println("world")

	out := outbox.Outbound{
		"test",
		outbox.Request{
			"",
			nil,
			"helloworld",
		},
	}

	respchan := make(chan outbox.Response)

	outbound := outbox.OutboundRequest{
		out,
		&respchan,
	}

	outbox.Box <- outbound
	
	x := <- respchan

	fmt.Printf("done: %#v\n", x.Body)
}
