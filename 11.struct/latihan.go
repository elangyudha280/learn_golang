// package main

// import "fmt"

// //! struct 1
// type Book struct {
// 	Title,Author string
// 	Pages int
// }

// //! struct width method
// type Rectangle struct {
// 	width, heigth int
// }

// func (r Rectangle) hitungLuas() int {
// 	return r.width * r.heigth
// }
// func (r *Rectangle) setSize(width,heigth int){
// 	r.width = width
// 	r.heigth = heigth
// }


// //! struct bersarang
// type Address struct {
// 	City , Country string
// }

// type Person struct {
// 	name string
// 	age int
// 	address Address 
// }

// //! mhs
// type Mahasiswa struct {
// 	id int
// 	name string
// 	grade float64
// }

// func (m *Mahasiswa) isPassed() bool {
// 	return m.grade >= 60.0
// }

// func main(){
// 	// book1 := Book{
// 	// 	Title: "Belajar Golang",
// 	// 	Author: "John Doe",
// 	// 	Pages: 100,
// 	// }
// 	// 	book2 := Book{
// 	// 	Title: "Belajar js",
// 	// 	Author: "John Doe",
// 	// 	Pages: 100,
// 	// }
// 	// 	book3 := Book{
// 	// 	Title: "Belajar TS",
// 	// 	Author: "John Doe",
// 	// 	Pages: 100,
// 	// }

// 	mhs1 := Mahasiswa{1,"raka", 75.0}
// 	mhs2 := Mahasiswa{2,"budi", 50.0}
// 	mhs3 := Mahasiswa{3,"siti", 80.0}

// 	 mhsGroup := []Mahasiswa{mhs1, mhs2, mhs3}
// 	// rectangle1 := Rectangle{
// 	// 	width: 10,
// 	// 	heigth: 5,
// 	// }

// 	// //! struct bersarang
// 	// Person1 := Person {
// 	// 	name: "Jane Doe",
// 	// 	age:30,
// 	// 	address: Address{
// 	// 		City: "New York",
// 	// 		Country: "USA",
// 	// 	},
// 	// }

// 	// fmt.Printf("book 1 berjudul %s ditulis oleh %s dan memiliki %d halaman\n", book1.Title, book1.Author, book1.Pages)
// 	// rectangle1.setSize(20, 10)

// 	// fmt.Println(rectangle1)
// 	// fmt.Printf("Luas persegi panjang adalah %d\n", rectangle1.hitungLuas())
// 	// fmt.Print(Person1)

// 	//! looping struct
// 	// for _, books := range []Book{book1, book2, book3} {
// 	// 	fmt.Printf("Buku berjudul %s ditulis oleh %s dan memiliki %d halaman\n", books.Title, books.Author, books.Pages)
// 	// }

// 	for _, mhs := range mhsGroup {
// 		if(!mhs.isPassed()){
// 			continue
// 		}
// 		fmt.Printf("Mahasiswa %s dengan ID %d lulus\n", mhs.name, mhs.id)
// 	}
// }