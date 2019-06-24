package main

import "fmt"

// Person Struct Definition
type Person struct {
	name string
	age  int
	job  string
}

func (person Person) say() {
	fmt.Printf("My name is %s", person.name)
}

func getOnePerson(name, job string, age int) Person {
	return Person{name, age, job}
}

func main() {
	p := getOnePerson("coma", "worker", 23)
	p.say()
}
