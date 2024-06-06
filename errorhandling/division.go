package errorhandling

import "errors"

func Divide(a, b float64) (float64, error) {

	if b == 0 {
		return 0, errors.New("Devision by zero is not allowed!")
	}

	return a / b, nil
}
