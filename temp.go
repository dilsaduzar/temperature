package temperature

import "errors"

var ErrTemp = errors.New("unrecognized symbol, want: F, f, C or c")

func Conv(temp float64, symbol string) (float64, string, error) {

	if symbol == "C" || symbol == "c" {
		return CtoF(temp)
	} else if symbol == "F" || symbol == "f" {
		return FtoC(temp)
	}
	return 0, "", ErrTemp
}

func FtoC(a float64) (float64, string, error) {
	return (a - 32) * 5 / 9, "C", nil
}
func CtoF(b float64) (float64, string, error) {
	return (b * 1.8) + 32, "F", nil
}
