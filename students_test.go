package coverage

import (
	"os"
	"testing"
	"time"
	"errors"
	"fmt"
	"strconv"
	"reflect"
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
	data := map[string]struct {
		TestString string
		Expected   *Matrix
		Err        error
	}{
		"ProperData": {
			"1 1 1\n2 2 2",
			&Matrix{2, 3, []int{1, 1, 1, 2, 2, 2}},
			nil,
		},
		"DiffRowsLength": {
			"1 1 1\n2 2 2 2",
			nil,
			errors.Unwrap(fmt.Errorf("Rows need to be the same length")), 
		},
		"ImproperData": {
			"a",
			nil,
			strconv.ErrSyntax,
		},
	}

	for name, v := range data {
		t.Run(name, func(t *testing.T) {
			got, err := New(v.TestString)
			if err != nil && !errors.Is(errors.Unwrap(err), v.Err) { 
				t.Errorf("[%s] error expected:\"%v\", got:\"%v\".\n", name, v.Err, err)
			}
			if !reflect.DeepEqual(v.Expected, got) { 
				t.Errorf("[%s] expected: %v, got %v", name, v.Expected, got)
			}
		})
	}
}


func TestSet(t *testing.T){
	baseMatrix := &Matrix{2, 3, []int{1, 1, 1, 2, 2, 2}}
	needMatrixT := &Matrix{2, 3, []int{1, 1, 1, 2, 5, 2}}
	needMatrixF := &Matrix{2, 3, []int{1, 1, 1, 2, 2, 2}}

	data := map[string]struct {
		row            int
		col            int
		value          int
		ExpectedMatrix *Matrix
		Expected       bool
	}{
		"ProperData":     {1, 1, 5, needMatrixT, true},
		"RowLessO":       {-1, 1, 5, baseMatrix, false},
		"ImproperRowNum": {2, 1, 5, needMatrixF, false},
		"ImproperColNum": {1, 3, 5, needMatrixF, false},
	}

	for name, v := range data {
		t.Run(name, func(t *testing.T) {
			got := baseMatrix.Set(v.row, v.col, v.value)
			if got != v.Expected && !reflect.DeepEqual(v.ExpectedMatrix, baseMatrix) {
				t.Errorf("[%s] expected: %v, got %v", name, v.Expected, got)
			}
		})
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