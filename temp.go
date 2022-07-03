package temperature

import "fmt"

func Conv(temp float64, symbol string) (float64, error) {
	if symbol == "C" || symbol == "c" {
		return CtoF(temp), nil
	} else if symbol == "F" || symbol == "f" {
		return FtoC(temp), nil
	}
	return 0, fmt.Errorf("unrecognized symbol: %q, want: F, f, C or c\n", symbol)
}

func FtoC(a float64) float64 {
	return (a - 32) * 5 / 9
}
func CtoF(b float64) float64 {
	return (b * 1.8) + 32
}
