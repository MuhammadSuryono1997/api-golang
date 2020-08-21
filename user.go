package main

import (
	"encoding/json"
	"log"
	"net/http"
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

func insertUsers(w http.ResponseWriter, r *http.Request) {
	var user Users
	var arr_users []Users
	var response Response

	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	first_name := r.Form.Get("first_name")
	last_name := r.Form.Get("last_name")

	_, err = db.Exec("insert into person (first_name, last_name) values (?,?)", first_name, last_name)
	if err != nil {
		log.Print(err)
	}

	user.FirstName = first_name
	user.LastName = last_name
	arr_users = append(arr_users, user)

	response.Status = 1
	response.Message = "Successfuly Add Data"
	response.Data = arr_users
	log.Print("Insert data to database")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func updateUsers(w http.ResponseWriter, r *http.Request) {
	var response Response

	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	id := r.Form.Get("user_id")
	first_name := r.Form.Get("first_name")
	last_name := r.Form.Get("last_name")

	_, err = db.Exec("UPDATE person set first_name = ?, last_name = ? where id = ?",
		first_name,
		last_name,
		id,
	)

	if err != nil {
		log.Print(err)
	}

	response.Status = 1
	response.Message = "Success Update Data"
	log.Print("Update data to database")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func deleteUsers(w http.ResponseWriter, r *http.Request) {
	var response Response

	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	id := r.Form.Get("user_id")

	_, err = db.Exec("DELETE from person where id = ?",
		id,
	)

	if err != nil {
		log.Print(err)
	}

	response.Status = 1
	response.Message = "Success Delete Data"
	log.Print("Delete data to database")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
