package roman_to_integer

func RomanToInt(romanString string) int {
	// Map to store roman numbers to int mapping
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

	// Iterate over the romanString backwards
	for i := len(romanString) - 1; i >= 0; i-- {
		// Get the current number
		currentNumber := romanMap[romanString[i]]

		// If the current number is greater than or equal to the biggest number, add it to the sum
		// Otherwise, subtract it
		if (currentNumber >= biggestNumber) || (biggestNumber == 0) {
			sum += currentNumber
			biggestNumber = currentNumber
		} else {
			sum -= currentNumber
		}
	}

	return sum
}
