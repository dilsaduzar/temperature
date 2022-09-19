package temperature

import "fmt"

func Conv(temp float64, symbol string) (float64, string, error) {
	if symbol == "C" || symbol == "c" {
		return CtoF(temp)
	} else if symbol == "F" || symbol == "f" {
		return FtoC(temp)
	}
	return 0, "", fmt.Errorf("unrecognized symbol: %q\n, want: F, f, C or c", symbol)
}

func FtoC(a float64) (float64, string, error) {
	return (a - 32) * 5 / 9, "C", nil
}
func CtoF(b float64) (float64, string, error) {
	return (b * 1.8) + 32, "F", nil
}
