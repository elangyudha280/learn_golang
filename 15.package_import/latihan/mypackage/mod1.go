package mypackage

import "fmt"


var version = "1" //* -> variable ini tidak bisa di akses diluar package karena menggunakan huruf kecil diawalannya

func HelloWorld() { //* -> function ini bisa di akses diluar fungsi karena pke huruf kapital awalannya
	fmt.Println("Hello, World!")
	HelloGolang()
}