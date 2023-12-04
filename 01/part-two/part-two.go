package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var wordToNumber = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func extractNumbersFromRunes(runes []rune) []int {
	// start empty word
	word := ""
	numbers := []int{}

	// iterate through runes
	for _, char := range runes {
		word += string(char)

		fmt.Println("word", word)

		// if char is a number, append to numbers array, reset word, and continue to next iteration
		if unicode.IsDigit(char) {
			numbers = append(numbers, int(char-'0'))
			fmt.Println(int(char - '0'))
			word = ""
			continue
		}

		// map through wordToNumber map
		for key, value := range wordToNumber {
			// if word contains key, append value to numbers array, reset word (to current char), and continue to next iteration
			if strings.Contains(word, key) {
				numbers = append(numbers, value)
				fmt.Println(value)
				word = string(char)
			}
		}
	}

	return numbers
}

func main() {
	// open file
	file, err := os.Open("../input.txt")
	defer file.Close()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// start total
	total := 0

	// create scanner
	scanner := bufio.NewScanner(file)

	// while scanner.Scan() is true, keep looping through the file
	for scanner.Scan() {
		fmt.Println(scanner.Text())

		// convert string (scanner.Text()) to array of runes
		chars := []rune(scanner.Text())

		numbers := extractNumbersFromRunes(chars)

		fmt.Println("numbers", numbers)

		if len(numbers) > 0 {
			// get first and last number in the numbersInString array (if there's only one number, it will be the first and last)
			first := numbers[0]
			last := numbers[len(numbers)-1]

			// number = first concatenated with last
			number := strconv.Itoa(first) + strconv.Itoa(last)
			fmt.Println("number", number)

			// convert number to int
			num, err := strconv.Atoi(number)
			if err != nil {
				panic(err)
			}
			// add number to total
			total += num
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Total:", total) // 54578
}
