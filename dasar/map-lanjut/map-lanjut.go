package main

import "fmt"

func main() {
	warna := map[string]string{
		"merah": "f111111",
		"hijau": "f222222",
	}
	cetakMap(warna)
}

func cetakMap(dataMap map[string]string) {
	for warna, kode := range dataMap {
		fmt.Println("Kode untuk warna ", warna, " adalah ", kode)
	}
}
