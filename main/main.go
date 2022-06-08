package main

import (
	"GoEmployeeMgmt/main/app"
	"fmt"
)

//dependency packages:
//go get -u github.com/lib/pq
//go get github.com/jmoiron/sqlx
//go get -u github.com/gorilla/mux

func main() {

	fmt.Println("app started")
	app.StartApp()
}
