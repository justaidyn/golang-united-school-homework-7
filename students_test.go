package coverage

import (
	"os"
	"testing"
	"time"
	"reflect"
	"errors"
	"strconv"
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

func TestLen(t *testing.T) {
	var people People

	if people.Len() != 0 {
		t.Errorf("Wrong People Len empty")
	}

	people = append(people, Person{})
	people = append(people, Person{})
	people = append(people, Person{})

	if people.Len() != 3 {
		t.Errorf("Wrong People Len 3")
	}

	people = people[0:1]

	if people.Len() != 1 {
		t.Errorf("Wrong People Len 1")
	}
}

func TestLess(t *testing.T) {
	var people People
	now := time.Time{}

	people = append(people, Person{"BBB", "BBB", now})
	people = append(people, Person{"BBB", "BBB", now.Add(5 * time.Minute)})
	people = append(people, Person{"AAA", "BBB", now})
	people = append(people, Person{"AAA", "AAA", now})
	people = append(people, Person{"AAA", "AAA", now})

	if people.Less(0, 1) {
		t.Errorf("Wrong People Less by Birthday")
	}

	if people.Less(0, 2) {
		t.Errorf("Wrong People Less by FirstName")
	}

	if people.Less(2, 3) {
		t.Errorf("Wrong People Less by LastName")
	}

	if people.Less(3, 4) {
		t.Errorf("Wrong People Less Eq")
	}
}

func TestSwap(t *testing.T) {
	var people People
	now := time.Time{}

	people1 := Person{"first", "first", now}
	people2 := Person{"first", "first", now}

	people = append(people, people1)
	people = append(people, people2)

	people.Swap(0, 1)

	if people[0] != people2 || people[1] != people1 {
		t.Errorf("Wrong People Swap")
	}
}

func TestNew(t *testing.T) {
	actual, err := New("---")
	if actual != nil || !errors.Is(err, strconv.ErrSyntax) {
		t.Errorf("Wrong String Error")
	}

	actual, err = New("1 1 \n 2")
	if actual != nil || err.Error() != "Rows need to be the same length" {
		t.Errorf("Wrong Matrix Error")
	}

	actual, err = New("1 1 \n 2 3")
	expects := &Matrix{2, 2, []int{1, 1, 2, 3}}

	if actual.cols != expects.cols || actual.rows != expects.rows || !reflect.DeepEqual(actual.data, expects.data) {
		t.Errorf("Wrong Empty Matrix")
	}
}

func TestRows(t *testing.T) {
	matrix := &Matrix{2, 2, []int{4, 5, 6, 7}}
	expects := [][]int{{4, 5}, {6, 7}}

	actual := matrix.Rows()

	if !reflect.DeepEqual(actual, expects) {
		t.Errorf("Wrong Rows Matrix")
	}
}

func TestCols(t *testing.T) {
	matrix := &Matrix{3, 3, []int{1, 10, 100, 2, 20, 200, 3, 30, 300}}
	expects := [][]int{{1, 2, 3}, {10, 20, 30}, {100, 200, 300}}

	actual := matrix.Cols()

	if !reflect.DeepEqual(actual, expects) {
		t.Errorf("Wrong Cols Matrix")
	}
}

func TestSet(t *testing.T) {
	matrix := &Matrix{3, 3, []int{1, 10, 100, 2, 20, 200, 3, 30, 300}}
	expectsData := []int{1, 10, 100, 2, 20, 200, 3, 30, 10000000}

	actual := matrix.Set(2, 2, 10000000)

	if !actual || !reflect.DeepEqual(matrix.data, expectsData) {
		t.Errorf("Wrong Set Matrix")
	}

	actual = matrix.Set(-1, 2, 10000000)
	if actual {
		t.Errorf("Wrong Set Error")
	}
}