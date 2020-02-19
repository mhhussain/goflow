package main

import (
	//"os"
	"fmt"
	//"log"
	//"net/http"

	//"github.com/mhhussain/goflow/router"
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

type Outbound struct {
	Caller	string	`json:"caller"`
	Request	Request	`json:"request"`
}

type OutboundRequest struct {
	Outbound	Outbound
	RespChan	*chan Response
}

type Request struct {
	ServiceLink		string		`json:"serviceLink"`
	Headers			interface{}	`json:"headers,omitempty"`
	Body			interface{} `json:"body"`
}

type Response struct {
	Body	interface{}	`json:"body"`
}

var q = make(chan OutboundRequest, 1)

func main() {

	//port = os.Getenv("SERVICE_PORT")
	/*port = 110
	servePort := fmt.Sprintf(":%d", port)

	rtr := router.GetRouter()

	fmt.Printf("Listening on port %s", servePort)
	log.Fatal(http.ListenAndServe(servePort, rtr))*/

	fmt.Println("world")

	out := Outbound{
		"test",
		Request{
			"",
			nil,
			"helloworld",
		},
	}

	respchan := make(chan Response)

	outbound := OutboundRequest{
		out,
		&respchan,
	}

	q <- outbound

	go func() {
		a := <- q
		fmt.Printf("in: %s\n", a.Outbound.Caller)

		rr := Response{
			"",
		}
		*(a.RespChan) <- rr
	}()

	x := <- respchan

	fmt.Printf("done: %#v\n", x.Body)
}
