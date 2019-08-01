package main

import (
	"fmt"
)

func main() {
	variableBoolean := true

	// Kondisi If
	if variableBoolean == true {
		fmt.Println("True")
	}

	// Kondsi If dengan Else
	if variableBoolean == false {
		fmt.Println("Bukan False")
	} else {
		fmt.Println("True")
	}

	variableAngka := 2
	// Kondisi If berulang / nested if
	if variableAngka == 1 {
		fmt.Println("Angka = 1")
	} else if variableAngka == 2 {
		fmt.Println("Angka Bukan 1, Angka = 2")
	} else {
		fmt.Println("Angka Bukan 1 & 2")
	}

	// Kondisi menggunakan switch
	switch angka := variableAngka; angka {
	case 1:
		fmt.Println("Angka = 1")
	case 2:
		fmt.Println("Angka Bukan 1, Angka = 2")
	default:
		fmt.Println("Angka Bukan 1 & 2")
	}
}
