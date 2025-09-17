package main

import "fmt"

type StructKosong struct {}

type interfaceKosong interface{}

func main() {
	var angka int
	var angkaPecahan float64
	var booleanvalue bool
	var teks string

	var sliceKosong []int
	var arrayKosong [3]int
	var structKosong StructKosong

	fmt.Printf("default value dari tipe data angka %d\n", angka)
	fmt.Printf("default value dari tipe data angka pecahan %f\n", angkaPecahan)
	fmt.Printf("default value dari tipe data boolean %t\n", booleanvalue)
	fmt.Printf("default value dari tipe data teks %q\n", teks)
	fmt.Printf("default value dari tipe data slice kosong %v\n", sliceKosong)
	fmt.Printf("default value dari tipe data array kosong %v\n", arrayKosong)
	fmt.Printf("default value dari tipe data struct kosong %v\n", structKosong)
	fmt.Printf("default value dari tipe data interface kosong %v\n", interfaceKosong(nil))


	var valueBebas any = "hey im string"

	// fmt.Println(valueBebas.(int)) //* -> ini golang akan panic karena nilainya tidak sesuai dengan tipe data yang ditentukan

	// fmt.Println(valueBebas.(string)) //* -> ini akan berhasil karena nilainya sesuai dengan tipe data yang ditentukan	

	//! check with if else
	if str,ok := valueBebas.(string); ok {
		fmt.Println("Nilai adalah string:", str,ok)
	} else {
		fmt.Println("Nilai bukan string",ok)
	}
	


}