package dbs

import (
	"fmt"
	"html/template"
	"net/http"

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
	http.HandleFunc("/", loadHome)
	readDb()
	http.ListenAndServe(":8080", nil)
}

func loadHome(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/home.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	t.Execute(w, nil)
}

func readDb() {
	sessionState, err := mgo.Dial("127.0.0.1:27017")

	if err != nil {
		panic(err)
	}

	defer sessionState.Close()

	reader := sessionState.DB("bank")
	
	d := Account{}

	coll := reader.C("Bank")
	
	err = coll.Find(bson.M{ "user": "JoyceL" }).Select(bson.M{"user": 0, "pass": 0}).One(&d)
	//fmt.Println("getting data")
	if err != nil {
		fmt.Println("Query Error")
		panic(err)
	}

	fmt.Println(d.User)
		
}
