package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type CustomScanner struct {
	io.Closer
	*bufio.Scanner
}

func main() {
	scanner := getFileScanner("data.data")
	partOneResponse := partOne(scanner)
	scanner = getFileScanner("data.data")
	partTwoResponse := partTwo(scanner)
	fmt.Println(partOneResponse)
	fmt.Println(partTwoResponse)
}

func partOne(scanner *CustomScanner) int {
	defer scanner.Close()
	validLines := 0
	for scanner.Scan() {
		line := scanner.Text()
		parsed := parseLine(line)
		if validateData(parsed) {
			validLines += 1
		}
	}
	return validLines
}

func partTwo(scanner *CustomScanner) int {
	defer scanner.Close()
	validLines := 0
	for scanner.Scan() {
		line := scanner.Text()
		parsed := parseLine(line)
		if validateData(parsed) {
			validLines += 1
		} else {
			for j := 0; j < len(parsed); j++ {
				reshapedParsed := deleteFromSlice(parsed, j)
				if validateData(reshapedParsed) {
					validLines += 1
					break
				}
			}
		}
	}
	return validLines

}

func getFileScanner(fileName string) *CustomScanner {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	return &CustomScanner{file, scanner}
}

func parseLine(line string) []int {
	splitted := strings.Split(line, " ")
	result := []int{}
	for _, element := range splitted {
		intElement, _ := strconv.Atoi(element)
		result = append(result, intElement)
	}
	return result
}

func validateData(line []int) bool {
	if len(line) == 1 {
		return true
	}
	isIncreasing := line[0] < line[1]
	for i := 0; i < len(line)-1; i++ {
		if line[i] > line[i+1] && isIncreasing {
			return false
		} else if line[i] < line[i+1] && !isIncreasing {
			return false
		}
		difference := abs(line[i] - line[i+1])
		if difference > 3 || difference == 0 {
			return false
		}
	}
	return true
}

func deleteFromSlice(slice []int, index int) []int {
	temporary := make([]int, len(slice))
	copy(temporary, slice)
	return append(temporary[:index], temporary[index+1:]...)
}

func abs(value int) int {
	if value < 0 {
		return value * -1
	}
	return value
}