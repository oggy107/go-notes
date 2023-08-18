package main

import "fmt"

type Person struct {
	name string
	age  int
	Employee
}

type Employee struct {
	no   int
	role string
}

func structsExample() {
	bob := Person{
		name: "Bob",
		age:  20,
		Employee: Employee{
			no:   1,
			role: "admin",
		},
	}

	bob.info()
	fmt.Println(bob.isAdult())
}

func (p Person) isAdult() bool { return p.age >= 18 }

func (p Person) info() {
	fmt.Println("name", p.name, "age", p.age, "Employee number", p.no, "role", p.role)
}
