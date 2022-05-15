package main

import "fmt"

const four = "IV"
const fiftyEight = "LVIII"

func main() {

	fmt.Println(fmt.Sprintf("%v in roman numerals is %v", fmt.Sprint(romanToInt(four)), four))
	fmt.Println(fmt.Sprintf("%v in roman numerals is %v", fmt.Sprint(romanToInt(fiftyEight)), fiftyEight))
	/*
		fmt.Println(fmt.Sprintf("%v in roman numerals is %v", fmt.Sprint(romanToInt(four)), four))
		fmt.Println(fmt.Sprintf("%v in roman numerals is %v", fmt.Sprint(romanToInt(four)), four))
		fmt.Println(fmt.Sprintf("%v in roman numerals is %v", fmt.Sprint(romanToInt(four)), four))
		fmt.Println(fmt.Sprintf("%v in roman numerals is %v", fmt.Sprint(romanToInt(four)), four))
		fmt.Println(fmt.Sprintf("%v in roman numerals is %v", fmt.Sprint(romanToInt(four)), four))
	*/
}

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
		currentNumeral := string(s[i])
		// fmt.Println(string(s[i]))
		// checks for character with offset greater than 0 because nested conditionals are checking against preceding characters
		if i > 0 {
			previousNumeral := string(s[i-1])
			// the following three ifs check for a preceding character to the current character that will dictate subtraction
			if currentNumeral == "V" || currentNumeral == "X" {
				if previousNumeral == "I" {
					runningTotal, err = outRomanToArabicState(currentNumeral, runningTotal, previousNumeral)
				}
			} else if currentNumeral == "L" || currentNumeral == "C" {
				if previousNumeral == "X" {
					runningTotal, err = outRomanToArabicState(currentNumeral, runningTotal, previousNumeral)
				}
			} else if currentNumeral == "D" || currentNumeral == "M" {
				if previousNumeral == "C" {
					runningTotal, err = outRomanToArabicState(currentNumeral, runningTotal, previousNumeral)
				}
			} else {
				runningTotal, err = outRomanToArabicState(currentNumeral, runningTotal, "")
			}
		} else {
			runningTotal, err = outRomanToArabicState(currentNumeral, runningTotal, "")
		}
	}
	fmt.Println(err)
	return runningTotal
}

func outRomanToArabicState(romaDigit string, runningTotal int, subtractivePrecedent string) (newRunningTotal int, err error) {
	romaValue := romanToArabic[romaDigit]
	if subtractivePrecedent != "" {
		romaValue -= romanToArabic[subtractivePrecedent] * 2
		runningTotal += romaValue
		fmt.Println(fmt.Sprintf("%v in roman numerals is %v", fmt.Sprint(romanToInt(four)), four))
		return runningTotal, nil
	}
	fmt.Println(fmt.Sprintf("%v in roman numerals is %v", fmt.Sprint(romanToInt(four)), four))
	runningTotal += romaValue
	return runningTotal, nil
}
