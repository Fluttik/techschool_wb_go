package main

import "fmt"

type Human struct {
	name    string
	surname string
	age     int8
}

func (h Human) Greet() {
	fmt.Println("Привет! Меня зовут", h.name, h.surname)
}

type Action struct {
	Human
}

func (h Action) Run() {
	fmt.Println(h.name, "побежал")
}

func main() {
	a := Action{
		Human: Human{
			name:    "Николай",
			surname: "Королёв",
			age:     29,
		},
	}
	a.Greet()
	a.Run()
}
