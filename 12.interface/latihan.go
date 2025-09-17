package main

import "fmt"

type Kendaran interface {
	Jalan() string
	GetMerk() string
}

type Mobil struct {
	merk string
}

func (m Mobil) Jalan() string {
	return "Mobil " + m.merk + " sedang berjalan"
}
func (m Mobil) GetMerk() string {
	return m.merk
}

type Motor struct {
	merk string
}

func (m Motor) Jalan() string {
	return "Motor " + m.merk + " sedang berjalan"
}
func (m Motor) GetMerk() string {
	return m.merk
}
func main() {

	//!! menggunakan object struct mobil dengan interface
	mobil1 := Mobil{
		merk: "Toyota",
	}

	var motor1 Motor;
	motor1.merk = "Honda";
	
	fmt.Println(mobil1.Jalan(),mobil1.GetMerk())
	fmt.Println(motor1.Jalan(),motor1.GetMerk())

}