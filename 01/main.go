package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	// open file
	file, err := os.Open("input.txt")
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

		// create an array of numbers in the string
		var numbers []int
		for _, char := range scanner.Text() {
			if unicode.IsDigit(char) {
				numbers = append(numbers, int(char-'0'))
			}
		}

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

	fmt.Println("Total:", total)
}
