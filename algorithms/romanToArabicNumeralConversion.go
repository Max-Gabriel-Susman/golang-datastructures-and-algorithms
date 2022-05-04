package conversion

import "fmt"

func romanToInt(s string) int {
	var err error
	runningTotal := 0
	current_total := fmt.Sprintf("Runningtotal is now: %d", runningTotal)
	fmt.Println(current_total)
	fmt.Println("---------------------------------------------")

	romanToArabic := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	fmt.Println(romanToArabic)

	for i := 0; i < len(s); i++ {
		// fmt.Println(string(s[i]))
		// checks for character with offset greater than 0 because nested conditionals are checking against preceding characters
		if i > 0 {
			// the following three ifs check for a preceding character to the current character that will dictate subtraction
			if string(s[i]) == "V" || string(s[i]) == "X" {
				if string(s[i]) == "I" {

				}
			} else if string(s[i]) == "L" || string(s[i]) == "C" {
				if string(s[i]) == "X" {

				}
			} else if string(s[i]) == "D" || string(s[i]) == "M" {
				if string(s[i]) == "C" {

				}
			} else {

			}
		}
		runningTotal, err = outRomanToArabicState(string(s[i]), i, runningTotal)
	}
	fmt.Println(err)
	return runningTotal
}

func outRomanToArabicState(romaDigit string, digitPlace int, runningTotal int) (newRunningTotal int, err error) {
	fmt.Println()
	current_digit := fmt.Sprintf("The %vth roman digit is : %v", digitPlace, romaDigit)
	fmt.Println(current_digit)
	runningTotal += romanToArabic[romaDigit]
	current_total = fmt.Sprintf("Runningtotal is now: %d", runningTotal)
	fmt.Println(current_total)
	fmt.Println("---------------------------------------------")
	return runningTotal, nil
}
