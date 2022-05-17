package main

import "fmt"

func main() {
	fmt.Println(intToRoman(3))
}

func intToRoman(num int) string {
	fmt.Println(fmt.Sprintf("initial value of num is %v", num))
	romanNum := ""
	numConvertedToRoman := false
	workedNum := num
	// let's iterate over num and subtractively distill it's roman numerals from it
	for !numConvertedToRoman {
		// it seems like switch isn't the way to go here
		if workedNum < 3999 && workedNum > 999 {
			romanNum += "M"
			workedNum -= 1000
		} else if workedNum < 1000 && workedNum > 899 {
			romanNum += "CM"
			workedNum -= 900
		} else if workedNum < 900 && workedNum > 499 {
			romanNum += "D"
			workedNum -= 500
		} else if workedNum < 500 && workedNum > 399 {
			romanNum += "CD"
			workedNum -= 400
		} else if workedNum < 400 && workedNum > 99 {
			romanNum += "C"
			workedNum -= 100
		} else if workedNum < 100 && workedNum > 89 {
			romanNum += "XC"
			workedNum -= 90
		} else if workedNum < 90 && workedNum > 49 {
			romanNum += "L"
			workedNum -= 50
		} else if workedNum < 50 && workedNum > 39 {
			romanNum += "XL"
			workedNum -= 40
		} else if workedNum < 40 && workedNum > 9 {
			romanNum += "X"
			workedNum -= 10
		} else if workedNum < 10 && workedNum > 5 {
			romanNum += "IX"
			workedNum -= 9
		} else if workedNum < 9 && workedNum > 4 {
			romanNum += "V"
			workedNum -= 5
		} else if workedNum < 5 && workedNum > 3 {
			romanNum += "IV"
			workedNum -= 4
		} else if workedNum < 4 && workedNum > 0 {
			romanNum += "I"
			workedNum -= 1
		} else if workedNum == 0 {
			numConvertedToRoman = true
		} else {
			fmt.Println(fmt.Sprintf("undetermined error for workedNum of %v, romanNum of %v, and num of %v", workedNum, romanNum, num))
		}
	}
	return romanNum
}
