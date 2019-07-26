package main

func main() {
	// dekBaru := dekbaru() // Membuat Dek Baru
	// dekBaru.cetak() // Cetak Kartu
	// dekBaru.acak() // Acak Kartu Dalam Dek
	// dekBaru.simpanKeFile("__filedekku") // Simpan Dek Ke bentuk File
	// dekDariFile := dakBaruDariFile("__filedekku") // Baca Dek Dari File
	// kartuku := bagiKartu(dekBaru, 5) // Mengambil 5 Kartu

	dekDariFile := dakBaruDariFile("my_cards")
	dekDariFile.acak()
	dekDariFile.cetak()

}
