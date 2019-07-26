package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Membuat Tipe Dek
// Dek adalah kumpulan text/string dalam array/larik

type dek []string

func (d dek) cetak() {
	for i, kartu := range d {
		fmt.Println(i, kartu)
	}
}

func (d dek) jadiHuruf() string {
	return strings.Join(d, ",")
}

func (d dek) simpanKeFile(namafile string) error {
	error := ioutil.WriteFile(namafile, []byte(d.jadiHuruf()), 0666)
	return error
}

func bagiKartu(d dek, jumlahKartu int) (dek, dek) {
	return d[:jumlahKartu], d[jumlahKartu:]
}

func dekbaru() dek {
	jenisKartu := []string{"Spades", "Diamonds", "Hearts", "Club"}
	angkaKartu := []string{"Ace", "Two", "Three", "Four"}
	kartuKartu := dek{}

	for _, jenis := range jenisKartu {
		for _, angka := range angkaKartu {
			kartuKartu = append(kartuKartu, angka+` `+jenis)
		}
	}

	return kartuKartu
}

func dakBaruDariFile(namafile string) dek {
	bs, err := ioutil.ReadFile(namafile)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	dekDalamSatuString := string(bs)
	dekBerbentukArrayString := strings.Split(dekDalamSatuString, ",")
	kartuKartu := dek(dekBerbentukArrayString)
	return kartuKartu
}

// Catatan, untuk menghasilkan data yang selalu acak
// kita perlu membuat randomGenerator yang menggunakan sumber/sumberdata yang berbeda setiap saat
// terdapat tiga tahap yaitu : - Membuat Bibit/ Angka untuk diolah, - Membuat Source/ Sumber dari bibit, - Membuat Generator dari sumber

func (d dek) acak() dek {
	bibit := time.Now().UnixNano()  // Seed
	sumber := rand.NewSource(bibit) // Create Random Source
	r := rand.New(sumber)

	for i := range d {
		posisiBaru := r.Intn(len(d) - 1)
		d[i], d[posisiBaru] = d[posisiBaru], d[i]
	}
	return d
}
