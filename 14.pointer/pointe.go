// package main

// import "fmt"

// func getData(data *int) {
// 	*data = 20 //* -> ini artinya kita akan mengubah nilai dari variabel yang ditunjuk oleh pointer data
// }

// type Person struct {
// 	name string
// 	age  int
// }

// func (p *Person) setAge(age int) {
// 	p.age = age
// }

// func main() {

// 	//! pointer variable
// 	// x := 10
// 	// p := &x //* ->  ini artinya kita akan menggunakan alamat asli dari variabel x sehingga tidak akan menduplikasi data yg ada di variable x

// 	// *p = 20 //* -> ini artinya kita akan mengubah nilai dari variabel yang ditunjuk oleh pointer p
// 	// // fmt.Print(*p,x)
// 	//! pointer fungsi
// 	// h := 15
// 	// getData(&h) //* ->  ini artinya kita akan menggunakan alamat asli dari variabel h sehingga tidak akan menduplikasi data yg ada di variable h
// 	// fmt.Print(h)

// 	//! pointer di struct
// 	// p1 := Person{"John", 30}
// 	// p1.setAge(35) 
// 	// // p2.age = 40
// 	// fmt.Print(p1.age)

// 	//! default pointer
// 	var dp *string;
// 	if dp == nil {
// 		dp = new(string)
// 		*dp = "hello go"
// 	}
// 	fmt.Print(*dp)

// }