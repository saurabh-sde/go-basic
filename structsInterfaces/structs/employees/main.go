package main

import (
	"encoding/json"
	"fmt"
)

/*
Nested Structs Declaration Syntax ->

	type ParentStruct struct {
	  	...
		Child   ChildStruct
		or
		ChildStruct (if Key is named as Struct name only)
	}

	type ChildStruct {
	  	...
	}
*/

// Department - store department details
type Department struct {
	DepartmentName string
	Empoyees       []Employee // list of employees in department
}

// Employee - store Employee details
type Employee struct {
	EmployeeId int
	Details    Person
	Contact
}

// Person - store Personal details
type Person struct {
	Name   string
	Age    int
	Job    string
	Salary int
}

type Contact struct {
	Mobile int
	Email  string
}

func main() {
	// PersonA
	var personA Person
	personA.Name = "A"
	personA.Age = 20
	personA.Job = "Engineer"
	personA.Salary = 100000
	// create employee A object
	employeeA := Employee{
		EmployeeId: 1,
		Details:    personA,
		Contact:    Contact{Mobile: 1234, Email: "personA@test"},
	}

	// PersonB
	personB := Person{
		Name:   "B",
		Age:    25,
		Job:    "Manager",
		Salary: 120000,
	}
	// create employee A object
	employeeB := Employee{
		EmployeeId: 2,
		Details:    personB,
		Contact:    Contact{Mobile: 1234, Email: "personB@test"},
	}

	employees := []Employee{employeeA, employeeB}
	tech := Department{
		DepartmentName: "TECH",
		Empoyees:       employees,
	}
	fmt.Printf("Tech : %+v\n", tech)

	// print struct data as JSON data to mock API response data
	jsonData, _ := json.MarshalIndent(tech, "", "\t")
	fmt.Printf("jsonData: %+v\n", string(jsonData))
}
