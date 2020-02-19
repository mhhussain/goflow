package outbox

type OutboundRequest struct {
	Outbound	Outbound
	RespChan	*chan Response
}

type Outbound struct {
	Caller	string	`json:"caller"`
	Request	Request	`json:"request"`
}

type Request struct {
	ServiceLink		string		`json:"serviceLink"`
	Headers			interface{}	`json:"headers,omitempty"`
	Body			interface{} `json:"body"`
}

type Response struct {
	Body	interface{}	`json:"body"`
}
