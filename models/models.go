package models

type Employee struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Calories int    `json:"calories"`
}
