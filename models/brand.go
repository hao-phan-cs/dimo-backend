package models

import "database/sql"

type Brand struct {
	ID			int64	`json:"id"`
	Name 		string 	`json:"string"`
	Category 	string 	`json:"string"`
	ImageUrl 	sql.NullString 	`json:"image_url"`
}