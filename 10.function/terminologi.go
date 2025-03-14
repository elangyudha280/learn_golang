


//! MATERI FUNCTION DI GOLANG

//! 1. APA ITU FUNCTION
/*
	function adalah sebuah sub program atau blok kode untuk membuat program yg dapat digunakan secara berulang-ulang
*/ //todo: jadi misal ketika kita mempunyai program yg dapat digunakan secara berulang kita bisa buat menjadi sebuah fungsi

// TODO: NOTE: SEBUAH FUNGSI BISA DIBUAT DILUAR FUNCTION MAIN TETAPI SAAT MEMANGGILNYA HARUS DI DALAM FUNCTION MAIN

//! 2. KENAPA MENGGUNAKAN FUNCTION
/*
*  1. DAPAT DIGUNAKAN UNTK MEMBUAT PROGRAM YG DAPAT DIGUNAKAN SECARA BERULANG ULANG
 */

//! 3. CARA DAN ATURAN MEMBUAT FUNCTION DI GOLANG
/*
	todo: NOTE: PARAMETER ADALAH SEBUAH NILAI ATAU VARIABLE YG AKAN DI OLAH ATAU DIGUNAKAN DI DALAM FUNGSI
	TODO: NOTE: ARGUMENT ADALAH SEBUAH NILAI YG DIKIRIMKAN KE SEBUAH FUNGSI SAAT FUNGSINYA DI PANGGIL


	! 3.1 BASIC FUNCTION
	/
		TODO: CARA PALING SEDERHANA MEMBUAT FUNGSI
		* func nama_function(){
		*	...kode program
		*}
		todo: cara menggunakannya
		* nama_function()
	/

	! 3.2 FUNCTION DENGAN PARAMETER
	/
		TODO: CARA MEMBUAT FUNGSI DENGAN PARAMETER
		* func nama_function(parameter1 tipedata, parameter2 tipedata){
        *    ...kode program
        *}
		todo: cara menggunakannya
		* nama_function(value1, value2)

		||
		todo: NOTE: JIKA FUNCTION DENGAN PARAMETER MENGUBAH ISINYA MENGGUNAKAN KONSTANTA MENJADI PARAMETER SAMA
		* func nama_function(parameter1, parameter2 int){
        *    parameter1++
        *    parameter2++
        *    fmt.Println(parameter1, parameter2)
        *}
	/

	! 3.3 FUNCTION DENGAN RETURN VALUE
	/
		todo: cara membuat fungsi yg mengembalikan nilai

		todoL return value 1 nilai
		* func nama_function() tipeData {
		*	return value
		*}
		* nilai := nama_function()

		todo: return multiple nilai
		* func nama_function() (tipeData1, tipeData2) {
        *    return value1, value2
        *}
		* value1, value2,...value(sesuaikan dengan return value di fungsinya) := nama_function()
	/



*/