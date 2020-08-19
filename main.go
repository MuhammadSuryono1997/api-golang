package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func allUsers(w http.ResponseWriter, r *http.Request) {
	var users Users
	var arr_users []Users
	var response Response

	db := connect()
	defer db.Close()

	rows, err := db.Query("Select * from person")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&users.Id, &users.FirstName, &users.LastName); err != nil {
			log.Fatal(err.Error())
		} else {
			arr_users = append(arr_users, users)
		}
	}

	response.Status = 1
	response.Message = "Success"
	response.Data = arr_users

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", allUsers).Methods("GET")
	http.Handle("/", router)
	fmt.Println("Connected to port 1000")
	log.Fatal(http.ListenAndServe(":1000", router))
}
