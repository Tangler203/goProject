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
	//var user = "joyceL"
	var d bson.D
	//var d []Account

	coll := reader.C("Bank")
	//err = coll.Find(bson.M{"user": user}).All(&d)
	err = coll.Find(bson.M{ "user": "JoyceL" }).Select(bson.M{"user": 0, "pass": 0}).One(&d)
	fmt.Println("getting data")
	if err != nil {
		fmt.Println("Query Error")
		panic(err)
	}

	for i, elem := range d {
		//if elem == "Liam Joyce" {
		i++
		fmt.Println(elem)

		//fmt.Println("Done")
		//return
		//else{
		//fmt.Println(elem.Name, elem.Value, i)
		//}

		//acc.name = elem.name
		//fmt.Println(acc.name, i)
	}
	//fmt.Println(d)
}
