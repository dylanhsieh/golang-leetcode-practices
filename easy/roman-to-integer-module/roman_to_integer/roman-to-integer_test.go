package roman_to_integer

import (
	"testing"
)

func Test_III(t *testing.T) {
	if RomanToInt("III") != 3 {
		t.Error("wrong result")
	}
}

func Test_LVIII(t *testing.T) {
	if RomanToInt("LVIII") != 58 {
		t.Error("wrong result")
	}
}

func Test_MCMXCIV(t *testing.T) {
	if RomanToInt("MCMXCIV") != 1994 {
		t.Error("wrong result")
	}
}
