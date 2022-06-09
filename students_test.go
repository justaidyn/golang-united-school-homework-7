package coverage

import (
	"os"
	"testing"
	"time"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW

var PersonList People = []Person{
	{
		firstName: "John",
		lastName:  "Doe",
		birthDay:  time.Date(1980, time.January, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		firstName: "Jane",	
		lastName:  "Doe",
		birthDay:  time.Date(1980, time.January, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		firstName: "John",
		lastName:  "Smith",
		birthDay:  time.Date(1980, time.January, 1, 0, 0, 0, 0, time.UTC),
	},
}

func TestLen(t *testing.T){
	if PersonList.Len() != 3 {
		t.Error("Expected 3, got ", PersonList.Len())
	}
}
func TestSwap(t *testing.T){
	PersonList.Swap(0, 1)
	if PersonList[0].firstName != "Jane" {
		t.Error("Expected Jane, got ", PersonList[0].firstName)
	}
	if PersonList[1].firstName != "John" {
		t.Error("Expected John, got ", PersonList[1].firstName)
	}
}




///////////////////////////////////////////////////////////////////////////////////////////////////////


func TestNew(t *testing.T){
	matrix, err := New("1 2\n3 4")
	if err != nil {
		t.Error("Expected nil, got ", err)
	}
	if matrix.rows != 2 {
		t.Error("Expected 2, got ", matrix.rows)
	}
	if matrix.cols != 2 {
		t.Error("Expected 2, got ", matrix.cols)
	}
	if matrix.data[0] != 1 {
		t.Error("Expected 1, got ", matrix.data[0])
	}
	if matrix.data[1] != 2 {
		t.Error("Expected 2, got ", matrix.data[1])
	}
	if matrix.data[2] != 3 {
		t.Error("Expected 3, got ", matrix.data[2])
	}
	if matrix.data[3] != 4 {
		t.Error("Expected 4, got ", matrix.data[3])
	}
}


func TestSet(t *testing.T){
	matrix, err := New("1 2\n3 4")
	if err != nil {
		t.Error("Expected nil, got ", err)
	}
	matrix.Set(0, 0, 5)
	if matrix.data[0] != 5 {
		t.Error("Expected 5, got ", matrix.data[0])
	}
}

func TestRows(t *testing.T){
	data := struct{
		TestMatrix *Matrix
		Expected [][]int
	}{
		&Matrix{2,3,[]int{1,2,3,4,5,6}},
		[][]int{
			{1,2,3},
			{4,5,6},
		},
	}
	get := data.TestMatrix.Rows()
	if len(get) != len(data.Expected) {
		t.Error("Expected ", len(data.Expected), " got ", len(get))
	}
	for i := 0; i < len(get); i++ {
		if len(get[i]) != len(data.Expected[i]) {
			t.Error("Expected ", len(data.Expected[i]), " got ", len(get[i]))
		}
		for j := 0; j < len(get[i]); j++ {
			if get[i][j] != data.Expected[i][j] {
				t.Error("Expected ", data.Expected[i][j], " got ", get[i][j])
			}
		}
	}
}

func TestCols(t *testing.T){
	data := struct{
		TestMatrix *Matrix
		Expected [][]int
	}{
		&Matrix{2,3,[]int{1,2,3,4,5,6}},
		[][]int{
			{1,4},
			{2,5},
			{3,6},
		},
	}
	get := data.TestMatrix.Cols()
	if len(get) != len(data.Expected) {
		t.Error("Expected ", len(data.Expected), " got ", len(get))
	}
	for i := 0; i < len(get); i++ {
		if len(get[i]) != len(data.Expected[i]) {
			t.Error("Expected ", len(data.Expected[i]), " got ", len(get[i]))
		}
		for j := 0; j < len(get[i]); j++ {
			if get[i][j] != data.Expected[i][j] {
				t.Error("Expected ", data.Expected[i][j], " got ", get[i][j])
			}
		}
	}	
}