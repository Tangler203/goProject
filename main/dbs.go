package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Account struct {
	//Id           bson.ObjectId 'bson:"_id,omiyempty"'
	Name         string //'bson:"name"'
	Number       int
	Amount       float32
	CreditRating string
	User         string
	Pass         string
}

func main() {
	makeDb()
}

func makeDb() {
	sessionState, err := mgo.Dial("127.0.0.1:27017")

	if err != nil {
		panic(err)
	}

	defer sessionState.Close()

	reader := sessionState.DB("bank")

	coll := reader.C("Bank")



	coll.RemoveAll(nil)
	err = coll.Insert(&Account{Name: "Liam Joyce", Number: 12345678, Amount: 2000.95, CreditRating: "Good", User: "JoyceL", Pass: "78ad45"} ,
	&Account{Name: "Edel Shaw", Number: 60985521, Amount: 7742.99, CreditRating: "Good", User: "YellowSquare", Pass: "DaisyB00k"},
	&Account{Name: "Liam Joyce", Number: 12345678, Amount: 2000.95, CreditRating: "Good", User: "JoyceL", Pass: "78ad45"} ,
	&Account{Name: "Liam Joyce", Number: 12345678, Amount: 2000.95, CreditRating: "Good", User: "JoyceL", Pass: "78ad45"})
	if err != nil {
		fmt.Println("Insert error")
		panic(err)
	}

	// Query One
	result := Account{}
	err = coll.Find(bson.M{ "user": "JoyceL" }).Select(bson.M{"name": 1, "user": 1}).One(&result)
	if err != nil {
		fmt.Println("Query Error")
		panic(err)
	}
	fmt.Println(result)
}
