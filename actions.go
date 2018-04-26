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


func Index(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hola mundo desde mi servidor web con GO")
}

func shoppingCartList(w http.ResponseWriter, r *http.Request){
	var results []Product
			err := collection.Find(nil).Sort("-_id").All(&results)

			if err != nil {
				log.Fatal(err)
			}else{
				fmt.Println("Resultados: ", results)
			}

			responseMovies(w, 200, results)

}



func shoppingCartAddProduct(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)

	var product Product 
	err := decoder.Decode(&product)

	if(err != nil){
		panic(err)
	}

	defer r.Body.Close()

	err = collection.Insert(product)

	if err != nil{
		w.WriteHeader(500)
		return
	}

	responseMovie(w, 200, product)
}