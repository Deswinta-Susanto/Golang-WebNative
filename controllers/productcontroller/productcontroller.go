package productcontroller

import (
	"golang-native/entities"
	"golang-native/models/productmodel"
	"html/template"
	"net/http"
	"time"

)

func Index(w http.ResponseWriter, r *http.Request) {
	products := productmodel.GetAll()
	data := map[string]any{
		"products" : products,
	}
	temp, err := template.ParseFiles("views/product/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/product/create.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(w, nil)
	}
	if r.Method == "POST" {
		var product entities.Product

		product.Name = r.FormValue("name")
		// product.CategoryId = r.FormValue("category_id")
		// product.Stock = r.FormValue("stock")
		product.Description = r.FormValue("description")
		product.CreatedAt = time.Now()
		product.UpdatedAt = time.Now()

	}
}