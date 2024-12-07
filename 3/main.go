package main

import (
	"adventOfCode2024/customScanner"
	"fmt"
	"log"
	"regexp"
	"strconv"
)

func main() {
	scanner := customScanner.GetFileScanner("data.data")
	partOneResponse := partOne(scanner)
	fmt.Println(partOneResponse)
	scanner = customScanner.GetFileScanner("data.data")
	partTwoResponse := partTwo(scanner)
	fmt.Println(partTwoResponse)
}

func partOne(scanner *customScanner.CustomScanner) int {
	defer scanner.Close()
	answer := 0
	for scanner.Scan() {
		line := scanner.Text()
		value, _ := getMulValue(line, true)
		answer += value
	}
	return answer
}

func partTwo(scanner *customScanner.CustomScanner) int {
	defer scanner.Close()
	answer := 0
	do := true
	var value int
	for scanner.Scan() {
		line := scanner.Text()
		value, do = getMulValue(line, do)
		answer += value
	}
	return answer
}

func getMulValue(line string, do bool) (int, bool) {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)|(do\(\)|don't\(\))`)
	matches := re.FindAllStringSubmatch(line, -1)
	if matches == nil {
		return 0, do
	}
	answer := 0
	for i := 0; i < len(matches); i++ {
		if matches[i][3] != "" {
			do = matches[i][3] == "do()"
			continue
		}
		if !do {
			continue
		}
		match1, err1 := strconv.Atoi(matches[i][1])
		match2, err2 := strconv.Atoi(matches[i][2])
		if err1 != nil || err2 != nil {
			log.Fatal("Error parsing strings to numbers")
		}
		answer += match1 * match2
	}
	return answer, do
}
