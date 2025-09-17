

//! MATERI PACKAGE AND IMPORT DI GOLANG


//! 1. APA ITU PACKAGE
/*
	? PACKA ADALAH SEBUAH FOLDER ATAU TEMPAT YANG BERISI SEKUMPULAN MODULE
*/
//TODO: NOTES: MODULE ADALAH SEBUAH FILE YANG BERISIKAN CODE UNTUK MEMBUAT PROGRAMS


//! 2. APA ITU IMPORT
/*
	? IMPORT ADALAH SEBUAH CARA UNTUK KITA MENGGUNAKAN SEBUAH MODULE DI PROJECT KITA
*/
//TODO:NOTES: SECARA DEFAULT GOLANG BISA MENGAKSES DATA ATAU CODE YG BERADA DI FILE LAIN SELAMA MASIH DALAM 1 FOLDER YG SAMA, JIKA BERADA DI FOLDER YG BERBEDA HARUS DI IMPORT TERLEBIH DAHULU

//! 3. CARA MEMBUAT PACKAGE
/*
	TODO: KITA BUAT FOLDER BARU DENGAN NAMA APAPUN
	TODO: TERUS FILE YG ADA DI DALAM TERSEBUT BISA DIANGGAP SEBAGAI MODULE
	TODO: di dalam file nanti kita harus bikin package nama_package(sesuai dengan nama folder yg dibuat)

	todo: cara import -> import "...url_path/name_package" || import (
	todo:	"...url_path/name_package"
	todo: )

	todo: notes: lingkup url_pathnya di mulai dari nama folder projectnya
*/


//!4. APA ITU access modifier
/*
	? access modifier adalah sebuah cara untuk mengatur apakah sebuah fungsi, variabel, atau tipe data di dalam package bisa di akses dari luar package tersebut
*/

//!5. cara menggunakan access modifier
/*
	* Untuk membuat sebuah fungsi, variabel, atau tipe data dapat diakses dari luar package, kita harus menuliskan nama dengan huruf kapital diawal
	
	* Sebaliknya, jika kita ingin membatasi aksesnya, kita cukup menuliskan dengan huruf kecil diawal
*/
