package main

import "fmt"

//! basic function
func sayHello() {
	fmt.Println("Hello, World!")
}

//! function with parameter
func Greeting(name string, umur string) {
	fmt.Println("Hello, saya ", name, " umur saya ", umur)
}

//! function with return value
func sum(value1, value2 int) int {
	hasil := value1 + value2
	return hasil
}

//! function with multiple return value
func calculate(value1, value2 int) (int, int) {
	return value1 + value2, value1 - value2
}

//! varidaic function
func variadcsFunc(anothervalue ...int) int {
	nilai := 0

	for _, value := range anothervalue {
		nilai += value
	}
	return nilai
}

// type declaration function
type filterWordType func(string) bool

//! function as paramter || callback
func filterWord(word string, ceckWord filterWordType) bool {
	return ceckWord(word)
}

func isNotABadWord(value string) bool {
	if value == "asshole" {
		return true
	}
	return false
}
func main() {
	//! simple function
	// sayHello()

	//! use function with paramter and arguments and single return value
	// Greeting("user", "18") // output: Hello, saya  user  umur saya  18
	// fmt.Println(sum(5, 10)) // output: 15

	//! use function with multer return valu
	// // hasil penjumlahan, pengurangan
	// _, perkalian := calculate(5, 10)
	// print(perkalian)

	//! use variadic function with just value
	// calculate_variadic := variadcsFunc(1,2,3,4,51)
	//! use variadic function with  slice
	// d_slice := []int{1, 2, 3, 4, 51}
	// d_Slice_value := variadcsFunc(d_slice...)
	// fmt.Println(d_Slice_value)

	//! function as value
	get_calculate := calculate
	fmt.Println(get_calculate(1, 2))

	//! function as parameter || callback with another declare function
	fmt.Println(filterWord("asshole2", isNotABadWord))
	//! function as parameter || callback with anonimous function
	fmt.Println(filterWord("asshole", func(value string) bool {
		if value == "asshole" {
			return true
		}
		return false
	}))

	//! anonymous function in variable
	any_func := func(name string) {
		fmt.Println(name)
	}
	any_func("user123")

}
