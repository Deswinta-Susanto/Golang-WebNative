package categorymodel

import (
	"golang-native/config"
	"golang-native/entities"
)

//function untuk komunikasi dg db

func GetAll() []entities.Category{
	//panggil koneksi
	rows, err := config.DB.Query("SELECT * FROM categories")
	if err != nil{
		panic(err)
	}

	defer rows.Close()

	//variable berisi struck yang telah dibuat
	var categories []entities.Category

	for rows.Next(){
		var category entities.Category
		if err := rows.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdatedAt); err != nil {
			panic(err)
				
	}

	categories = append(categories, category)
	}

	return categories
}