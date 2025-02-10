package main

func main() {

	//! infinite loop
	// for {
	// 	fmt.Println("Hello, World!")
	// }

	//! loop like a while
	// nilai_awal := 1

	// for nilai_awal <= 10 {
	// 	nilai_awal++
	// 	if nilai_awal%2 == 1 {
	// 		continue
	// 	}
	// 	fmt.Println("Hello, World ke-", nilai_awal, "!")
	// }

	//! loop  dengan statement langsung
	// for nilai_awal := 1; nilai_awal <= 10; nilai_awal++ {
	// 	if nilai_awal <= 5 {
	// 		fmt.Println("angkot ke - ", nilai_awal, "sedang beroperasi")
	// 		continue
	// 	} else if nilai_awal == 6 {
	// 		fmt.Println("angkot ke - ", nilai_awal, "sedang di")
	// 		break
	// 	}
	// 	fmt.Println("angkot ke - ", nilai_awal, "sedang tidak beroperasi")
	// }

	//! loop data collection

	//! loop data array
	// mahasiswa := [...]string{"mahasiswa1", "mahasiswa2", "mahasiswa3", "mahasiswa4"}
	// for index, value := range mahasiswa {
	// 	fmt.Println("mahasiswa ke -", index+1, "adalah", value)
	// }

	//! LOOP DATA SLICE
	// dosen := []string{"dosen1", "dosen2", "dosen3", "dosen4"}
	// for index, value := range dosen {
	// 	fmt.Println("dosen ke -", index+1, "adalah", value)
	// }

	//! loop data map
	// person := map[string]string{
	// 	"nama": "John Doe",
	// 	"umur": "25",
	// }
	// for key, value := range person {
	// 	fmt.Println(key, ":", value)
	// }

	// segitiga
	// for i := 0; i <= 10; i++ {
	// 	for j := 0; j <= i; j++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// }

}
