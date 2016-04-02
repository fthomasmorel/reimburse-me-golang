package main

import (
	"encoding/json"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
)

func GetNotificationController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]
	var res = GetNotifications((bson.ObjectIdHex(userID)))
	json.NewEncoder(w).Encode(res)
}
