package outbox

var Box = make(chan OutboundRequest, 100)

func ProcessOutbox() {
	go func() {
		for {
			select {
			case request := <- Box:
				// OutboundRequest
				// make request and respond on queue
				resp := Response{
					"world",
				}
				*(request.RespChan) <- resp
			}
		}
	}()
}