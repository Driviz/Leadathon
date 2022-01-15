package service

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lead/chessgames"
)

func NewService(data chessgames.DataMap) *service {
	return &service{
		Data: data,
	}
}

type service struct {
	Data chessgames.DataMap
}

func (s *service) StartService() {
	r := mux.NewRouter()
	r.HandleFunc("/", s.GetAll)
	r.HandleFunc("/{code}", s.GetByCode)

	http.ListenAndServe(":8080", r)
}
