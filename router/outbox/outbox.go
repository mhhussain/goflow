package outbox

import (
	"encoding/json"
	"net/http"

	"github.com/mhhussain/goflow/outbox"
)

func OutboxRouteHandler(w http.ResponseWriter, r *http.Request) {
	var incomingReq outbox.Outbound

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&incomingReq); err != nil {
		resp, _ := json.Marshal("invalid request")
		w.WriteHeader(500)
		w.Write(resp)
		return
	}

	// create response channel
	respchan := make(chan outbox.Response)

	outbound := outbox.OutboundRequest{
		incomingReq,
		&respchan,
	}

	// send to outbox channel
	outbox.Box <- outbound

	// wait on response channel
	x := <-respchan
	w.WriteHeader(200)
	w.Write(x.Body.([]byte))

	//fmt.Printf("done: %#v\n", string(x.Body.([]byte)))
}
