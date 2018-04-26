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

<<<<<<< HEAD
func productAdd(writer http.ResponseWriter, reader *http.Request){
	decoder := json.NewDecoder(reader.Body)

	var movie_data Movie
	err := decoder.Decode(&movie_data)

	if(err != nil){
		panic(err)
	}

	defer reader.Body.Close()

	err = collection.Insert(movie_data)

	if err != nil{
		writer.WriteHeader(500)
		return
	}

	responseMovie(writer, 200, movie_data)
}
func producUpdate()

func productRemove(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	movie_id := params["id"]

	if !bson.IsObjectIdHex(movie_id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(movie_id)

	decoder := json.NewDecoder(r.Body)

	var movie_data Movie
	err := decoder.Decode(&movie_data)

	if err != nil {
		panic(err)
		w.WriteHeader(500)
		return
	}

	defer r.Body.Close()

	document := bson.M{"_id": oid}
	change := bson.M{"$set": movie_data}
	err = collection.Update(document, change)

	if err != nil {
		w.WriteHeader(404)
		return
	}

	responseMovie(w, 200, movie_data)
}
=======
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
>>>>>>> kaku
