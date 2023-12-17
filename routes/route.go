package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"

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
func createStudent(ctx *gofr.Context) (interface{}, error) {
	name := ctx.PathParam("name")
	calories := ctx.PathParam("calories")
	age, err := strconv.Atoi(ctx.PathParam("age"))
	if err != nil {
		return nil, errors.New("invalid age parameter")
	}

	_, err = db.ExecContext(ctx, "INSERT INTO health (name,age,calories) VALUES (?,?,?)", name, age, calories)
	if err != nil {
		return nil, err
	}

	return map[string]string{"message": "Member created successfully"}, nil
}

func getAllStudents(ctx *gofr.Context) (interface{}, error) {
	var students []Student

	rows, err := db.QueryContext(ctx, "SELECT * FROM health")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var student Student
		err := rows.Scan(&student.ID, &student.Name, &student.Age, &student.Calories)
		if err != nil {
			return nil, err
		}

		students = append(students, student)
	}

	return students, nil
}

func updateStudent(ctx *gofr.Context) (interface{}, error) {
	id, err := strconv.Atoi(ctx.PathParam("id"))
	if err != nil {
		return nil, errors.New("invalid id parameter")
	}

	body, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		return nil, err
	}

	var requestBody struct {
		NewName string `json:"new_name"`
	}

	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		return nil, errors.New("invalid JSON payload")
	}

	_, err = db.ExecContext(ctx, "UPDATE health SET name = ? WHERE id = ?", requestBody.NewName, id)
	if err != nil {
		return nil, err
	}

	return map[string]string{"message": "Member updated successfully"}, nil
}

func deleteStudent(ctx *gofr.Context) (interface{}, error) {
	id, err := strconv.Atoi(ctx.PathParam("id"))
	if err != nil {
		return nil, errors.New("invalid id parameter")
	}

	_, err = db.ExecContext(ctx, "DELETE FROM health WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	return map[string]string{"message": "Member deleted successfully"}, nil
}
