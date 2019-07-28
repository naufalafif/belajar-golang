package main

import "fmt"

func main() {

	// Cara Deklaasi map ke 1 // deklarasi ekspresi memerlukan make
	contohMap := make(map[string]string)
	contohMap["nama"] = "naufal"
	contohMap["alamat"] = "jogja"

	delete(contohMap, "alamat")

	fmt.Println(contohMap)

	// Cara Deklarasi Map ke 2
	warna := map[string]string{
		"merah": "f21312312",
		"hijau": "f5435345",
	}

	fmt.Println(warna)

	// Cara Deklarasi Map ke 3
	var satuan map[string]string

	fmt.Println(satuan)
}
