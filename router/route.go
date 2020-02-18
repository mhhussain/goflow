package router

import (
	"net/http"

	"github.com/mhhussain/goflow/router/ping"
)

type Route struct {
	Name		string
	Method		string
	Pattern		string
	HandlerFunc	http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Ping",
		"GET",
		"/ping",
		ping.PingRouteHandler,
	},
}