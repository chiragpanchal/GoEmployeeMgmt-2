package app

import (
	"GoEmployeeMgmt/main/service"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type PeopleHandler struct {
	service service.PeopleService
}

func (ps *PeopleHandler) getAll(w http.ResponseWriter, r *http.Request) {

	fmt.Println("entering app.getAll")

	people, err := ps.service.GetAllPeople()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(people)
	if err != nil {
		return
	}

	fmt.Println("leaving app.getAll")

}

func (ps *PeopleHandler) getOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	searchedPersonId, parseErr := strconv.ParseInt(vars["personId"], 10, 0)
	if parseErr != nil {
		return
	}
	person, err := ps.service.GetOnePerson(int(searchedPersonId))
	if err != nil {
		return
	}

	err = json.NewEncoder(w).Encode(person)
	if err != nil {
		return
	}

}
