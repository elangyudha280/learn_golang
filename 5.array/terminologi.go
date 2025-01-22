

//! MATERI ARRAY DI GOLANG

//! 1. APA ITU ARRAY
/*
 *	ARRAY ADALAH SEBUAH TIPE DATA YANG DIGUNAKAN UNTUK MENAMPUNG BANYAK NILAI YG DI DALAM NYA ADA INDEX DAN ELEMENT
 */
//todo: index = urutan elemen array
//todo: element = nilai yg dimasukan ke dalam array

//! 2. CARA MEMBUAT ARRAY DI GOLANG
/*

	! ADA BEBERAPA CARA DALAM MENDEKLARASIKAN ARRAY DI GOLANG

	* 1. menentukan secara jelas jumlah index dan element yg ada di dalam array
	/
		* var||const nama_var [jumlah_index]tipedata;
		||
		* var||const nama_var  = [jumlah_index]tipedata{
		   * value1,
		   * value2,
           * value3,
		* }
		todo: var angka [3]int; -> ini artinya variable ini akan bertipe data array yg jumlah index elementnya 3 dan tipe datanya harus number || int

		! cara set value
		* angka[0] = 1;
		* angka[1] = 2;
		* angka[2] = 3;
		||
		* var||const angka = [3]int{1, 2, 3}

		! kelebihan
		*1. kita bisa secara jelas menentukan jumlah element atau nilai yg ada di dalam array

		! kekurangan
		*1. kita harus mengetahui jumlah index yang akan diisi array terlebih dahulu
		*2. kita tidak bisa menambahkan nilai lebih dari jumlah yg di tentukan

		todo: note: jika kita memberikan nilai kurang dari jumlah element yg ditentukan, maka sisanya akan menggunakan nilai default,
		todo: contoh klo di awal kita set [3]int maka ketika memberikan nilai kurang dari 3 maka sisanya akan menggunakan nilai defaultnya yaitu klo int adalh 0
	/

	* 2. menentukan secara otomatis jumlah element yg ada di dalam array
	/
		* var||const nama_var  = [...]tipedata{
		   * value1,
		   * value2,
           * value3,
		   * ...
		* }

		todo: var angka [...]int; -> ini artinya variable ini akan bertipe data array yg jumlah elementnya bebas dan tipe datanya harus number || int

		! cara set value
        * var||const angka = [...]int{1, 2, 3}

		! kelebihan
        *1. kita bisa secara otomatis menentukan jumlah element atau nilai yg ada di dalamn array

		! kekurangan
		*1. kita harus meng-assigment atau memberikan nilainya saat membuat arraynya


		! OPERATOR YG BISA DIGUNAKAN UNTUK BERINTERAKSI DENGAN ARRAY
			* 1. len(nama_var) -> mengambil panjang atau jumlah nilai di dlam array
			* 2. nama_var[index] -> mengambil nilai pada index tertentu di array
			* 3. nama_var[index] = value -> mengubah nilai pada index tertentu di array
			* 4. append(nama_var, value) -> menambahkan nilai baru ke array
	/


*/