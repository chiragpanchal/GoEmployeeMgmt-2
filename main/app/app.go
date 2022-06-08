package app

import (
	"GoEmployeeMgmt/main/service"
	"GoEmployeeMgmt/main/util"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func StartApp() {

	fmt.Println("entering StartApp")
	router := mux.NewRouter()

	dbClient := util.DbConnection()

	//ph := PeopleHandler{service: service.NewPeopleService(util.RepoDb{DbClient: dbClient})}
	ph2 := PeopleHandler{service: service.NewPeopleService(util.RepoDb{DbClient: dbClient})}
	router.HandleFunc("/people", ph2.getAll).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8080", router))

	fmt.Println("leaving StartApp")
}
