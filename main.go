package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/users", allUsers).Methods("GET")
	router.HandleFunc("/users", insertUsers).Methods("POST")
	router.HandleFunc("/users", updateUsers).Methods("PUT")
	router.HandleFunc("/users/{id}", deleteUsers).Methods("DELETE")
	http.Handle("/", router)
	fmt.Println("Connected to port 1000")
	log.Fatal(http.ListenAndServe(":1000", router))
}
