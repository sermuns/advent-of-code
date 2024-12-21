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

func getDiffs(report []int) []int {
	diffs := []int{}
	for i := 1; i < len(report); i++ {
		diffs = append(diffs, report[i]-report[i-1])
	}
	return diffs
}

func isSafe(report []int) bool {
	diffs := getDiffs(report)
	numFails := 0

	positives := 0
	negatives := 0
	for _, diff := range diffs {
		// safe diff?
		absDiff := diff
		if absDiff < 0 {
			absDiff *= -1
		}
		if absDiff < 1 || absDiff > 3 {
			numFails++
		}

		// monotonous?
		if diff > 0 {
			positives++
		} else if diff < 0 {
			negatives++
		}
	}
	if positives > 0 && negatives > 0{
		numFails++
	}

	return numFails < 2
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
