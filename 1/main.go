package main

import "fmt"

type Human struct {
	name string
	age  int
}

func (h *Human) sayHello() {
	fmt.Printf("Hello, my name is %s and I'm %d years old.\n", h.name, h.age)
}

func (h *Human) walk() {
	fmt.Println("I'm walking.")
}

// структура Action сможет вызывать все методы
// которые реализованы в Human
// т.е. аналог наследования
type Action struct {
	Human
	action string
}

func (a *Action) doAction() {
	fmt.Printf("I'm doing %s.\n", a.action)
}

func main() {
	action := &Action{
		Human: Human{
			name: "John",
			age:  30,
		},
		action: "running",
	}

	action.sayHello()
	action.walk()
	action.doAction()
}
