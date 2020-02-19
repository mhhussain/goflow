package outbox

import (
	"fmt"
	"net/http"

	"github.com/mhhussain/goflow/outbox"
)

func OutboxRouteHandler(w http.ResponseWriter, r *http.Request) {
	out := outbox.Outbound{
		"test",
		outbox.Request{
			"GET",
			"http://buckets-dev.xby2-rnd.com/buckets/drop-four/testfile0009110000000",
			nil,
			`{"hello":"world"}`,
		},
	}

	// create response channel
	respchan := make(chan outbox.Response)

	outbound := outbox.OutboundRequest{
		out,
		&respchan,
	}

	// send to outbox channel
	outbox.Box <- outbound

	// wait on response channel
	x := <-respchan

	fmt.Printf("done: %#v\n", string(x.Body.([]byte)))
}
