package main

// func ubahValue(value *int)  {
// 	*value += 10
// }

//! function memebuat custom pointer
// func buatPointer(value int) *int{
// 	return &value
// }

type Person struct {
	name string;
	umur int;
}

func (P *Person) tambahUmur(){
	P.umur += 1
}
func (P *Person) cekUlangTahun() int{
	return P.umur
}
func main() {
	// //! simple pointer
	// var x = 10
	// var p *int = &x //* -> ini artinya variable p akan menjadi pointer dari varible x sehingga ketika kita ubah data p maka varible x juga akan berubah
	// //! ubah variable x lewat pointer p
	// *p= 20
	// fmt.Printf("variable x is %v and variable p as pointer is %v", x, *p)

	//! pointer di parameter function
	// var nilaiAwal int = 5
	// fmt.Printf("value nilai awal belum diubah %v \n",nilaiAwal)
	// ubahValue(&nilaiAwal)
	// fmt.Printf("value nilai awal sudah diubah %v \n",nilaiAwal)

	// var customPointer *int = buatPointer(10)
	// *customPointer +=10
	// fmt.Printf("%v",*customPointer)

	//! POINTER DI STRUCT
	// var user1 Person = Person{
	// 	name: "user1",
	// 	umur: 20,
	// }

	//! ubah umur user1
	// user1.tambahUmur()
	// fmt.Printf("umur user %v \n",user1.cekUlangTahun())
	// fmt.Printf("saya user %v umur %v",user1.name,user1.umur)

	//! pointer di slice
	// var peSlice = new(int)
	// var peSlice2 = new(int)
	// var testSlice = []*int{peSlice,peSlice2}
	// *testSlice[0] = 1
	// *testSlice[1] = 12

	// fmt.Printf("slice value %v, slice pointer %v \n", *testSlice[0],testSlice)

	// for index,value := range testSlice {
	// 	fmt.Printf("value testsclie index ke-%v adalah %v \n",index,value)
	// }
}

