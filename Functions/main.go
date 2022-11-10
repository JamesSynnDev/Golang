package Functions

func divAndRemainer(numerator int, denominator int) (result int, remainder int, err error) {
	if denominator == 0 {
		return 0, 0, error.New("Cannot divide by zero")

	}
	return numerator / denominator, numerator % denominator, nil

}
