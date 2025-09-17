package main

import "fmt"

// func logger(){
// 	fmt.Println("program selesai")
// }

// func divide(a,b int) int {

// 	defer fmt.Println("function divide end execution")
// 	defer func ()  {
// 		recover()
// 	} ()

// 	if(a == 0 || b == 0){
// 		panic("tidak bisa membagi angka 0")
// 	}

// 	fmt.Println(a,b)

// 	return a/b
// }

func step1(){
	// defer func ()  {
	// 	if r := recover(); r != nil {
	// 		fmt.Println("Tangkap panic: step1", r)
	// 	}
	// }()
	panic("step 1  error")
	
	// fmt.Println("this is step1")
}
func step2(){

	fmt.Println("this is step2")
}
func step3(){
	
	fmt.Println("this is step3")
}

func process()  {
	// defer
	defer func ()  {
		fmt.Println(recover())	
	}()

	step1()
	step2()
	step3()

	fmt.Println("end functino process")
}

func checkUser(name string,age int) {
	defer func ()  {
		recover()
	}()

	if(name ==  "" || age < 18){
		panic("nama tidak boleh kosong atau umur harus di atas 18")
	}

	fmt.Printf("helo user %v umur kamu %v umur \n",name,age)
}

func main() {
	// defer logger()
	// panic("somesthing wrong")
	// defer fmt.Println("hello world")
	// defer fmt.Println("hello world2")
	// defer fmt.Println("hello world3")


	// divide1 := divide(0,1)
	// divide2 := divide(5,2)
	// fmt.Println("divide1",divide1)
	// fmt.Println("divide1",divide2)

	// checkUser("yudha",20)
	// checkUser("pram",22)
	process()
	fmt.Println("last programs")
	}	