package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readReports(fileName string) [][]int {
	file, err := os.Open(fileName)
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	reports := [][]int{}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break // assume EOF
		}

		strReport := strings.Split(line, " ")
		report := make([]int, 0, len(strReport))
		for _, strLevel := range strReport {
			level, err := strconv.Atoi(
				strings.TrimSpace(strLevel),
			)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			report = append(report, level)
		}
		reports = append(reports, report)
	}
	return reports
}

func getMonotinyFails(report []int) int {
	fails := 0
	increasing := report[1] > report[0]
	for i := 2; i < len(report); i++ {
		if (report[i] > report[i-1]) != increasing {
			fails++
		}
	}
	return fails
}

func getDiffFails(report []int) int {
	fails := 0
	for i := 1; i < len(report); i++ {
		diff :=	report[i] - report[i-1]
		if diff < 0{
			diff *= -1
		}
		if diff < 1 || diff > 3 {
			fails++
		}
	}
	return fails
}

func isSafe(report []int) bool {
	numFails := 0	
	
	numFails += getMonotinyFails(report)

	numFails += getDiffFails(report)

	return numFails < 1
}

func main() {
	reports := readReports("input")

	numSafeReports := 0
	for _, report := range reports {
		if isSafe(report) {
			numSafeReports++
		}
	}
	fmt.Println(numSafeReports)
}
