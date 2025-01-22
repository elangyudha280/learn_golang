package main

import "fmt"

func main() {

	// * secara jelas menentukan jumlah index dan element yg ada di dalam array
	var angka [3]int
	name := [3]string{
		"user1",
		"user2",
		"user3",
	}
	//* set value
	angka[0] = 1
	angka[1] = 2
	angka[2] = 3

	//* get value
	fmt.Println(len(angka))
	fmt.Println(name)

	// *  menentukan secara otomatis jumlah element yg ada di dalam array
	mahasiswa := [...]string{
		"mahasiswa1",
		"mahasiswa2",
		"mahasiswa3",
	}
	var dosen = [...]string{
		"dosen1",
		"dosen2",
	}

	fmt.Println(mahasiswa)
	fmt.Println(dosen)

}
