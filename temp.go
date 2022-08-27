package temperature

import "fmt"

func Conv(temp float64, symbol string) (float64, error) {
	if symbol == "C" || symbol == "c" {
		return CtoF(temp)
	} else if symbol == "F" || symbol == "f" {
		return FtoC(temp)
	}
	return 0, fmt.Errorf("unrecognized symbol: %q, want: F, f, C or c\n", symbol)
}

func FtoC(a float64) (float64, error) {
	return (a - 32) * 5 / 9, nil
}
func CtoF(b float64) (float64, error) {
	return (b * 1.8) + 32, nil
}
