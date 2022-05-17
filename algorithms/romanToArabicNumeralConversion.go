package main

import "fmt"

/*
func main() {
	fmt.Println(romanToInt("III"))
}
*/

var romanToArabic = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	var err error
	runningTotal := 0

	for i := 0; i < len(s); i++ {
		// fmt.Println(string(s[i]))
		// checks for character with offset greater than 0 because nested conditionals are checking against preceding characters
		if i > 0 {
			// the following three ifs check for a preceding character to the current character that will dictate subtraction
			if string(s[i]) == "V" || string(s[i]) == "X" {
				if string(s[i]) == "I" {
					runningTotal, err = outRomanToArabicState(string(s[i]), runningTotal, "I")
				}
			} else if string(s[i]) == "L" || string(s[i]) == "C" {
				if string(s[i]) == "X" {
					runningTotal, err = outRomanToArabicState(string(s[i]), runningTotal, "X")
				}
			} else if string(s[i]) == "D" || string(s[i]) == "M" {
				if string(s[i]) == "C" {
					runningTotal, err = outRomanToArabicState(string(s[i]), runningTotal, "C")
				}
			} else {
				runningTotal, err = outRomanToArabicState(string(s[i]), runningTotal, "")
			}
		} else {
			runningTotal, err = outRomanToArabicState(string(s[i]), runningTotal, "")
		}
	}
	fmt.Println(err)
	return runningTotal
}

func outRomanToArabicState(romaDigit string, runningTotal int, subtractivePrecedent string) (newRunningTotal int, err error) {
	if subtractivePrecedent != "" {
		runningTotal += romanToArabic[romaDigit] - romanToArabic[subtractivePrecedent]
		return runningTotal, nil
	}
	runningTotal += romanToArabic[romaDigit]
	return runningTotal, nil
}
