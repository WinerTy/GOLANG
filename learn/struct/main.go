package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	age       int
}

type Android struct {
	Person
	Model string
}

func main() {
	p := Android{
		Person: Person{
			"Ivan",
			"Ivanov",
			30,
		},
		Model: "First Model",
	}
	fmt.Println(p.full_name() + " " + p.Model)
}

func (p *Person) full_name() string {
	return p.FirstName + " " + p.LastName
}
