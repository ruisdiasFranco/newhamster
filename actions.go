package main

import (
	"fmt"
	"net/http"
	"gopkg.in/mgo.v2"
	"log"
	"encoding/json"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

//-----------------------------------------------------COLLECTION OF SHOPPINGCART
var shoppingCartItems = getSession().DB("Ecommerce").C("shoppingcart")

//-----------------------------------------------------PRODUCT COLLECTION
var productsCollection = getSession().DB("Ecommerce").C("products")

func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost")

	if err != nil{
		panic(err)
	}

	return session
}

//-------------------------------------------------------RESPONSES
func responseShoppingCart(w http.ResponseWriter, status int, results []ShoppingCart){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(results)
}

func responseProduct(w http.ResponseWriter, status int, results Product){

	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(results)
}


func responseProducts(w http.ResponseWriter, status int, results []Product){

	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(results)
}

func Index(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hola mundo desde mi servidor web con GO")
}


//---------------------------------------------------SHOPPING-CART FUNCTIONS
func shoppingCartList (writer http.ResponseWriter, request* http.Request){
	var results []ShoppingCart
	err := shoppingCartItems.Find(nil).Sort("-_id").All(&results)

	if err != nil {
		log.Fatal(err)
	}else{
		fmt.Println("Resultados: ", results)
	}

	responseShoppingCart(writer, 200, results)
}

func addProductToShoppingCart(w http.ResponseWriter, r *http.Request){
	product_id := r.FormValue("producto")

	defer r.Body.Close()
	if !bson.IsObjectIdHex(product_id) {
		w.WriteHeader(404)
		fmt.Println("ERROR")
		return
	}

	oid := bson.ObjectIdHex(product_id)
	fmt.Println(oid)

	cantidad, _ := strconv.ParseInt(r.FormValue("cantidad"), 10, 0)
	datos := Dato{cantidad ,string(oid)}
	fmt.Println(datos)
	err := shoppingCartItems.Insert(datos)

	if err != nil{
		w.WriteHeader(500)
		return
	}

	var results []ShoppingCart
	shoppingCartItems.Find(nil).Sort("-_id").All(&results)

	responseShoppingCart(w, 200, results)
}

func removeProductFromShoppingCart(writer http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	itemId := params["id"]

	if !bson.IsObjectIdHex(itemId) {
		writer.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(itemId)

	err := shoppingCartItems.RemoveId(oid)
	if err != nil{
		writer.WriteHeader(404)
		return
	}

	message := new(Message)

	message.setStatus("success")
	message.setMessage("El producto con ID "+itemId+" ha sido eliminado correctamente")

	results := message
	writer.Header().Set("Content-Type","application/json")
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(results)
}

func shoppingCartUpdate(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	productId := params["id"]

	if !bson.IsObjectIdHex(productId) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(productId)

	decoder := json.NewDecoder(r.Body)

	var product_data Product
	err := decoder.Decode(&product_data)

	if err != nil {
		panic(err)
		w.WriteHeader(500)
		return
	}

	defer r.Body.Close()

	document := bson.M{"_id": oid}
	change := bson.M{"$set": product_data}
	err = shoppingCartItems.Update(document, change)

	if err != nil {
		w.WriteHeader(404)
		return
	}

	var results []ShoppingCart
	shoppingCartItems.Find(nil).Sort("-_id").All(&results)

	responseShoppingCart(w, 200, results)
}


//-----------------------------------------------------------PRODUCTS FUNCTIONS

func productsList(w http.ResponseWriter, r *http.Request){
	var results []Product
	err := productsCollection.Find(nil).Sort("-_id").All(&results)

	if err != nil {
		log.Fatal(err)
	}else{
		fmt.Println("Resultados: ", results)
	}

	responseProducts(w, 200, results)
}

func showProduct(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	productId := params["id"]

	if !bson.IsObjectIdHex(productId){
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(productId)

	results := Product{}
	err := productsCollection.FindId(oid).One(&results)

	if err != nil{
		w.WriteHeader(404)
		return
	}

	responseProduct(w, 200, results)
}

func addProduct(writer http.ResponseWriter, reader *http.Request){
	decoder := json.NewDecoder(reader.Body)

	var product_data Product
	err := decoder.Decode(&product_data)

	if(err != nil){
		panic(err)
	}

	defer reader.Body.Close()

	err = productsCollection.Insert(product_data)

	if err != nil{
		writer.WriteHeader(500)
		return
	}

	responseProduct(writer, 200, product_data)
}

func updateProduct(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	productId := params["id"]

	if !bson.IsObjectIdHex(productId){
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(productId)
	decoder := json.NewDecoder(r.Body)

	var productData Product
	err := decoder.Decode(&productData)

	if err != nil {
		panic(err)
		w.WriteHeader(500)
		return
	}

	defer r.Body.Close()

	document := bson.M{"_id" : oid}
	change := bson.M{"$set": productData}
	err2 := productsCollection.Update(document, change)

	if err2 != nil {
		panic(err2)
		w.WriteHeader(404)
		return
	}

	responseProduct(w, 200, productData)
}

func removeProduct(writer http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	itemId := params["id"]

	if !bson.IsObjectIdHex(itemId) {
		writer.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(itemId)

	err := productsCollection.RemoveId(oid)
	if err != nil{
		writer.WriteHeader(404)
		return
	}

	message := new(Message)

	message.setStatus("success")
	message.setMessage("El producto con ID "+itemId+" ha sido eliminado correctamente")

	results := message
	writer.Header().Set("Content-Type","application/json")
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(results)
}

//-----------------------------------------------------------MESSAGE STRUCT
type Message struct {
	Status string `json:"status"`
	Message string `json:"message"`
}

func (this *Message) setStatus(data string){
	this.Status = data
}

func (this *Message) setMessage(data string){
	this.Message = data
}