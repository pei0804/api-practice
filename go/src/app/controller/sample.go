package controller

import (
	"app/model"
	"encoding/json"
	"net/http"

	"github.com/zenazn/goji/web"
)

type SampleController struct {
}

func SetUpSample(m *web.Mux) {
	api := SampleController{}
	m.Get("/api/v1/sample", api.Get)
}

func (a *SampleController) Get(c web.C, w http.ResponseWriter, r *http.Request) {
	s := model.Sample{}
	sa := s.Get()

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sa)
}
