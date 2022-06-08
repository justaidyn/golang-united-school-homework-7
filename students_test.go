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
func TestSwap(t *testing.T){
	var people People

	people = append(people, Person{
		firstName: "John",
		lastName: "Doe",
		birthDay: time.Date(1980, 1, 1, 0, 0, 0, 0, time.UTC),
	})
	people = append(people, Person{
		firstName: "Jane",
		lastName: "Doe",
		birthDay: time.Date(1980, 1, 1, 0, 0, 0, 0, time.UTC),
	})
	people = append(people, Person{
		firstName: "Joe",
		lastName: "Doe",
		birthDay: time.Date(1980, 1, 1, 0, 0, 0, 0, time.UTC),
	})

	people.Swap(0,1)

	if people[0].firstName != "Jane" {
		t.Errorf("Wrong Swap")
	}
	if people[1].firstName != "John" {
		t.Errorf("Wrong Swap")
	}
}