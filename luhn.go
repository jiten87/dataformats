package dataformats

import (
	"errors"
	"strconv"
)

func ValidateLuhn(data []byte) (bool, error) {
	sum, err := calculateSum(string(data), len(data))
	if err != nil {
		return false, err
	}
	return sum%10 == 0, nil
}

// This function calculate the luhn and append it to data as well.
func CalculateAndAddLuhnDigit(data []byte) ([]byte, error) {
	l := len(data)
	if l == 0 {
		return nil, errors.New("invalid input: data too short")
	}

	sum, err := calculateSum(string(data), len(data))
	if err != nil {
		return nil, err
	}

	luhnDigit := 10 - sum%10

	return append(data, byte(luhnDigit)), nil
}

func calculateSum(number string, len int) (int, error) {
	var alternate bool
	var sum int
	for i := len - 1; i > -1; i-- {
		mod, err := strconv.Atoi(string(number[i]))
		if err != nil {
			return 0, errors.New("invalid input: not a number")
		}
		if alternate {
			mod *= 2
			if mod > 9 {
				mod = (mod % 10) + 1
			}
		}
		alternate = !alternate
		sum += mod
	}
	return sum, nil
}
