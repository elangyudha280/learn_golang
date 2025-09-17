

//! MATERI TIPE DATA DI GOLANG

//! 1. APA ITU TIPE DATA
/*
	TIPE DATA ATAU DATAYPE ADALAH SEBUAH JENIS DATA YG KITA BERIKAN DI PROGRAMN UNTUK MEMBERITAHUKAN KOMPUTER DATA APA YG AKAN KITA BUAT
*/

///todo: NOTES: SETIAP TIPE DATA MEMILIKI NILAI DEFAULTNYA MASING"

//! 2. MACAM-MACAM TIPE DATA DI GOLANG
// * ADA BEBERAPA TIPE DATA

//! 2.1 NUMBER
/*
	TODO: TIPE DATA UNTUK MEMBUAT ANGKA

	* INT = TIPE DATA UNTUK ANGKA BILANGAN BULAT (123)
	* FLOAT = TIPE DATA UNTUK ANGKA BILANGAN PECAHAN (1.2)
*/

//! 2.2 BOOLEAN
/*
	TODO: TIPE DATA UNTUK MENENTUKAN KEBENARAN

	* TRUE = TIPE DATA UNTUK BENAR
	* FALSE = TIPE DATA UNTUK SALAH
*/

//! 2.3 String
/*
	TODO: TIPE DATA UNTUK MEMBUAT KARAKTER
	TODO: CARA NYA DENGAN MENGGUNAKAN TANDA PETIK 2 ATAU "KARAKTER"

	TODO: len('karakter') = fungsi untuk menghitung panjang karakter
	TODO: "karater"[index || angka] ++ "karater"[0] = untuk mengambil kata pada karakter

	
	///! method atau fungsi untuk berinteraksi dengan string
	todo: kita import dulu pkg yg namanya strings
	
	* 1. strings.ToLower(string_value) -> untuk membuat string menjadi huruf kecil

	* 2. strings.ToUpper(string_value) -> untuk membuat string menjadi huruf capital

	* 3. strings.HasPrefix(string_value,"kata_yg_ingin_di_Cek") -> untuk mengecek apakah kata awalan string sesuai dengan kata yg kita tentukan

	* 4. strings.Contains(string_value,"kata_yg_ingin_di_Cek") -> untuk mengecek apakah sebuah string mengandung atau ada kata tertentu yg kita tentukan

	* 5. strings.Split(...string_value,"separator||pemisah_katanya") -> untuk memisahkan sebuah string dan nanti menjadi array

	* 6. strings.ReplaceAll(string_value,"kata_yg_ingin_diganti","kata_baru") -> untuk mengubah atau mengganti kata yg ada di string menjadi kata baru
*/

//! CONVERSION TIPE DATA
/*
	TODO: CONVERSION ADALAH SEBUAH CARA UNTUK MENGUBAH TIPE DATA MENJADI TIPE DATA LAIN
	todo: strconv adalah package golang untuk mengubah data string menjadi tipe data lain
	* CARANYA DENGAN MENGUNAKAN METHOD NAMA TIPE DATANYA 

	* CONTOH
	* 1. int("1") -> ini artinya mengubah string 1 menjadi number 1
	* 2. float64(10) -> ini untuk mengubah data int 10 menjadi float64
*/