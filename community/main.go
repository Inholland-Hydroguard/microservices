package main

import (
	"fmt"
	"log"
	"net/http"

	app "microservices/community/App"
	"microservices/community/Domain"
	"microservices/community/Services"

	"github.com/gorilla/mux"
)

const webPort = ":8080"

func main() {
	fmt.Println("Starting App")

	var router = mux.NewRouter()

	GroupRepo := Domain.NewGroupRepositoryDB()
	GroupServices := Services.NewGroupService(GroupRepo)

	var GroupHandlers = app.Grouphandlers{Service: GroupServices}

	router.HandleFunc("/Groups", GroupHandlers.GetAllGroups).
		Methods("GET").
		Name("GetAllGroups")

	router.HandleFunc("/Groups/{id}", GroupHandlers.FindById).
		Methods("GET").
		Name(" Group")

	fmt.Println("Starting Web Server on port", webPort)
	log.Fatal(http.ListenAndServe(webPort, router))

}
