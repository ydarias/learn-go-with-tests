package iteration

func Repeat(character string, times int) string {
	// TODO check the times is larger than zero

	var repeated string

	for i := 0; i < times; i++ {
		repeated += character
	}
	return repeated
}
