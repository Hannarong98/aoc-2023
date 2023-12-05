package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return solve2()
	}
	// solve part 1 here
	return solve1()
}

func getNumberPairString(s1 string, s2 string) string {
	return fmt.Sprintf("%s%s", s1, s2)
}

func omitAlphabets(s string) string {

	m1 := regexp.MustCompile("[a-zA-z]")

	return m1.ReplaceAllLiteralString(s, "")
}

func solve1() int {
	file, err := os.Open("input-user.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	result := 0
	for scanner.Scan() {

		cleaned := omitAlphabets(string(scanner.Text()))
		lineLen := len(cleaned)

		if lineLen > 1 {

			first := string(cleaned[0])
			last := string(cleaned[lineLen-1])

			num, err := strconv.Atoi(getNumberPairString(first, last))
			if err != nil {
				log.Fatal(err)
			}

			result += num
		} else if lineLen == 1 {

			first := string(cleaned[0])
			last := string(cleaned[0])

			num, err := strconv.Atoi(getNumberPairString(first, last))
			if err != nil {
				log.Fatal(err)
			}

			result += num
		}
	}

	return result
}

func solve2() int {
	file, err := os.Open("input-user.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	numWordSlice := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	result := 0
	curr := 1
	for scanner.Scan() {

		line := scanner.Text()

		var first string
		var last string

		firstCharIsDigit := isDigit(0, line)

		if !firstCharIsDigit {
			firstNumWord := ""
			firstNumWordIndex := -1

			firstNumDigit := ""
			firstNumDigitIndex := -1

			m1 := regexp.MustCompile(`\d`)

		out:
			for _, char := range line {
				if m1.MatchString(string(char)) {
					firstNumDigit = string(char)
					firstNumDigitIndex = strings.Index(line, string(char))
					break out
				}
			}

			for _, word := range numWordSlice {
				if strings.Contains(line, word) {

					indexOfCurrentNumber := strings.Index(line, word)

					if firstNumWordIndex == -1 && firstNumWord == "" {
						firstNumWord = word
						firstNumWordIndex = indexOfCurrentNumber
					}

					if indexOfCurrentNumber < firstNumWordIndex {
						firstNumWord = word
						firstNumWordIndex = indexOfCurrentNumber
					}

				}
			}

			if firstNumDigitIndex > firstNumWordIndex && firstNumWord == "" && firstNumWordIndex == -1 {
				first = firstNumDigit
			} else if firstNumDigitIndex < firstNumWordIndex && firstNumDigitIndex != -1 {
				first = firstNumDigit
			} else {
				first = firstNumWord
			}

		} else {
			first = string(line[0])
		}

		lastCharIsDigit := isDigit(1, line)

		if !lastCharIsDigit {
			lastNumberWord := ""
			lastNumberWordIndex := -1

			lastNumberDigit := ""
			lastNumberDigitIndex := -1

			m1 := regexp.MustCompile(`\d`)

			for _, char := range line {
				if m1.MatchString(string(char)) {
					lastNumberDigit = string(char)
					lastNumberDigitIndex = strings.LastIndex(line, string(char))
				}
			}

			for _, word := range numWordSlice {
				if strings.Contains(line, word) {

					lastIndexOfCurrentNumber := strings.LastIndex(line, word)

					if lastIndexOfCurrentNumber > lastNumberWordIndex && lastIndexOfCurrentNumber != -1 {
						lastNumberWord = word
						lastNumberWordIndex = lastIndexOfCurrentNumber
					}

				}
			}

			if lastNumberDigitIndex > lastNumberWordIndex {
				last = lastNumberDigit
			} else {
				last = lastNumberWord
			}

		} else {
			last = string(line[len(line)-1])
		}

		pair := fmt.Sprintf("%s%s", convertToNumericString(string(first)), convertToNumericString(string(last)))

		num, err := strconv.Atoi(pair)
		if err != nil {
			log.Fatal(err)
		}

		result += num
		curr++
	}

	return result
}

func isDigit(mode int, line string) bool {

	if mode == 0 {
		matched, err := regexp.MatchString("\\d", string(line[0]))
		if err != nil {
			log.Fatal(err)
		}

		return matched
	} else {
		matched, err := regexp.MatchString("\\d", string(line[len(line)-1]))
		if err != nil {
			log.Fatal(err)
		}

		return matched
	}

}

func convertToNumericString(s string) string {

	matched, err := regexp.MatchString("\\d", s)
	if err != nil {
		log.Fatal(err)
	}

	if matched {
		return s
	} else {
		sMap := map[string]string{
			"one":   "1",
			"two":   "2",
			"three": "3",
			"four":  "4",
			"five":  "5",
			"six":   "6",
			"seven": "7",
			"eight": "8",
			"nine":  "9",
		}

		return sMap[s]
	}

}
