package iteration

//Repeat a string 5 times
func Repeat(character string) (repeated string) {
	for i := 0; i < 5; i++ {
		repeated += character
	}
	return
}
