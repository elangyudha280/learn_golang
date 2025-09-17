

//! materi pointer di golang


//! 1. apa itu pointer
/*
	* Pointer adalah tipe data yang digunakan untuk menyimpan alamat memori dari sebuah variabel ketika digunakan di tempat lain, entah itu varible lain,parameter pada function atau method pada struct
*/
//todo: jadi dengan pointer kita bisa menunjuk alamat memori dari data pada sebuah variable

//! 2. KAPAN PAKAI POINTER
/*
	* Kalau mau hemat memori (tidak copy data besar).
	* Kalau mau ubah nilai asli dari variabel lewat function/method.
	* Kalau butuh referensi ke objek yang sama.

*/

//! 3. operator pointer
/*
	* 1. & (address-of operator) → mengambil atau mendaftarkan alamat memori dari sebuah variabel. -> jadi ini ambil alamat tempat atau objectnya
	* 2. * (dereference operator)→ mengambil nilai dari alamat memori yang ditunjuk oleh pointer. atau tanda & -> ini untuk ambil nilai dari tempat atau objectnya
	

	* function new(value) = >function dari golang untuk membuat pointer kita sendiri

	* gampangnya
	* & -> membuat alamat atau buat pointer
	* * -> buka alamat (pakai/ubah nilai di balik pointer).


	todo: tipe data yg bertipe passing by value -> artinya yg dikirim nilainya bukan alamat memorinya langsung
	/
		* int , string, boolean,struct, array
	/

	todo: tipe data yg bertipe passing by reference -> artinya yg dikirim nilainya alamat memorinya langsung
	/
		* pointer, slice, map, chanel, function, interface 
	/
	*/


//! 4. CARA MENGGUNAKAN POINTER
/*
	TODO: POINTER BISA DI PAKE KE SEBUAH VARIABLE, PARAMETER FUNCTION, METHOD DI STRUCT

	* VARIABLE
	/
		* x := 10
		* p := &x -> INI ARTINYA KITA AKAN MENYIMPAN ALAMAT MEMORI DARI VARIABLE x KE p dan variable p otomatis akan menjadi pointer
		* fmt.Println(*p) -> INI ARTINYA KITA AKAN MENGAMBIL NILAI DARI ALAMAT MEMORI YANG DITUNJUK atau pointer p YAITU NILAI DARI VARIABLE X
	/

	* FUNCTION 
	/
		* func ubahNama(u *User) { -> INI ARTINYA KITA MENGIRIMKAN REFERENSI DARI OBJEK ATAU STRUCT User KE FUNCTION INI
		*	u.Name = "Budi"
		* }
		*  u := User{Name: "Andi"}
    	*  ubahNama(&u) -> INI ARTINYA KITA AKAN MENDAFTARKAN ALAMAT atau membuat pointer DARI VARIABLE u ke dalam parameter function ubah nama
	/

	//* STUCT METHOD
	/
		* type Counter struct {
		*	Value int
		* }

		* func (c *Counter) Increment() {
		* 	c.Value++
		* }
	/
*/