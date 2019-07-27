package main

import "fmt"

func main() {

	contohMap := make(map[string]string)
	contohMap["nama"] = "naufal"
	contohMap["alamat"] = "jogja"

	delete(contohMap, "alamat")

	fmt.Println(contohMap)
}
