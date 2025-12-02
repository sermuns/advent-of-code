package main

import (
	"fmt"
	"strconv"
	"unicode"
	"os"
)

func getFactor(inputText string, startIndex int) (factor int, factorLen int, err error) {
	factorStr := ""

	var i int
	for i = 0; i < 3; i++ {
		digit := inputText[startIndex+i]
		if unicode.IsDigit(rune(digit)) {
			factorStr += string(digit)
		} else {
			break
		}
	}

	factor, err = strconv.Atoi(factorStr)
	if err != nil {
		return -1, 0, err
	}

	return factor, i, nil
}

func getSumFromInput(inputText string) int {
	sum := 0
	i := 0
	for i < len(inputText) {
		if inputText[i] == 'm' &&
			inputText[i+1] == 'u' &&
			inputText[i+2] == 'l' &&
			inputText[i+3] == '(' {

			i += 4

			var f1 int
			f1, factorLen, err := getFactor(inputText, i)
			i += factorLen
			if err != nil {
				fmt.Println("unable to get f1")
				continue
			}

			if inputText[i] != ',' {
				fmt.Println("no , found")
				continue
			}
			i++

			var f2 int
			f2, factorLen, err = getFactor(inputText, i)
			i += factorLen
			if err != nil {
				fmt.Println("unable to get f2")
				continue
			}

			if inputText[i] != ')' {
				fmt.Println("no end ) found")
				continue
			}

			sum += f1 * f2
		} else {
			fmt.Println(i, "no mul( found")
		}

		i++
	}

	// scanna genom sträng, hitta "mul("
	// hitta f1: "1-3st siffror"
	// hitta ","
	// hitta f2: "1-3st siffror"
	// hitta ")"

	// OM ALLT: sum += f1*f2
	return sum
}

func main() {
	fileBytes, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	input := string(fileBytes)
	// input := "mul(2,5)mul(2,10)"
	fmt.Println(getSumFromInput(input))
}
