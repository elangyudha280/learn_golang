

//! MATERI INTERFACE DI GOLANG

//? 1. APA ITU INTERFACE
/*
	INTERFACE ADALAH SEBUAH TEKNIK UNTUK MEMBUAT ATURAN ATAU KONTRAK PADA  SEBUAH OBJECT CONTONYA STRUCT OBJECT
*/
//TODO: NOTES: JADI KITA MENENTUKAN METHOD APA AJA YG HARUS ADA DI OBJECT STRUCT ATAU VARIABLE

//? 2. CARA MEMBUAT DAN MENGGUNAKAN INTERFACE
/*
	TODO: NOTES: INTERFACE DI GOLANG HANYA BISA MEMBUAT ATURAN METHOD APA SAJA YG HARUS ADA DI OBJECT STRUCTNYA
	todo: Jika sebuah struct memilki semua method yang ada di dalam interface, maka struct tersebut otomatis dianggap menerapkan interface tersebut.
	TODOl: NOTES: MEMBUAT ITNERFACE ITU DENGAN MENGGUNAKAN TYPE DECLARATIONS

	* * * * * * * * * * * * * * * * *
	/
		type nama_interface interface{
			nama_method() tipe_data_return_value || kosong
		}


		* cara menggunakan interface
		? 1. interface di params function
        * type nama_struct struct{}
		* func (nama_struct nama_struct) nama_method_sesuai_di_interfacenya(){
        *    ...programs
        * }

		* func nama_function(nama_param nama_interface){
        *    nama_param.nama_method()
        * }
			* nama_variable := nama_struct{}
			* nama_variable.nama_method()
	/
*/

//? 4. interface kosong
/*
	interface kosong adalah sebuah teknik untuk menentukan sebuah object bisa bernilai dengan tipe data apapun
*/
//todo: contohnya kaya object variable,parameter function atau property di struct jika di set aturannya interaface {} atau any itu bisa bernialai dengan tipe data apapun

//todo: jadi kaya any di typescript

//? 4.1 cara menggunakan inteface kosong
/*

	* kasus biasanya menggunakan inteface kosong
	! Saat kita tidak tahu tipe data yang akan diterima.
	todo: Contohnya saat menerima input dari JSON API atau database yang bisa berisi berbagai jenis data.

	! inteface kosong di variable
	* type nama_interface interface{} atau any // -> ini artinya variable ini bisa bernilai dengan tipe data apapun

	! interface kosong di parameter function
	* func nama_function(nama_param interface{}||any interface{}){//-> ini artinya parameter ini bisa bernilai dengan tipe data apapun
        *    nama_param.nama_method()
        * }

	! interface kosong di property di struct
	* type nama_struct struct{
    *    nama_property interface{}||any --> ini artinya property pada struct ini bisa bernilai dengan tipe data apapun
    * }

	! interface kosong di map
	* person := map[string]interface{}||any{
	*    nama:"person1",
    *    alamat: "jalan1",
    *    umur: 20,
    *    //...} --> artinya value dari key pada mapnya bisa bernilai dengan tipe data apapun

	!  interfacec di slice
	* nama_slice []interface{}||any //-> ini artinya slice ini bisa bernilai dengan tipe data apapun


*/