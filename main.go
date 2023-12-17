package main

import (
	"database/sql"
	"fmt"

	"gofr.dev/pkg/gofr"
)

type Student struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Calories int    `json:"calories"`
}

var db *sql.DB

func main() {
	// Connect to your database
	conn, err := initDatabase()
	if err != nil {
		fmt.Println("Error initializing database:", err)

	}
	//db = conn
	defer db.Close()

	app := gofr.New()

	app.POST("/healthify/{name}/{age}/{calories}", createStudent)
	app.GET("/healthify", getAllStudents)
	app.PUT("/healthify/{id}", updateStudent)
	app.DELETE("/healthify/{id}", deleteStudent)

	app.Start()
}
func initDatabase() (*sql.DB, error) {

	host := "localhost"
	port := "3306"
	user := "root"
	password := "zopsmart123#"
	database := "health"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database)
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("error validating sql.open")
		return nil, err
	}

	return db, nil
}
