package main

import "fmt"

// struct di golang
type Person struct {
	//* property di struct
	Name   string
	age    int
	health int
}

//* method di struct
func (paramStruct *Person) sayHello() {
	fmt.Printf("hello %v senang bertemu dengan mu \n", paramStruct.Name)
}

func (paramStruct *Person) healing(jumlah int) {
	if paramStruct.health < 100 {
		paramStruct.health += jumlah
		if paramStruct.health > 100 {
			paramStruct.health = 100
		}
		fmt.Printf("Health setelah healing: %v\n", paramStruct.health)
		return
	}
	fmt.Printf("tidak butuh healing, health: %v", paramStruct.health)
}

func main() {
	user1 := Person{
		Name:   "user1",
		age:    20,
		health: 50,
	}

	// user1.healing(10)
	user1.healing(10)
	// user1.healing(10)

	fmt.Println(user1)
	// user2.sayHello()
}
