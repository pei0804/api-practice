package main

import (
	"app/controller"
	"encoding/json"
	"net/http"

	"github.com/zenazn/goji"
)

type ErrorResponse struct {
	Status   int
	Messages []string
}

func main() {
	m := goji.DefaultMux
	controller.SetUpSample(m)
	goji.Serve()
}

func (er *ErrorResponse) Write(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(er.Status)
	json.NewEncoder(w).Encode(er.Messages)
}
