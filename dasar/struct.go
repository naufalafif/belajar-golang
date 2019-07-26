package main

import "fmt"

type orang struct {
	namaDepan   string
	namaBelakag string
}

func main() {
	// CARA ASSING VALUE KE Struct

	// Cara 1 // namaDepan = Budi, namaBelakang = Ananta // Mengikuti Urutan Struct // Not Recommended
	budi := orang{"Budi", "Ananta"}

	// Cara 2 // namaDepan = Budi, namaBelakang = Ananta // Mengikuti Key Struct // Recommended
	budi2 := orang{namaDepan: "Budi", namaBelakag: "Ananta"} // namaDepan = Budi, namaBelakang = Ananta

	// Cara 3 // namaDepan = Budi, namaBelakang = Ananta // Mengikuti Key Struct Setelah Di Assign // Recommended
	var budi3 orang
	budi3.namaDepan = "Budi"
	budi3.namaBelakag = "Ananta"

	fmt.Println("Budi", budi)
	fmt.Println("Budi2", budi2)
	fmt.Println("Budi3", budi3)

	// Cara Membaca Isi Struct by Key
	fmt.Println("Nama Depan = ", budi.namaDepan)
	fmt.Println("Nama Belakang = ", budi.namaBelakag)
}
