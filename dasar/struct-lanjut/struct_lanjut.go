package main

import "fmt"

type kontak struct {
	email   string
	kodePos string
}

type orang struct {
	namaDepan   string
	namaBelakag string
	kontak      kontak
}

func main() {

	budi := orang{
		namaDepan:   "Budi",
		namaBelakag: "Ananta",
		kontak: kontak{
			email:   "naufalafif58@gmail.com",
			kodePos: "94364",
		},
	}

	// Pointer
	budi.ubahNamaDepan("Budi Baru")

	budi.cetak()

}

func (o *orang) ubahNamaDepan(namaDepan string) {
	(*o).namaDepan = namaDepan
}

func (o orang) cetak() {
	fmt.Printf("%+v", o)
}
