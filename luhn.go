package dataformats

import (
	"errors"
	"strconv"
)

func ValidateLuhn(data []byte) (bool, error) {
	sum, err := calculateSum(string(data), len(data)%2)
	if err != nil {
		return false, err
	}
	return sum%10 == 0, nil
}

// function calculate the luhn and append it to data as well.
func CalculateAndAddLuhnDigit(data []byte) ([]byte, error) {
	l := len(data)
	if l == 0 {
		return nil, errors.New("invalid input: data too short")
	}

	sum, err := calculateSum(string(data), (len(data)+1)%2)
	if err != nil {
		return nil, err
	}

	luhnDigit := (sum * 9) % 10

	return ([]byte)(string(data) + strconv.Itoa(luhnDigit)), nil
}

func calculateSum(number string, parity int) (int, error) {
	var sum int
	for i, d := range number {
		mod, err := strconv.Atoi(string(d))
		if err != nil {
			return 0, errors.New("invalid input: not a number")
		}
		if i%2 == parity {
			mod *= 2
			if mod > 9 {
				mod = (mod % 10) + 1
			}
		}

		sum += mod
	}
	return sum, nil
}
