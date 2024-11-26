package main

import (
	"golang-native/config"
	"golang-native/controllers/categorycontroller"
	"golang-native/controllers/homecontroller"
	"golang-native/controllers/productcontroller"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	//memanggil homepage controller.function
	http.HandleFunc("/", homecontroller.Welcome)

	//memanggil menu kategory
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)

	http.HandleFunc("/product", productcontroller.Index)
	http.HandleFunc("/product/add", productcontroller.Add)
	// http.HandleFunc("/product/edit", productcontroller.Edit)
	// http.HandleFunc("/product/delete", productcontroller.Delete)

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)

}
