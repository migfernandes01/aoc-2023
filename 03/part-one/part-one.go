package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

// LOGIC: SUMS all numbers in the grid that are adjacent to a special symbol (other than a period).
// ADJAENT: above, below, left, right, above left, above right, below left, below right

func isSpecialSymbol(ch rune) bool {
	// Check if the character is a special symbol excluding numeric digits and period
	return ch != '.' && !unicode.IsDigit(ch)
}

func containsSpecialSymbol(s string) bool {
	for _, ch := range s {
		// Check if the character is a special symbol excluding numeric digits and period
		if isSpecialSymbol(ch) {
			return true
		}
	}
	return false
}

func extract(grid [][]string) []int {
	numbers := []int{}

	// Print the grid for verification
	for i, line := range grid {
		// start number string
		number := ""
		isNumber := false
		check := true

		for j, char := range line {
			// if char is a number, add to number string
			digit := regexp.MustCompile(`^[0-9]$`)
			isNumber = digit.MatchString(char)

			// if char is NOT a number and no need to check, add number to numbers array
			if !isNumber && !check {
				num, err := strconv.Atoi(number)
				if err != nil {
					panic(err)
				}
				numbers = append(numbers, num)
			}

			// if char is NOT a number, reset number string and check
			if !isNumber {
				number = ""
				check = true
			}

			// if char is a number and check is true (needed), check for special symbols around it (mark check as false if there is a special symbol adjacent)
			if isNumber && check {
				if i == 0 {
					// if current char is all the way to the left
					if j == 0 {
						// check RIGHT, BELOW, and BELOW RIGHT for symbols
						if containsSpecialSymbol(grid[i][j+1]) {
							check = false
						}
						if containsSpecialSymbol(grid[i+1][j]) {
							check = false
						}
						if containsSpecialSymbol(grid[i+1][j+1]) {
							check = false
						}
					} else if j == len(line)-1 { // if current char is all the way to the right
						// check LEFT, BELOW, and BELOW LEFT for symbols
						if containsSpecialSymbol(grid[i][j-1]) {
							check = false
						}
						if containsSpecialSymbol(grid[i+1][j]) {
							check = false
						}
						if containsSpecialSymbol(grid[i+1][j-1]) {
							check = false
						}
					} else { // check BELOW, LEFT, RIGHT, BELOW LEFT, and BELOW RIGHT for symbols
						if containsSpecialSymbol(grid[i+1][j]) {
							check = false
						}
						if containsSpecialSymbol(grid[i][j-1]) {
							check = false
						}
						if containsSpecialSymbol(grid[i][j+1]) {
							check = false
						}
						if containsSpecialSymbol(grid[i+1][j-1]) {
							check = false
						}
						if containsSpecialSymbol(grid[i+1][j+1]) {
							check = false
						}
					}
				} else if i == len(grid)-1 { // if current char is all the way down
					// if current char is all the way to the left
					if j == 0 {
						// check RIGHT, ABOVE, and ABOVE RIGHT for symbols
						if containsSpecialSymbol(grid[i][j+1]) {
							check = false
						}
						if containsSpecialSymbol(grid[i-1][j]) {
							check = false
						}
						if containsSpecialSymbol(grid[i-1][j+1]) {
							check = false
						}
					} else if j == len(line)-1 { // if current char is all the way to the right
						// check LEFT, ABOVE, and ABOVE LEFT for symbols
						if containsSpecialSymbol(grid[i][j-1]) {
							check = false
						}
						if containsSpecialSymbol(grid[i-1][j]) {
							check = false
						}
						if containsSpecialSymbol(grid[i-1][j-1]) {
							check = false
						}
					} else { // check ABOVE, LEFT, RIGHT, ABOVE LEFT, and ABOVE RIGHT for symbols
						if containsSpecialSymbol(grid[i-1][j]) {
							check = false
						}
						if containsSpecialSymbol(grid[i][j-1]) {
							check = false
						}
						if containsSpecialSymbol(grid[i][j+1]) {
							check = false
						}
						if containsSpecialSymbol(grid[i-1][j-1]) {
							check = false
						}
						if containsSpecialSymbol(grid[i-1][j+1]) {
							check = false
						}
					}
				} else { // if current char is not all the way up or down
					// if current char is all the way to the left
					if j == 0 {
						// check RIGHT, ABOVE, BELOW, ABOVE RIGHT, and BELOW RIGHT for symbols
						if containsSpecialSymbol(grid[i][j+1]) {
							check = false
						}
						if containsSpecialSymbol(grid[i-1][j]) {
							check = false
						}
						if containsSpecialSymbol(grid[i+1][j]) {
							check = false
						}
						if containsSpecialSymbol(grid[i-1][j+1]) {
							check = false
						}
						if containsSpecialSymbol(grid[i+1][j+1]) {
							check = false
						}
					} else if j == len(line)-1 { // if current char is all the way to the right
						// check LEFT, ABOVE, BELOW, ABOVE LEFT, and BELOW LEFT for symbols
						if containsSpecialSymbol(grid[i][j-1]) {
							check = false
						}
						if containsSpecialSymbol(grid[i-1][j]) {
							check = false
						}
						if containsSpecialSymbol(grid[i+1][j]) {
							check = false
						}
						if containsSpecialSymbol(grid[i-1][j-1]) {
							check = false
						}
						if containsSpecialSymbol(grid[i+1][j-1]) {
							check = false
						}
					} else { // check ABOVE, BELOW, LEFT, RIGHT, ABOVE LEFT, ABOVE RIGHT, BELOW LEFT, and BELOW RIGHT for symbols
						if containsSpecialSymbol(grid[i-1][j]) {
							check = false
						}
						if containsSpecialSymbol(grid[i+1][j]) {
							check = false
						}
						if containsSpecialSymbol(grid[i][j-1]) {
							check = false
						}
						if containsSpecialSymbol(grid[i][j+1]) {
							check = false
						}
						if containsSpecialSymbol(grid[i-1][j-1]) {
							check = false
						}
						if containsSpecialSymbol(grid[i-1][j+1]) {
							check = false
						}
						if containsSpecialSymbol(grid[i+1][j-1]) {
							check = false
						}
						if containsSpecialSymbol(grid[i+1][j+1]) {
							check = false
						}
					}
				}
			}

			// if char is a number, add to number string
			if isNumber {
				number += char
			}
		}

		// if char is a number and no need to check, add number to numbers array
		if isNumber && !check {
			num, err := strconv.Atoi(number)
			if err != nil {
				panic(err)
			}
			numbers = append(numbers, num)
		}
	}
	fmt.Println("numbers", numbers)

	return numbers
}

func main() {
	rows := 140
	file, err := os.Open("../input.txt")
	defer file.Close()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// create scanner
	scanner := bufio.NewScanner(file)

	grid := make([][]string, rows)

	lineCount := 0
	for scanner.Scan() {
		line := scanner.Text()

		// Create a new inner slice for each line
		grid[lineCount] = make([]string, len(line))

		// Loop through each character in the line
		for charPosition, char := range line {
			charString := string(char)
			grid[lineCount][charPosition] = charString
		}

		lineCount++
	}

	total := 0
	numbers := extract(grid)
	for _, num := range numbers {
		total += num
	}
	fmt.Println("Total:", total)
}
