package main

import "fmt"

// func factorial(num int) int{
// 	if(num == 1){
// 		return num
// 	}
// 	return  factorial(num - 1) * num
// }

// func sayhello(user string){
// 	fmt.Printf(`helo dunia saya golang senang bertemu dengan mu %v `,user)
// }
// func sum(a,b int) int {
// 	return a+b
// }

// func multiplySum(a,b int) (sum,divided,mod int) {
// 	sum = a+b
// 	divided = a/b
// 	mod = a%b
// 	return sum,divided,mod
// }
// func sumAll(num ...int) int{
// 	// print()\
// 	var total int;
// 	for _,v := range num {
// 		total += v
// 	}
// 	return total
// }
type funcProcess func(angka int) int;
func processNumber(num int,procesFunc funcProcess) int {
	return procesFunc(num)
}

func main() {
	// user := [3]string{
	// 	"user1",
	// 	"user2",
	// 	"user3",
	// }

	// for index, value := range user {
	// 	println(index, value)
	// }

	//! program angkot beroprasi
	// for angkot := 1; angkot <= 10; angkot++ {
	// 	if angkot == 8 {
	// 		fmt.Printf("angkot ke-%v  sedang lembur \n", angkot)
	// 		continue
	// 	} else if angkot >= 5 {
	// 		fmt.Printf("angkot ke-%v tidak beroperasi \n", angkot)
	// 		continue
	// 	}
	// 	fmt.Printf("angkot ke-%v sedang beroperasi \n", angkot)
	// }

	// var inputFile string
	// file, _ := os.ReadFile(`test.txt`)
	
	// fmt.Print("tulis isi file: ")

	// fmt.Scan(file,"&inputFile")
	// fmt.Fprint(file,&inputFile)
	// fmt.Println(factorial(5))
	// fmt.Println("Nama:", name, "Umur:", age)
	// sayhello(`user1`)
	// fmt.Print(sum(1,2))
	// fmt.Print(multiplySum(5,5))
	// fmt.Print(sumAll(1,2))
	// sumAll(1,2)
	var a = processNumber(5, func(value int) int {
		return value * 2	
	})
	fmt.Println(a)

}