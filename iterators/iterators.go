package iterators

func Repeat(character string, iterations uint8) (repeated string) {
	if (iterations == 0) {
		iterations = 5
	}
	for i := 0; i < int(iterations); i++ {
		repeated += character
	}
	return repeated
}