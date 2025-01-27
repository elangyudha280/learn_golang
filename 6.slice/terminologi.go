

//! MATERI SLICE DI GOLANG

//! 1. APA ITU SLICE
/*
	SLICE ADALAH TIPE DATA UNTUK MENAMPUNG BANYAK NILAI  ATAU ELEMENT  DENGAN UKURAN YG DINAMIS
*/

//todo: jadi slice sama seperti array cmn yg membedakan adalah ukuran slice bisa dinamis sedangkan array statis atau di tentukan di awal saat tahap deklarasi arraynya

//TODO: NOTES: SLICE DAN ARRAY INI SALING TERKONEKSI JADI NANTI DATA YG ADA DI SLICE ITU BISA AJA MENGGUNAKAN DATA YG ADA DI ARRAY

//! 2. cara menggunakan slice
/*

	TODO: NOTES: DI SLICE ADA 3 DATA ATAU INFORMASI PENTING
	*
		* 1. pointer = penunjuk data awal dari array untuk slicenya
		* 2. length = panjang data yg ada di slicenya
		* 3. capacity = kapasitas atau jumlah data slicenya yg dimulai dari pointer arraynya
	*

	* CARA BUAT SLICE LANGSUNG
	* rumus : keyword_var nama_var = []tipedata{...value}
	||
	* rumus :  nama_var := []tipedata{...value}
	||
	todo: menggunakan fungsi make
	* rumus :  nama_var := make([]tipedata,pangjang_data,kapasitas)


*/

//! 3. cara membuat slice yg datanya dari array
/*

	todo: pertama kita harus membuat arraynya dulu yg akan di jadikan referensi datanya

	* nama_var_Array := [...]tipedata{'value1','value2','value3'}
	todo: rumus nge-slice yg datanya dari array =  var_arr[index_awal:index_akhir]

	* 1. var_arr[:] = untuk nge-slice data array dimulai dari index awal sampai index akhir

	* contoh:
	* var_arr[1:4] -> ini artinya mengambil atau nge-slice data array dari index 1 sampai index ke 4



*/

/*

	fungsi yg bisa digunakan

	* len(data_slice) = unutk mengetahui panjang data slicenya
	* cap(data_slicee) = untuk mengetahui kapasitas slicenya
	* append(data_Slice1,data_slice2) = untuk menambahkan data ke slice
	TODO: JIKA KITA MENAMBHAKAN DATA DENGAN APPEND JIKA DATA YG DI TAMBAHKAN LEBIH DARI KAPASITAS SLICENYA MAKA AKAN MEMBUAT ARRAY BARU JADI TIDAK MENGGUNAKAN ARRAY YG SEBELUMNNYA LAGI
*/