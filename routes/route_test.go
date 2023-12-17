package route

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gofr.dev/gofr"
	// Modify to your actual package path
)

func TestCreateStudent(t *testing.T) {
	// Setup mock database
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDB := mock_database.NewMockDB(ctrl)

	// Mock expected behavior
	student := &Student{Name: "John Doe", Age: 30, Calories: 2000}
	mockDB.EXPECT().ExecContext(gomock.Any(), "INSERT INTO health (name,age,calories) VALUES (?,?,?)", student.Name, student.Age, student.Calories).Return(sql.Result{}, nil)

	// Prepare request body
	body, err := json.Marshal(student)
	if err != nil {
		t.Fatal(err)
	}

	// Create mock request
	req, err := http.NewRequest(http.MethodPost, "/healthify/John Doe/30/2000", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	// Create mock context
	ctx := gofr.NewContext(req)

	// Test endpoint
	response, err := createStudent(ctx)
	assert.Nil(t, err)

	// Check response type
	message, ok := response.(map[string]string)
	assert.True(t, ok)

	// Assert response message
	assert.Equal(t, "Member created successfully", message["message"])
}

func TestGetAllStudents(t *testing.T) {
	// Setup mock database
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDB := mock_database.NewMockDB(ctrl)

	// Mock expected behavior
	students := []*Student{
		{ID: 1, Name: "Alice", Age: 25, Calories: 1500},
		{ID: 2, Name: "Bob", Age: 32, Calories: 2200},
	}
	mockDB.EXPECT().QueryContext(gomock.Any(), "SELECT * FROM health").Return(&mock_database.MockRows{
		Result: mock_database.MockRows{
			Rows: [][]interface{}{
				{students[0].ID, students[0].Name, students[0].Age, students[0].Calories},
				{students[1].ID, students[1].Name, students[1].Age, students[1].Calories},
			},
		},
	}, nil)

	// Create mock context
	ctx := gofr.NewContext(&http.Request{})

	// Test endpoint
	response, err := getAllStudents(ctx)
	assert.Nil(t, err)

	// Check response type and length
	actualStudents, ok := response.([]*Student)
	assert.True(t, ok)
	assert.Equal(t, len(students), len(actualStudents))

	// Assert student data
	for i, actualStudent := range actualStudents {
		assert.Equal(t, students[i].ID, actualStudent.ID)
		assert.Equal(t, students[i].Name, actualStudent.Name)
		assert.Equal(t, students[i].Age, actualStudent.Age)
		assert.Equal(t, students[i].Calories, actualStudent.Calories)
	}
}
