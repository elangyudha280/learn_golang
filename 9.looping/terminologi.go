

//! MATERI LOOPING DI GOLANG

//! 1. APA ITU LOOPING ATAU PERULANGAN
/*
	PERULANGAN ADALAH SALAH SATU ALGORITMA UNTUK MENJALAKAN KODE SECARA BERULANG-ULANG
*/

//! 2. CARA MENGGUNAKAN LOOPING DI GOLANG
/*

	todo: NOTE: break adalah sebuah keyword untuk menghentikan sebuah looping

	todo: NOTE: continue sebuah keyword untuk menghentikan loop nya dan lanjut ke count looping selanjutnya
	* ada beberapa cara

	! 1.
	/
		for condition{
			program
		}
		*\ contoh
		var nilai_awal = 0;
		for nilai_awal < 10 {
			fmt.printL('print')
			nilai_awal++
		} //todo: -> ini artinya kita akan loop nilai awal dengan kondisi jika masih di bawah 10 jika diatas atau sama dengan 10 maka hentikan loopnya
	/

	! 2. for dengan statement langsung
	/
		for statement{
			program
		}
		* contoh
		for var_nilai_awal = 0; var_nilai_awal < 10; var_nilai_awal++ {
			fmt.printL('print')
		} //todo: -> ini artinya kita akan loop nilai awal dengan kondisi jika masih di bawah 10 jika diatas atau sama dengan 10 maka hentikan loopnya
	/

	! 3. for range || untuk meloop element collection seperti array,slice dan map
	/
		for nama_var_index, nama_variable_value := range nama_data_collection {
			program
			fmt.printl(nama_var_index,nama_variable_value)
		}
		* contoh
		names := []string["name1","name2",...]
		for index, name := range names{
			fmt.printLn(index,name)
		} // todo: -> ini artinya kita akan meloop element yg ada di slice akan berenti ketika sudah di element terakhir slicenya

	/


*/