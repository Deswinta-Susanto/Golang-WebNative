package entities

import "time"

type Product struct {
	Id int
	Name string
	CategoryId int
	Stock int
	Description string
	CreatedAt time.Time
	UpdatedAt time.Time
}