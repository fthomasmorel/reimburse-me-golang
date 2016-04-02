package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
)

func GetDebtController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	debtID := vars["id"]
	var res = GetDebt(bson.ObjectIdHex(debtID))
	json.NewEncoder(w).Encode(res)
}

func GetMyDebtsController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userID"]
	var res = GetMyDebts(bson.ObjectIdHex(userID))
	json.NewEncoder(w).Encode(res)
}

func GetTheirDebtsController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userID"]
	var res = GetTheirDebts(bson.ObjectIdHex(userID))
	json.NewEncoder(w).Encode(res)
}

func CreateDebtController(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var debt Debt
	decoder.Decode(&debt)
	vars := mux.Vars(r)
	userID := vars["userID"]
	debt.Payer = bson.ObjectIdHex(userID)
	res := CreateDebt(debt)
	json.NewEncoder(w).Encode(res)
}

func DeleteDebtController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	debt := DeleteDebt(bson.ObjectIdHex(vars["id"]))
	json.NewEncoder(w).Encode(debt)
}

func ReimburseDebtController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	res := ReimburseDebt(bson.ObjectIdHex(vars["id"]))
	json.NewEncoder(w).Encode(res)
}

func AddImageDebtController(w http.ResponseWriter, r *http.Request) {
	fileName := UploadImage(r)
	if fileName == "error" {
		fmt.Println("wrong file name")
		w.Header().Set("status", "400")
		fmt.Fprintln(w, "{}")
	} else {
		vars := mux.Vars(r)
		res := AddImageDebt(bson.ObjectIdHex(vars["id"]), fileName)
		json.NewEncoder(w).Encode(res)
	}
}
