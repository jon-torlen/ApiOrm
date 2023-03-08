package main

import (
	//"orm/models"
	 "orm/db/handlers"
	"net/http"
	 "log"
	 "github.com/gorilla/mux"
)


func main(){

	//models.MigrarUser()
	
	//Rutas
	mux := mux.NewRouter()

	//EndPoints
	mux.HandleFunc("/api/user/", handlers.GetUsers).Methods("GET")
	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.GetUser).Methods("GET")
	mux.HandleFunc("/api/user/", handlers.CreateUser).Methods("POST")
	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.UpdateUser).Methods("PUT")
	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", mux))
}