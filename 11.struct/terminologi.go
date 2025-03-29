

//! MATERI STRUCT DI GOLANG

//! 1. APA ITU STRUCT
/*
	STRUCT ADALAH SINGKATAN DARI STRUCTURE sebuah tipe data di Golang yang bisa menyimpan beberapa nilai dan tipe data dalam satu kesatuan.
*/
//todo: jadi hampir mirip seperti object atau Class di JavaScript  tanpa fungsi inheritance di OOP.
//todo: notes jadi bisa dibilang struct ini adalah cetakan atau blueprint dari object yg memiliki property dan juga method

//! 2. kegunaan struct
/*
* ✅ Untuk mengelompokkan beberapa data menjadi satu kesatuan.
* ✅ Bisa digunakan untuk membuat objek yang punya banyak atribut atau property.
 */

//!3 cara menggunakan struct
/*
	TODO: NOTES: STRUCT HARUS DI BUAT DI LUAR FUNCTION MAIN
	TODO:NOTES: STRUCT INI HAMPIR MIRIP DENGAN OBJECT LITERAL PADA JAVSCRIPT YG MEMBEDANKAN DI STRUCT TIDAK MEMILIKI KONSEP SEPERTI INHERITANCE DLL

	todo: notes: kebiasaan membuat nama property di struct itu menggunakan huruf kapitan

	TODO: BASIC
	* type nama_struct struct {
	*  nama_field tipe_data;
	*  || jika property memiliki tipedata yg sama
	*  nama_field1, nama_field2 tipe_data;
	* }

	*  || method di struct
	*  funct (nama_param_struct nama_struct) nama_function(...parameter:tipedata) tipedata(optional){
	*  ...program
	*};

	* cara menggunakannya
	todo: deklarasi dulu baru masukan data
	* var nama_variable nama_struct;
	* nama_variable.nama_property_struct = value;
	* nama_variable.nama_method()

	todo: deklarasi langsung assignmen data dengan menentukan nama key nya
	* nama_variable := nama_struct{property: "value",...property:value}

	todo: deklarasi langsung assignmen data tanpa menentukan nama key nya
	todo: notes: cara ini memberikan nilainya harus sesuai dengan urutan key yg ada di structnya
	* nama_variable := nama_struct{"value",...value}
*/
