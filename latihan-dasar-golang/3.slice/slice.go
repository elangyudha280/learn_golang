package main

func main() {

	//todo: MEMBUAT SLICE DARI DATA ARRAY YG SUDAH ADA
	// data mahasiswa
	// mahasiswa := [...]string{"yudha", "raka", "helmi", "pram", "reja", "nanda", "rehan"}

	// // data slice dari data array
	// kelompok1 := mahasiswa[0:3] //todo: -> ini artinya kita akan nge-slice atau mengambil data array dari index pertama sampai index ke 3
	// kelompok2 := mahasiswa[3:]  //todo: -> ini artinya kita akan nge-slice atau mengambil data array dari index ke 3 sampai index terakhir

	// fmt.Println("data parent array ->", mahasiswa)
	// fmt.Println("slice dari array ->", kelompok1)
	// fmt.Println("slice dari  array ->", kelompok2)

	// // len slice
	// fmt.Println("jumlah data yg ada di slice kelompok1 ->", len(kelompok1))

	// // kapasisitas yg ada di slice kelompok 1
	// fmt.Println("kapasitas yg ada di slice kelompok1 ->", cap(kelompok1))

	// kelompok_GABUNGAN := append(kelompok1, "user3", "user4", "user5")
	// fmt.Println(kelompok_GABUNGAN)
	// fmt.Println(kelompok1)
	// fmt.Println(kelompok2)
	// fmt.Println(mahasiswa)

	//todo: MEMBUAT SLICE LANGSUNG
	// dosen := []string{
	// 	"dosen1",
	// 	"dosen2",
	// 	"dosen3",
	// }

	//todo membuat slicde dengan fungsi make
	// mahasiswa := make([]string, 4)
	// mahasiswa[0] = "mhs1"
	// mahasiswa[1] = "mhs2"
	// mahasiswa[2] = "mhs3"
	// mahasiswa[3] = "mhs4"
	// fmt.Println(cap(mahasiswa))
}
