package main

import (
	"fmt"
	"net/http"
	"time"
	"math/rand"

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
	m.Post("/createDb", makeDb)
	m.Get("/createaccount", create)
	m.Post("/account", account)
	m.Run()
}

func login(ctx *macaron.Context, req *http.Request) {
	var username string
	var password string
	err := req.ParseForm()
	if err != nil {
		panic(err)
	}
	username = req.FormValue("user")
	password = req.FormValue("password")
	if username == "" {

		fmt.Println("No input written")
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
	err = coll.Find(bson.M{"user": username, "pass": password}).Select(bson.M{"name": 1, "number": 1, "amount": 1, "creditrating": 1, "user": 1, "pass": 1}).One(&d)
	if err != nil {
		fmt.Println("Query Error")
		panic(err)
	}
	return d

}

func makeDb(ctx *macaron.Context, req *http.Request) {
	sessionState, err := mgo.Dial("127.0.0.1:27017")

	if err != nil {
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
		panic(err)
	}
	ctx.HTML(200, "home")
}

func create(ctx *macaron.Context, req *http.Request) {
	ctx.HTML(200, "create")
}

func account(ctx *macaron.Context, req *http.Request){
	var username string
	var password string
	var name string
	err := req.ParseForm()
	if err != nil {
		panic(err)
	}

	name = req.FormValue("name")
	username = req.FormValue("user")
	password = req.FormValue("password")
	r1 := rand.NewSource(time.Now().UnixNano())
	r2 := rand.New(r1)
	accountnum := r2.Intn(100000000)

	addaccount(name,username,password,accountnum)
	
	ctx.HTML(200, "home")
}

func addaccount(name string, user string, password string, acc int){
	sessionState, err := mgo.Dial("127.0.0.1:27017")

	if err != nil {
		panic(err)
	}

	defer sessionState.Close()

	reader := sessionState.DB("bank")

	coll := reader.C("Bank")

	err = coll.Insert(&Account{Name: name, Number: acc, Amount: 1000, CreditRating: "Good", User: user, Pass: password})
		if err != nil {
		panic(err)
	}
}
