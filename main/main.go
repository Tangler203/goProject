package main

import (
	"fmt"
	"net/http"

	"gopkg.in/macaron.v1"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Account struct {
	//Id           bson.ObjectId 'bson:"_id,omiyempty"'
	Name         string
	Number       int
	Amount       float32
	CreditRating string
	User         string
	Pass         string
}

func main() {
	m := macaron.Classic()
	m.Use(macaron.Renderer())
	m.Get("/", func(ctx *macaron.Context) {
		ctx.HTML(200, "home") // 200 is the response code.
	})
	m.Post("/login", login)
	m.Post("/login/", login)
	m.Run()
}

func login(ctx *macaron.Context, req *http.Request) {
	var username string
	var password string
	err := req.ParseForm()
	if err != nil {
		fmt.Println("No input written")
		username = "JoyceL"
		password = "78ad45"
	} else {
		username = req.FormValue("user")
		password = req.FormValue("password")
	}

	results := readDb(username, password)
	ctx.Data["Name"] = results.Name
	ctx.Data["Number"] = results.Number
	ctx.Data["Amount"] = results.Amount
	ctx.Data["CreditRating"] = results.CreditRating
	ctx.Data["User"] = results.User
	ctx.Data["Pass"] = results.Pass
	ctx.HTML(200, "result")
}

func readDb(username string, password string) Account {
	sessionState, err := mgo.Dial("127.0.0.1:27017")

	if err != nil {
		panic(err)
	}

	defer sessionState.Close()

	reader := sessionState.DB("bank")
	d := Account{}

	coll := reader.C("Bank")
	fmt.Println(username, " : ", password)
	err = coll.Find(bson.M{"user": username}).Select(bson.M{"name": 1, "number": 1, "amount": 1, "creditrating": 1, "user": 1, "pass": 1}).One(&d)
	if err != nil {
		fmt.Println("Query Error")
		panic(err)
	}
	return d

}
