package main

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Debt defines what an Debt is
type Debt struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Payee       bson.ObjectId `json:"payee" ` //the guy who has to reimburse
	Payer       bson.ObjectId `json:"payer"`  //the guy who paid
	Date        time.Time     `json:"date"`
	PhotoURL    string        `json:"photoURL"`
	Reimbursed  time.Time     `json:"reimbursed"`
}

// Debts is an array of Debts
type Debts []Debt

// GetDebt returns a Debt object from the given ID
func GetDebt(id bson.ObjectId) Debt {
	session, _ := mgo.Dial("127.0.0.1")
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	db := session.DB("reimburse-me").C("debt")
	var debt Debt
	db.FindId(id).One(&debt)
	return debt
}

// GetMyDebts returns an array of Debt objects
// that the userID owe
func GetMyDebts(id bson.ObjectId) Debts {
	session, _ := mgo.Dial("127.0.0.1")
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	db := session.DB("reimburse-me").C("debt")
	var debt Debts
	db.Find(bson.M{"payee": id}).All(&debt)
	return debt
}

// GetTheirDebts returns an array of Debt objects
// that the userID ask to be reimbursed
func GetTheirDebts(id bson.ObjectId) Debts {
	session, _ := mgo.Dial("127.0.0.1")
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	db := session.DB("reimburse-me").C("debt")
	var debt Debts
	db.Find(bson.M{"payer": id}).All(&debt)
	return debt
}

// CreateDebt will add the Debt to the database
func CreateDebt(debt Debt) Debt {
	session, _ := mgo.Dial("127.0.0.1")
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	db := session.DB("reimburse-me").C("debt")
	debt.Date = time.Now()
	db.Insert(debt)
	var result Debt
	db.Find(bson.M{"title": debt.Title, "date": debt.Date}).One(&result)
	return result
}

// DeleteDebt will delete the given debtId
func DeleteDebt(id bson.ObjectId) Debt {
	session, _ := mgo.Dial("127.0.0.1")
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	db := session.DB("reimburse-me").C("debt")
	db.Remove(bson.M{"_id": id})
	var debt Debt
	db.Find(id).One(&debt)
	return debt
}

func ReimburseDebt(id bson.ObjectId) Debt {
	session, _ := mgo.Dial("127.0.0.1")
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	db := session.DB("reimburse-me").C("debt")
	debtID := bson.M{"_id": id}
	change := bson.M{"$set": bson.M{
		"reimbursed": time.Now(),
	}}
	db.Update(debtID, change)
	var debt Debt
	db.FindId(id).One(&debt)
	return debt
}

// AddImageDebt will set the image of the given debtId
func AddImageDebt(id bson.ObjectId, fileName string) Debt {
	session, _ := mgo.Dial("127.0.0.1")
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	db := session.DB("reimburse-me").C("debt")
	eventID := bson.M{"_id": id}
	change := bson.M{
		"photoUrl": fileName,
	}
	db.Update(eventID, change)
	var result Debt
	db.Find(bson.M{"_id": id}).One(&result)
	return result
}
