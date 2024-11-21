package categorycontroller

import (
	"golang-native/models/categorymodel"
	"net/http"
	"html/template"
)

// klik hand enter
func Index(w http.ResponseWriter, r *http.Request) {
	categories := categorymodel.GetAll()
	data := map[string]any{
		"categories" : categories,
	}
//jika ingin melempar data ke file perlu di mapping
	temp, err := template.ParseFiles("views/category/index.html")
	if err != nil {
		panic(err)	
	}

	temp.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {

}

func Edit(w http.ResponseWriter, r *http.Request) {

}

func Delete(w http.ResponseWriter, r *http.Request) {

}
