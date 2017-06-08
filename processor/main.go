package processor

// Process a string of digits, and return an integer array if found a valid Lotto pick, otherwise return empty integer array
// Essentially, we are doing a depth first search, and each node has max 2 children
//  1. Input string length needs to be in the range of [7, 14]
//  2. Given the length, we can pre-determine the number of single digits and the number of double digits, to improve performance
// NOTE: this is separated out so that it can be unit tested

var lottoPickLength = 7
var maxInputLength = 14

// Process the digits string, and return lotto picks
func Process(input string) []int {
	inputLength := len(input)

	// filter out input that's too long or too short
	if inputLength > maxInputLength || inputLength < lottoPickLength {
		return nil
	}

	// we keep track of the total number of double digits to avoid unnecessary search
	numOfDoubleDigits := (inputLength - lottoPickLength)

	// convert string into array of int
	digits := make([]int, inputLength)
	for index, element := range input {
		// char - '0'
		digits[index] = int(element - '0')
	}

	output, _ := recursiveSearch(digits, make(map[int]bool), 0, numOfDoubleDigits)

	// reverse the output
	for i, j := 0, len(output)-1; i < j; i, j = i+1, j-1 {
		output[i], output[j] = output[j], output[i]
	}

	return output
}

// lottoPickSet: keeps track the current lotto picks, also to ensure uniqueness with good performance. NOTE: Golang doesn't have Set
// curIndex: points to the current digit
// numOfDoubleDigitsLeft: keeps track of the number of double digits left
func recursiveSearch(digits []int, lottoPickSet map[int]bool, curIndex int, numOfDoubleDigitsLeft int) ([]int, bool) {
	// return true if reached the end, and has 7 unique items
	if curIndex > len(digits)-1 && len(lottoPickSet) == lottoPickLength {
		return []int{}, true
	}

	// return if reached the end, or we didn't use all the digits
	if curIndex > len(digits)-1 || len(lottoPickSet) > lottoPickLength {
		return nil, false
	}

	// 2 cases here: pick a single digit, or a double digit

	// 1. try the pick a single digit number
	curDigit := digits[curIndex]
	// check if the curDigit has NOT already been picked, and within range
	if _, exists := lottoPickSet[curDigit]; curDigit > 0 && curDigit < 60 && !exists {
		// add to the set, and continue
		lottoPickSet[curDigit] = true
		if lottoPicks, found := recursiveSearch(digits, lottoPickSet, curIndex+1, numOfDoubleDigitsLeft); found {
			// add the curDigit to the result array
			lottoPicks = append(lottoPicks, curDigit)
			return lottoPicks, true
		}

		// otherwise, if not found, remove the curDigit from the set
		delete(lottoPickSet, curDigit)
	}

	// 2. try pick a double digit number
	if numOfDoubleDigitsLeft > 0 && curIndex+1 < len(digits) {
		curDigit := digits[curIndex]*10 + digits[curIndex+1]
		// check if the curDigit has NOT already been picked, and within range
		if _, exists := lottoPickSet[curDigit]; curDigit > 0 && curDigit < 60 && !exists {
			// add to the set, and continue
			lottoPickSet[curDigit] = true
			if lottoPicks, found := recursiveSearch(digits, lottoPickSet, curIndex+2, numOfDoubleDigitsLeft-1); found {
				// add the curDigit to the result array
				lottoPicks = append(lottoPicks, curDigit)
				return lottoPicks, true
			}

			// otherwise, if not found, remove the curDigit from the set
			delete(lottoPickSet, curDigit)
		}
	}

	return nil, false
}
