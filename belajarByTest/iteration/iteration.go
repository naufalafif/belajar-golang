package iteration

// Repeat - mengulang & menggabungkan string sebanyak 5x
func Repeat(alphabet string) string {
	var fiveAlphabet string
	for i := 0; i < 5; i++ {
		fiveAlphabet = fiveAlphabet + alphabet
	}
	return fiveAlphabet
}
