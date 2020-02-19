package outbox

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var Box = make(chan OutboundRequest, 100)

func ProcessOutbox() {
	go func() {
		for {
			select {
			case outbound := <-Box:
				// OutboundRequest
				// make request and respond on queue
				makeOutboundRequest(outbound)
			}
		}
	}()
}

func makeOutboundRequest(outbound OutboundRequest) {
	// marshal body of request
	j, _ := json.Marshal(outbound.Outbound.Request.Body.(interface{}))

	// send request
	req, _ := http.NewRequest(outbound.Outbound.Request.Verb, outbound.Outbound.Request.ServiceLink, bytes.NewBuffer(j))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	// handle fail
	if err != nil {
		errResp := Response{
			"500",
			err,
		}

		*(outbound.RespChan) <- errResp
		return
	}
	defer resp.Body.Close()

	// handle success
	body, _ := ioutil.ReadAll(resp.Body)
	sucResp := Response{
		"200",
		body,
	}
	*(outbound.RespChan) <- sucResp
}
