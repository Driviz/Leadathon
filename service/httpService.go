package service

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Driviz/Leadathon/chessgames"
	"github.com/gorilla/mux"
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

	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), r)
}
