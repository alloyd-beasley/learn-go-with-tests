package iteration

func Repeat(val string) (repeated string) {

	for i := 0; i < 5; i++ {
		repeated += val
	}

	return
}
