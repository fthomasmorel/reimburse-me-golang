package main

import (
	"os/exec"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// User defines what an User is
type User struct {
	ID       bson.ObjectId   `bson:"_id,omitempty"`
	Username string          `json:"username"`
	Name     string          `json:"name"`
	Token    string          `json:"token"`
	Payees   []bson.ObjectId `json:"payees" bson:"payees,omitempty"`
}

// LogUser returns a session token
func LogUser(id bson.ObjectId, token string) string {
	///....
	return "fdsqjkm"
}

// GetUserFromUsername returns an user object from the given ID
func GetUserFromUsername(username string) User {
	session, _ := mgo.Dial("127.0.0.1")
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	db := session.DB("reimburse-me").C("user")
	var result User
	db.Find(bson.M{"username": username}).One(&result)
	result.Token = ""
	//result.Payees = []bson.ObjectId{} problem with public content
	return result
}

// GetUser returns an user object from the given ID
func GetUser(id bson.ObjectId) User {
	session, _ := mgo.Dial("127.0.0.1")
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	db := session.DB("reimburse-me").C("user")
	var result User
	db.FindId(id).One(&result)
	result.Token = ""
	//result.Payees = []bson.ObjectId{} problem with public content
	return result
}

// CreateUser will create the user in the database
func CreateUser(user User) User {
	session, _ := mgo.Dial("127.0.0.1")
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	db := session.DB("reimburse-me").C("user")
	token, _ := exec.Command("uuidgen").Output()
	user.Token = string(token[:])
	db.Insert(user)
	var result User
	db.Find(bson.M{"username": user.Username}).One(&result)
	return result
}

// DeleteUser will delete the given User
func DeleteUser(id bson.ObjectId) User {
	session, _ := mgo.Dial("127.0.0.1")
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	db := session.DB("reimburse-me").C("user")
	db.Remove(bson.M{"_id": id})
	var result User
	db.Find(id).One(result)
	return result
}

// AddPayee add the given payeeID to the given payees as a payee
func AddPayee(id bson.ObjectId, payeeID bson.ObjectId) User {
	session, _ := mgo.Dial("127.0.0.1")
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	db := session.DB("reimburse-me").C("user")
	userID := bson.M{"_id": id}
	change := bson.M{"$addToSet": bson.M{
		"payees": payeeID,
	}}
	db.Update(userID, change)
	var user User
	db.Find(bson.M{"_id": id}).One(&user)
	return user
}

// RemovePayee remove the given payeeID to the given payees list
func RemovePayee(id bson.ObjectId, payeeID bson.ObjectId) User {
	session, _ := mgo.Dial("127.0.0.1")
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	db := session.DB("reimburse-me").C("user")
	userID := bson.M{"_id": id}
	change := bson.M{"$pull": bson.M{
		"payees": payeeID,
	}}
	db.Update(userID, change)
	var user User
	db.Find(bson.M{"_id": id}).One(&user)
	return user
}
