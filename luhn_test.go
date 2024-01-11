package dataformats

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type (
	validationTest struct {
		number  string
		isValid bool
		err     error
	}
	calculateLuhnAndAdd struct {
		number      string
		finalNumber string
	}
)

func Test_ValidateLuhn(t *testing.T) {
	testCases := []validationTest{

		{"1234567812345670", true, nil},
		{"1111222233334444", true, nil},
		{"49927398716", true, nil},
		{"79927398713", true, nil},
		{"374652346956782346957823694857692364857368475368", true, nil},
		{"0", true, nil},
		{"1111222233334441", false, nil},
		{"49927398717", false, nil},
		{"1234567812345678", false, nil},
		{"79927398710", false, nil},
		{"79927398711", false, nil},
		{"79927398712", false, nil},
		{"79927398714", false, nil},
		{"79927398715", false, nil},
		{"79927398716", false, nil},
		{"79927398717", false, nil},
		{"79927398718", false, nil},
		{"79927398719", false, nil},
		{"374652346956782346957823694857692364857387456834", false, nil},
		{"8", false, nil},
	}

	for _, tc := range testCases {
		t.Run("luhn validation test cases for "+tc.number, func(t *testing.T) {
			res, err := ValidateLuhn(([]byte)(tc.number))
			require.Nil(t, err)
			require.Equal(t, tc.isValid, res)
		})

	}
	t.Run("luhn validation validationTest cases for invalid digit", func(t *testing.T) {
		res, err := ValidateLuhn(([]byte)("abcdef"))
		require.Equal(t, false, res)
		require.NotNil(t, err)
		require.Error(t, err, "invalid input: not a number")
	})
}

func Test_CalculateAndAddLuhnDigit(t *testing.T) {
	testCases := []calculateLuhnAndAdd{
		{"12345", "123455"},
		{"111122223333444", "1111222233334444"},
		{"7992739871", "79927398713"},
		{"37465234695678234695782369485769236485736847536", "374652346956782346957823694857692364857368475368"},
	}

	for _, tc := range testCases {
		t.Run("calculateLuhnAndAdd test cases for "+tc.number, func(t *testing.T) {
			luhnData, err := CalculateAndAddLuhnDigit(([]byte)(tc.number))
			require.Nil(t, err)
			require.Equal(t, tc.finalNumber, string(luhnData))
		})

	}
}
