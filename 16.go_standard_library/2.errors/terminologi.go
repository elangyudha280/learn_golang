

//! MATERI STANDARD LIBRARY ERRORS DI GOLANG


//! 1. APA ITU ERRORS
/*
	* ERRORS ADALAH SALAH SATU PACKAGE GO YANG DIGUNAKAN UNTUK MENANGANI ERROR ATAU KESALAHAN PADA SAAT PROGRAM DIJALANKAN.
*/
//todo: notes: SAAT MENGGUNAKAN PKG ERRORS PROGRAM TIDAK AKAN LANGSUNG BERHENTI JIKA TERJADI ERROR JADI GA KAYA FUNGSI PANIC


//! 2. CARA MENGGUNAKAN PACKAGE ERRORS
/*
	TODO: IMPORT DULU PKG ERRORNYA import "errors"
	* syntax
	* errors.nama_method()

	* contoh method yg biasanya digunakan
	* errors.New() = membuat error baru dengan pesan tertentu
	* errors.Is() = mengecek apakah error tertentu sama dengan error yg diinginkan
	* errors.As() = mengecek apakah error tertentu bisa di konversi ke tipe error yg diinginkan
	* errors.Unwrap() = mengambil error yg di bungkus oleh error lain
	* errors.Join() = menggabungkan beberapa error menjadi satu error
	
TODO: NOTES: LEBIH LENGKAPNYA BISA CEK DI DOKUMENTASINYA https://pkg.go.dev/errors

	* contoh penggunaan errors.New()
	
	*	err := errors.New("ini adalah error baru")
	*	if err != nil {
	*		fmt.Println(err.Error())
	*	}

	*/
*/