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

	for _, url := range daftarWebsite {
		go cekLink(url)
	}

	fmt.Scanln()
}

func cekLink(url string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "sedang down")
		return
	}
	fmt.Println(url, "jalan normal")
}
