package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
)

func LogUserController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user := GetUserWithToken(bson.ObjectIdHex(vars["id"]))
	token := vars["token"]
	fmt.Println(token)
	fmt.Println(user.Token)
	if user.Token == token {
		sessionToken := LogUser(user.ID)
		json.NewEncoder(w).Encode(bson.M{"token": sessionToken})
	} else {
		json.NewEncoder(w).Encode(bson.M{"errorr": "wrong id"})
	}
}

func GetUserController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]
	if len(userID) > 15 {
		var res = GetUser(bson.ObjectIdHex(userID))
		json.NewEncoder(w).Encode(res)
	} else {
		var res = GetUserFromUsername(userID)
		json.NewEncoder(w).Encode(res)
	}

}

func CreateUserController(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user User
	decoder.Decode(&user)
	res := CreateUser(user)
	json.NewEncoder(w).Encode(res)
}

func DeleteUserController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]
	res := DeleteUser(bson.ObjectIdHex(userID))
	json.NewEncoder(w).Encode(res)
}

func AddPayeeController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	res := AddPayee(bson.ObjectIdHex(vars["id"]), bson.ObjectIdHex(vars["payeeID"]))
	json.NewEncoder(w).Encode(res)
}

func RemovePayeeController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	res := RemovePayee(bson.ObjectIdHex(vars["id"]), bson.ObjectIdHex(vars["payeeID"]))
	json.NewEncoder(w).Encode(res)
}
