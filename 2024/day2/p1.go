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

func isSafeDiff(diff int) bool {
	absDiff := diff
	if absDiff < 0 {
		absDiff *= -1
	}
	if absDiff < 1 || absDiff > 3 {
		return false
	}
	return true
}

func isMonotonous(diff int, isIncreasing bool) bool {
	if isIncreasing && diff < 0 {
		return false
	} else if !isIncreasing && diff > 0 {
		return false
	}
	return true
}

func trySetFailIndex(i int, failIndex *int) bool {
	if *failIndex != -1 {
		return false
	}

	*failIndex = i
	return true
}

func isSafe(report []int) (int, bool) {
	diffs := getDiffs(report)

	increases := 0
	decreases := 0

	// Figure out if increasaing/decreasing
	for _, diff := range diffs {
		if diff > 0 {
			increases++
		} else if diff < 0 {
			decreases++
		}
	}
	isIncreasing := increases > decreases

	failIndex := -1
	for i, diff := range diffs {
		newFailIndex := -1
		if !isSafeDiff(diff) {
			newFailIndex = i
		}
		if !isMonotonous(diff, isIncreasing) {
			newFailIndex = i
		}
		if newFailIndex != -1 {
			if !trySetFailIndex(i, &failIndex) {
				return failIndex + 1, false
			}
		}
	}
	return -1, true
}

func main() {
	reports := readReports("input")

	numSafeReports := 0
	for _, report := range reports {
		failIndex, wasSafe := isSafe(report)
		if wasSafe {
			numSafeReports++
			continue
		}

		rescuedReport := append(report[:failIndex], report[failIndex+1:]...)

		_, wasSafe = isSafe(rescuedReport)
		if wasSafe {
			numSafeReports++
		}
	}
	fmt.Println(numSafeReports)
}
