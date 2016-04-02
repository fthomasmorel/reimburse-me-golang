package main

import (
	"strconv"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Notification struct {
	ID      bson.ObjectId `bson:"_id,omitempty"`
	User    bson.ObjectId `json:"user"`
	Content string        `json:"content"`
	Debt    Debt          `json:"debt"`
	Date    time.Time     `json:"date"`
}

type Notifications []Notification

func GetNotifications(id bson.ObjectId) Notifications {
	session, _ := mgo.Dial("127.0.0.1")
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	db := session.DB("reimburse-me").C("notification")
	var notifs Notifications
	db.Find(bson.M{"user": id}).Sort("-date").Limit(30).All(&notifs)
	return notifs
}

func AddReimbursedNotification(debt Debt) Notifications {
	session, _ := mgo.Dial("127.0.0.1")
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	db := session.DB("reimburse-me").C("notification")

	var notification1 = GenerateNotificationForReimbursingPayee(debt)
	var notification2 = GenerateNotificationForReimbursingPayer(debt)

	db.Insert(notification1)
	db.Insert(notification2)
	var notif1 Notification
	db.Find(bson.M{"user": notification1.User, "date": notification1.Date}).One(&notif1)
	var notif2 Notification
	db.Find(bson.M{"user": notification2.User, "date": notification2.Date}).One(&notif2)
	return Notifications{notif1, notif2}
}

func AddNewDebtNotification(debt Debt) Notifications {
	session, _ := mgo.Dial("127.0.0.1")
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	db := session.DB("reimburse-me").C("notification")

	var notification1 = GenerateNotificationForNewDebtPayee(debt)
	var notification2 = GenerateNotificationForNewDebtPayer(debt)

	db.Insert(notification1)
	db.Insert(notification2)
	var notif1 Notification
	db.Find(bson.M{"user": notification1.User, "date": notification1.Date}).One(&notif1)
	var notif2 Notification
	db.Find(bson.M{"user": notification2.User, "date": notification2.Date}).One(&notif2)
	return Notifications{notif1, notif2}
}

func GenerateNotificationForNewDebtPayee(debt Debt) Notification {
	var payer = GetUser(debt.Payer)
	var amount = strconv.FormatFloat(float64(debt.Amount), 'f', 2, 32)
	var content = payer.Name + " a ajouté une dette de " + amount + " € pour " + debt.Title + "."
	return Notification{User: debt.Payee, Content: content, Debt: debt, Date: time.Now()}
}

func GenerateNotificationForNewDebtPayer(debt Debt) Notification {
	var payee = GetUser(debt.Payee)
	var amount = strconv.FormatFloat(float64(debt.Amount), 'f', 2, 32)
	var content = "Vous avez ajouté une dette à " + payee.Name + " de " + amount + " € pour " + debt.Title + "."
	return Notification{User: debt.Payer, Content: content, Debt: debt, Date: time.Now()}
}

func GenerateNotificationForReimbursingPayee(debt Debt) Notification {
	var payer = GetUser(debt.Payer)
	var amount = strconv.FormatFloat(float64(debt.Amount), 'f', 2, 32)
	var content = "Vous avez remboursé " + payer.Name + " de " + amount + " € pour " + debt.Title + "."
	return Notification{User: debt.Payee, Content: content, Debt: debt, Date: time.Now()}
}

func GenerateNotificationForReimbursingPayer(debt Debt) Notification {
	var payee = GetUser(debt.Payee)
	var amount = strconv.FormatFloat(float64(debt.Amount), 'f', 2, 32)
	var content = payee.Name + " vous a remboursé(e) " + amount + " € pour " + debt.Title + "."
	return Notification{User: debt.Payer, Content: content, Debt: debt, Date: time.Now()}
}
