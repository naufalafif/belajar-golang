package main

import (
	"os"
	"testing"
)

func TestDekBaru(t *testing.T) {
	dek := dekbaru()

	if len(dek) != 16 {
		t.Errorf("Seharusnya Panjang Dek berjumlah 16, tetapi panjang dek adalah %v", len(dek))
	}

	if dek[0] != "Ace Spades" {
		t.Errorf("Seharusnya Kartu Pertama adalah `Ace Spades`, tetapi kartu pertama adalah %v", dek[0])
	}

	if dek[len(dek)-1] != "Four Club" {
		t.Errorf("Seharusnya Kartu Terakhir adalah `Four Club`, tetapi kartu terakhir adalah %v", dek[len(dek)-1])
	}
}

func TestSimpanDekdanBacaDekDariFile(t *testing.T) {
	namaFileDek := "__dekTest"
	os.Remove(namaFileDek)

	dekUntukDisimpan := dekbaru()
	dekUntukDisimpan.simpanKeFile(namaFileDek)

	dekDariFile := dakBaruDariFile(namaFileDek)

	if len(dekDariFile) != 16 {
		t.Errorf("Seharusnya Panjang Dek berjumlah 16, tetapi panjang dek adalah %v", len(dekDariFile))
	}
	os.Remove(namaFileDek)
}
