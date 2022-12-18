// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"roman-to-integer-module/roman_to_integer"
)

func romanToInt(romanString string) int {
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	sum := 0
	biggestNumber := 0
	// check if the are no roman symbol in string ?
	// string convert to char array
	for i := len(romanString) - 1; i >= 0; i-- {
		// fmt.Printf("%d,", romanMap[romanString[i]])
		// if current number >= biggest number then sum it up,
		// and update biggest number
		// otherwise, subtract it
		currentNumber := romanMap[romanString[i]]

		if (currentNumber >= biggestNumber) || (biggestNumber == 0) {
			sum += currentNumber
			biggestNumber = currentNumber
		} else {
			sum -= currentNumber
		}

	}

	// make an map to store roman numbers int mapping
	// make sure the are in ordered
	// fmt.Println(romanMap)
	return sum
}

func main() {
	fmt.Println(roman_to_integer.RomanToInt("III"))
	fmt.Println(roman_to_integer.RomanToInt("LVIII"))
	fmt.Println(roman_to_integer.RomanToInt("MCMXCIV"))

}
