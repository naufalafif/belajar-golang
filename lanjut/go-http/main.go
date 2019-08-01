package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// Writter Buatan sendiri
type logWriter struct{}

func main() {
	resp, err := http.Get("https://www.google.com")

	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}

	// // Cara Manual
	// // -- Membuat Variable untuk menampung byte dari response body
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)      // mengisi data bs dengan response body
	// fmt.Println(string(bs)) // mengubah byte ke string

	// // Cara ke 2
	// // -- fungsi io.copy digunakan untuk untuk menuliskan/menampilkan data berupa reader ke output yang di inginkan
	// // kode dibawah menggunakan os.stdout/terminal sebagai output/writter dan resp.body sebagai output
	// io.Copy(os.Stdout, resp.Body)

	// Cara ke 3
	// Menggunakan Writer - buatan
	lg := logWriter{}
	io.Copy(lg, resp.Body)
}

func (logWriter) Write(p []byte) (n int, err error) {
	fmt.Println(string(p))

	return len(p), nil
}
