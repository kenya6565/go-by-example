package main

import "fmt"

type Stringfy interface {
	ToString() string
}

type Person struct {
	Name string
	Age  int
}

func (p *Person) ToString() string {
	return fmt.Sprintf("Name=%v, Age=%v", p.Name, p.Age)
}

type Car struct {
	Number string
	Model  string
}

func (c *Car) ToString() string {
	return fmt.Sprintf("Number=%v, Model=%v", c.Number, c.Model)
}

type MyError struct {
	Message string
	ErrCode int
}

func (e *MyError) Error() string {
	return e.Message
}

func RaiseError() error {
	return &MyError{
		Message: "some error",
		ErrCode: 1,
	}
}

type Animal interface {
	Speak() string
}

type Dog struct {
	name string
}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {
}

func (c Cat) Speak() string {
	return "Meow!"
}

func getAnimal() []Animal {
	animals := []Animal{Dog{name: "jiggy"}, Cat{}}
	return animals
}

func interface_main() {
	vs := []Stringfy{
		&Car{Number: "124-22", Model: "AB-1234"},
	}

	fmt.Println(vs[0])

	// for _, v := range vs {
	// 	fmt.Println(v.ToString())
	// }

	// animals := getAnimal()
	// fmt.Println(animals)
	// for _, animal := range animals {
	// 	fmt.Println(animal.Speak())
	// }

	// err := RaiseError()
	// fmt.Println(err.Error())

}
