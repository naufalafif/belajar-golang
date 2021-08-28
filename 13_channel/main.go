package main

import (
	"fmt"
	"net/http"
)

func main() {
	daftarWebsite := []string{
		"https://google.com",
		"https://facebook.com",
		"https://youtube.com",
		"https://golang.org",
	}

	c := make(chan string)

	for _, url := range daftarWebsite {
		go cekLink(url, c)
	}

	// Cara 1 menampilkan hasil channel. print dilakukan 5 kali, karena main routine akan menutup program jika semua line telah tereksekusi
	// fmt.Println(<-c) // Bloking Call
	// fmt.Println(<-c) // Bloking Call
	// fmt.Println(<-c) // Bloking Call
	// fmt.Println(<-c) // Bloking Call
	// fmt.Println(<-c) // Bloking Call

	// Cara 2
	// Cara yang sama dengan cara ke 1, tetapi menggunakan iterasi
	// for i := 1; i <= len(daftarWebsite); i++ {
	// 	fmt.Println(<-c)
	// }

	// Cara ke 3
	for l := range c {
		fmt.Println(l)
	}

}

func cekLink(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "sedang down")
		c <- url
		return
	}
	fmt.Println(url, "jalan normal")
	c <- "jalan normal"
}
