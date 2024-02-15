package models

import "gorm.io/gorm"

//Task is the base for our task
type Task struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
}
