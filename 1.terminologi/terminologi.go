
//! MATERI TERMINOLOGI GOLANG

//! 1. APA ITU GOLANG
//TODO: GOLANG ADALAH BAHASA PEMOGRAMAN YG DIBUAT DENGAN BAHASA C OLEH GOOGLE PADA TAHUN 2009 UNTUK MEMBUAT PERANGKAT LUNAK YANG CEPAT, EFISIEN DAN MUDAH DI PAHAMI

///TODO: NOTES BAHASA GOLANG ADALAH BAHASA YG STRONGLY TYPE
// TODO: NOTES GOLANG BAHASA PERMOGRAMAN YG SENSITIVE CASE
//! 1.2 KELEBIHAN GOLANG
/*
* Sangat cocok untuk aplikasi berbasis jaringan (seperti web server atau API).
* Kode yang sederhana, mudah dibaca, dan dirawat.
* Komunitas yang berkembang pesat dan dukungan yang kuat dari Google.
* Digunakan oleh perusahaan besar seperti Google, Uber, Dropbox, dan Netflix.
 */

//! 1.3 KEKURANGAN GOLANG
/*

* Tidak memiliki fitur-fitur canggih seperti generics (meskipun sudah ditambahkan di Go 1.18).
* Tidak mendukung pemrograman berorientasi objek secara penuh.
* Kurang fleksibel untuk pengembangan aplikasi frontend atau desktop.

 */
//! 2. CARA INSTAL GOLANG
/*
	LEBIH LENGKAP BISA CEK DI DOKUMENTASI WEB GOLANGNNYA
*/

//! 3. LANGKAH" ATAU CARA MEMBUAT PROJECT DENGAN GOLANG
/*

	* 1. MASUK KE DALAM SEBUAH DIREKTORI
	* 2. KETIKAN PERINTAH  = go mod init nama-projectnya

	TODO: SAAT KITA MENINISIALISASIKAN PROJECT PERTAMA KALI KITA AKAN DAPA FILE go.mod  yang digunakan untuk mengelola module dalam proyek Golang

	TODO: SAAT KITA INGIN MEMBUAT FILE GO ITU EKTENSI ATAU FORMAT NYA = nama_file.go

*/

//! 4. MAIN FUNCTION
/*
	TODO: MAIN FUNCTION ADLAHA FUNGSI UTAMA YG AKAN DI JALANKAN KETIKA PROGRAM DIJALANKAN
	* jadi fungsi ini digunakan untuk meng ekse atau menjalakna code program dari module yg kita buat

	TODO: NOTE main function harus berada di dalam main package atau file utama\
	todo: NOTE: tanda titik koma di golang tidak wajib

	TODO: NOTES: UNTUK BUILD CODE GOLANGYA BISA MENGGUNAKAN PERINTAH go build
	TODO: untuk  menjalakan file compile codenya bisa gunakan path/nama_file_hasil_compilenya.exe || ./latihan-dasar-golang.exe
	TODO: NOTES: UNTUK MENJALAKAN CODE GOLANG TANPA COMPILE KITA BISA GUNAKAN PERINTAH go run nama_filenya.go

	todo: NOTES:BEST PRACTICENYA FILE MAIN ATAU MAIN FUNCTION HANYA BOLEH ADA SATU DI PROJECT YG KITA BUAT DENGAN GOLANG
	todo: sebenarnya bisa membuat lebih dari satu file main atau main function asalkan berbeda folder tapi tidak di sarankan

	* CONTOH SYNTAX
	package main -> unutk menandakan bahwa filnya akan menjadi file utama

	func main() {
		...CODE PROGRAMS
	}
*/