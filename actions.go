package newhamster

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)



func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost")

	if err != nil{
		panic(err)
	}

	return session
}


var collection = getSession().DB("Ecommerce").C("products")


func shoppingCartList



func shoppingCartAddProduct