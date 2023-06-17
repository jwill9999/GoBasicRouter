package cmd

import (
  "net/http"
  "main/Router"
)

func ServerStart() {
  r := &router.Router{}
  router.RegisterRoutes(r);
  http.ListenAndServe(":8000", r)
  
}