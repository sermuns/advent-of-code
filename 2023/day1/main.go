package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func wordToDigit(word string) int{
	return map[string] int{
			"one": 1,
			"two": 2,
			"three": 3,
			"four": 4,
			"five": 5,
			"six": 6,
			"seven": 7,
			"eight": 8,
			"nine": 9,
	}[word]
}

func readCalibrationValues(fileName string) []int {
	file, err := os.Open(fileName)
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	values := []int{}
	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break // assume EOF
		}

		digits := []rune{0, 0}
		for _, letter := range line {
			if !unicode.IsDigit(letter){
				continue
			}
			if digits[0] == 0 {
				digits[0] = letter
			}
			digits[1] = letter
		}
		val, err := strconv.Atoi(string(digits))
		values = append(values, val)
	}
	return values
}

func sum(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}

func main() {
	calVals := readCalibrationValues("input")

	
}
