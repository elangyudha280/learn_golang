


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

	! VARIADIC FUNCTION
	TODO:Variadic function adalah fungsi yang bisa menerima jumlah argumen yang tidak terbatas .
	TODO: NOTES: JADI HAMPIR MIRIP DENGAN REST PARAMETER DI JAVASCRIPT
	todo: notes: Tidak bisa memiliki lebih dari satu variadic parameter dalam satu fungsi.
	/
		* func nama_function(nama_paramter ...tipedata){
			* paramater varidic function tipe datanya slice
		*}

		* nama_function(valu1,value2,...value)
		||
		data_slice := int{1,2,3,4}
		* nama_function(data_slice...)
	/


	!! FUNCTION AS VALUE
	TODO: function yg disimpan sebagai value di variable
	todo:note jadi ketika kita ingin memanggil functionnya kita panggil nama variablenya
	* func nama_function(...paramter tipe_Data) koson | tipe_data {
	* ...programs
	* }

	* nama_var := nama_function;
	* nama_var(...arguments)

	! FUNCTION AS PARAMETER || CALLBACK
	TODO: CALLBACK ADALAH SEBUAH FUNGSI YG DI BERADA DI PARAMETER ATAU DIKIRIM SEBAGAI ARGUMENTS
	* func nama_function(nama_parameter func(tipe_Data) tipe_data_return_value | 'kosong'){
	*	...program
	* contoh: nama_parameter_as_function(...arguments)
	* }

	! ANONYMOUS FUNCTION || FUNCTION TANPA NAMA
	TODO: ANONYMOUS FUNCTION ADALAH FUNCTION TANPA YG DIBUAT TANPA MENGGUNAKAN NAMA
	TODO: NOTES: BIASANYA ANONYMOUS FUNCTION DIBUAT DI DALAM VARIABLE ATAU DI GUNAKAN LANSUNG SEBAGAI CALLBACK
	* nama_variable := func nama_function(...parameter) tipe_data {
	* ...programs
	* }
	||
	* func nama_function(nama_params_func_as_callback func(tipedata) tipe_Data {
	* ...programs
	* contoh : nama_params_func_as_callback(value)
	* }
	* nama_function((value){
	* ...programs
	*})

*/