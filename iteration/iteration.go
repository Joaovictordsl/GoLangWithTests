package iteration

func Repeat(repeatCount int, character string) string {
	var repeated string

	for i := 0; i < repeatCount; i++ {
		repeated += character
	}

	return repeated

}
