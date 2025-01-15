

//! MATERI VARIABLE DI GOLANG

//! 1. APA ITU VARIABLE
/*
	VARIABLE ADALAH SEBUAH WADAH ATAU TEMPAT UNTUK MENYIMPAN SEBUAH NILAI
*/
// TODO: JADI KETIKA KITA INGIN MENGGUNAKAN NILAI TERSEBUT KITA BISA LANGSUNG PANGGIL NAMA VARIBLENYA

//! 2. CARA MEMBUAT VARIBLE DI GOLANG
/*

* ADA BEBERAPA CARA
todo: syntax = keywword tipedata = value
todo: var nama string = 'value'
todo: shor declaration nama_var := value

* 1. var = sebuah keyword yg digunakan untuk membuat variable yg nilainya bisa diubah
* 2. const = sebuah keywword yg digunakan untuk memebuat variable yg nilainya tidak bisa diubah
* 3. short declaration || := = sebuah cara untuk medeklarasikan variable dan harus di isi nilainya dan juga cara ini hanya bisa digunakan di dalam sebuah fungsi
* 4. multiple deklarasi ||
* var (
	* nama_var1 = value1
	* ...
* ) = sebuah cara untuk mendeklarasikan lebih dari satu variable secara bersamaan dan variable harus digunakan
*/

//!3. type declaration
/*
	todo: type declaration adalah sebuah cara untuk membuat alias atau nama lain dari sebuah tipedata atau aturan yg kita buat

	* contoh : misal kita ingin menentukan struktur object agar lebih rapi kita bisa menggunakan type decralaration sehingga nanti ketika kita ingin menggunakan aturan atau tipe data tersebut kita bisa panggil nama type declarationnya

	* styntax
	todo: type nama_alias = aturan atau tipe data

	* contoh
	type nama_user = string; -> ini artinya kita membuat alias nama_User yg aturannya harus bertipe data string

	var nama nama_user = "user1" -> ini artinya kita menggunakan alias nama user yg artinya variable ini harus bertipe data string
*/