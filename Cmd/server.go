package cmd

import (
	"log"
	router "main/Router"
	"net/http"
)

func ServerStart() {
	const PORT = ":8000"
	r := &router.Router{}
	router.RegisterRoutes(r)
	log.Printf("Server running on port %v", PORT)
	http.ListenAndServe(PORT, r)

}
