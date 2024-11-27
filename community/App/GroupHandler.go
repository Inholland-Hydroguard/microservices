package app

import (
	"encoding/json"
	"log"
	Services "microservices/community/Services"
	"net/http"

	"github.com/gorilla/mux"
)

type Grouphandlers struct {
	Service Services.GroupService
}

func (ch *Grouphandlers) GetAllGroups(w http.ResponseWriter, r *http.Request) {
	Groups, err := ch.Service.GetAllGroup()
	if err != nil {
		log.Println(err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Groups)
}

func (ch *Grouphandlers) FindById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	Group, err := ch.Service.FindGroupById(params["id"])
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, Group)
	}
}
