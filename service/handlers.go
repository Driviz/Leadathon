package service

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type GetResponse struct {
}

func (s *service) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	data, err := json.Marshal(s.Data)
	if err != nil {
		log.Fatalln("error marshalling json", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (s *service) GetByCode(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	params := mux.Vars(r)
	code := params["code"]

	res := s.Data[code]
	if res.Moves == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	data, err := json.Marshal(s.Data[code])
	if err != nil {
		log.Fatalln("error marshalling json", err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
