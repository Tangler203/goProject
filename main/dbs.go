package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

func makeDb() {
	sessionState, err := mgo.Dial("127.0.0.1:27017")

	if err != nil {
		fmt.Println("Connection Error")
		panic(err)
	}

	defer sessionState.Close()

	reader := sessionState.DB("bank")

	coll := reader.C("Bank")

	coll.RemoveAll(nil)
	err = coll.Insert(&Account{Name: "Liam Joyce", Number: 12345678, Amount: 2000.95, CreditRating: "Good", User: "JoyceL", Pass: "78ad45"},
		&Account{Name: "Edel Shaw", Number: 60985521, Amount: 7742.99, CreditRating: "Good", User: "YellowSquare", Pass: "DaisyB00k"},
		&Account{Name: "Michael Sheehan", Number: 46872439, Amount: 1078, CreditRating: "Poor", User: "sheehan87", Pass: "password"},
		&Account{Name: "Lilly Jones", Number: 77896235, Amount: 5050.50, CreditRating: "Good", User: "LJones", Pass: "050690"})
	if err != nil {
		fmt.Println("Insert error")
		panic(err)
	}

	fmt.Println("Database created")

}
