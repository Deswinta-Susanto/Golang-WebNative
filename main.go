package main

import (
	"golang-native/config"
	"golang-native/controllers/categorycontroller"
	"golang-native/controllers/homecontroller"
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

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)

}
