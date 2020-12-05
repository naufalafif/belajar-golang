package arraySlices

func Sum(angka2 [5]int) int {
	hasil := 0
	for _, angka := range angka2 {
		hasil = hasil + angka
	}
	return hasil
}
