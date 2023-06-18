package router

import (
	"net/http"
)

func RegisterRoutes(r *Router) {

	r.Route(http.MethodGet, "/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("The server is now running in the browser!"))
	})

	r.Route(http.MethodGet, `/hello/(?P<Message>\w+)`, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello " + URLParam(r, "Message")))
	})

	r.Route(http.MethodGet, "/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("something bad happened!")
	})

}
