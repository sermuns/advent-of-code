package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func getFactor(letter rune) {
	if unicode.IsDigit(letter) {
		factorStr += string(letter)
	} else {
		factorInt, err := strconv.Atoi(factorStr)
		if err != nil {
			panic(err)
		}
	}
}

func getSumFromInput(inputText string) int {
	// sum := 0
	mulTemplate := "mul("
	mulIndex := 0

	factorStr := ""

	for _, letter := range inputText {
		// handle digits
		if mulIndex >= len(mulTemplate) {

			continue
		}

		if letter == rune(mulTemplate[mulIndex]) {
			mulIndex++
		} else {
			mulIndex = 0
		}
	}

	// scanna genom sträng, hitta "mul("
	// hitta f1: "1-3st siffror"
	// hitta ","
	// hitta f2: "1-3st siffror"
	// hitta ")"

	// OM ALLT: sum += f1*f2
	return 0
}

func main() {
	// fileBytes, err := os.ReadFile("input")
	// if err != nil {
	// 	panic(err)
	// }
	//
	// inputText := string(fileBytes)
	getSumFromInput("mul8(21,100)mul(21,100)")
}
