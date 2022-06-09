package coverage

import (
	"os"
	"testing"
	"time"
	"reflect"
	"errors"
	"strconv"
	"fmt"
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
		firstName: "Ivan",
		lastName:  "Drago",
		birthDay:  time.Date(1985, time.May, 25, 0, 0, 0, 0, time.UTC),
	},
	{
		firstName: "Olexa",
		lastName:  "Кvitka",
		birthDay:  time.Date(1995, time.December, 05, 0, 0, 0, 0, time.UTC),
	},
	{
		firstName: "Olexa",
		lastName:  "Drobush",
		birthDay:  time.Date(1995, time.December, 05, 0, 0, 0, 0, time.UTC),
	},
	{
		firstName: "John",
		lastName:  "Doe",
		birthDay:  time.Date(1995, time.December, 05, 0, 0, 0, 0, time.UTC),
	},
}

func TestPeople_Len(t *testing.T) {
	Expected := 4 // who knows whether we need here len(PersonList) or not :)
	got := PersonList.Len()

	if got != Expected {
		t.Errorf("Expected %d, got %d", Expected, got)
	}
}

func TestPeople_Less(t *testing.T) {
	tData := map[string]struct {
		FirstPersonIndex  int
		SecondPersonIndex int
		Expected          bool
	}{
		"diffDates":             {0, 1, false},
		"sameDateDiffNames":     {1, 3, false},
		"sameDateSameFirstName": {1, 2, false},
	}

	for name, v := range tData {
		t.Run(name, func(t *testing.T) {
			got := PersonList.Less(v.FirstPersonIndex, v.SecondPersonIndex)
			if got != v.Expected {
				t.Errorf("[%s] expected: %v, got %v", name, v.Expected, got)
			}
		})
	}
}

func TestPeople_Swap(t *testing.T) {
	tData := struct {
		FirstPeopleInd  int
		SecondPeopleInd int
		Expected        Person
	}{
		0,
		1,
		PersonList[0],
	}
	PersonList.Swap(tData.FirstPeopleInd, tData.SecondPeopleInd)
	got := PersonList[tData.SecondPeopleInd]
	if got != tData.Expected {
		t.Errorf("Expected %v, got %v", tData.Expected, got)
	}
}

/////////////////////////////////////////////////////////////////////

func TestNew(t *testing.T) {
	tData := map[string]struct {
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
			errors.Unwrap(fmt.Errorf("Rows need to be the same length")), // this is strange
			// could not find better solution
		},
		"ImproperData": {
			"a",
			nil,
			strconv.ErrSyntax,
		},
	}

	for name, v := range tData {
		t.Run(name, func(t *testing.T) {
			got, err := New(v.TestString)
			if err != nil && !errors.Is(errors.Unwrap(err), v.Err) { // could not find better solution, think
				//  the error message from the New() needs some changes for better logic
				t.Errorf("[%s] error expected:\"%v\", got:\"%v\".\n", name, v.Err, err)
			}
			if !reflect.DeepEqual(v.Expected, got) { // needs to be remembered when compare structs with slices&maps
				t.Errorf("[%s] expected: %v, got %v", name, v.Expected, got)
			}
		})
	}
}

func TestMatrix_Rows(t *testing.T) {
	tData := struct {
		TestMatrix *Matrix
		Expected   [][]int
	}{
		&Matrix{2, 3, []int{1, 1, 1, 2, 2, 2}}, // should I use New() here?
		[][]int{{1, 1, 1}, {2, 2, 2}},
	}
	got := tData.TestMatrix.Rows()
	if !reflect.DeepEqual(tData.Expected, got) {
		t.Errorf("Expected %v, got %v", tData.Expected, got)
	}
}

func TestMatrix_Cols(t *testing.T) {
	tData := struct {
		TestMatrix *Matrix
		Expected   [][]int
	}{
		&Matrix{2, 3, []int{1, 1, 1, 2, 2, 2}}, // should I use New() here?
		[][]int{{1, 2}, {1, 2}, {1, 2}},
	}
	got := tData.TestMatrix.Cols()
	if !reflect.DeepEqual(tData.Expected, got) {
		t.Errorf("Expected %v, got %v", tData.Expected, got)
	}
}

func TestMatrix_Set(t *testing.T) {
	baseMatrix := &Matrix{2, 3, []int{1, 1, 1, 2, 2, 2}}
	needMatrixT := &Matrix{2, 3, []int{1, 1, 1, 2, 5, 2}}
	needMatrixF := &Matrix{2, 3, []int{1, 1, 1, 2, 2, 2}}

	tData := map[string]struct {
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

	for name, v := range tData {
		t.Run(name, func(t *testing.T) {
			got := baseMatrix.Set(v.row, v.col, v.value)
			if got != v.Expected && !reflect.DeepEqual(v.ExpectedMatrix, baseMatrix) {
				t.Errorf("[%s] expected: %v, got %v", name, v.Expected, got)
			}
		})
	}
}