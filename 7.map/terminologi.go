

//! MATERI MAP DI GOLANG

//!! 1. APA ITU MAP
/*
	MAP ADALAH SEBUAH TIPE DATA YG DIGUNAKAN UNTUK MENAMPUNG BANYAK NILAI DAN INDEXNYA BISA BERUPA NAMA ATAU ANGKA
*/

//TODO: NOTES JADI KETIKA KITA INGIN AKSES ELEMENTNNYA ITU BISA MENGGUNAKAN NAMA INDEX ATAU KEYNYA

//TODO:NOTES: JADI INI MIRIP SEPERTI ASSOCIATIVE ARRAY DIPHP ATAU OBJECT DI JAVASCRIPT

//! 2. CARA MENGGUNAKAN MAP DI GOLA
/*

	todo: RUMUS
	* keyword_var nama_var = map[typedata_key_atau_indexnya]typeDataValue ={
	* nama_key:value,
	* ...nama_key:value
	*}

	* 2.1 cara akses datanya
	* nama_Var["nama_key"]


	todo: CONTOH
	* person := map[string]string {
	*	nama:"person1"
	* } // todo: -> ini artinya kita akan membuat data map untuk menampung data person yg tipedata keynya string dan value keynya string


	todo: function dan teknik yg bisa digunakan untuk berinteraksi dengan data map

	* len(var_map) = untuk mengetahui panjang atau jumlah data map
	* map[key] = untuk mengakses data mapnya
	* map[key] = value = untuk mengubah data pada map dengan data baru
	* make(map[typeDataKey]TypeDataValue) = unutk membuat map baru
	* delete(var_map,nama_key) = untuk menghapus key pada map
	* value, nama_var_exist := namaMap["key"] = untuk mengecek apkah key ada di map
*/