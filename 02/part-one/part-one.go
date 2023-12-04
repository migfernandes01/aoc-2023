package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// LOGIC:
// reads file, extracts game number and buckets, and returns a Game struct
// then it returs the ids of the possible games, we then loop through the ids and add them to a total
// RULES:
// possible games are games where the buckets have less than or equal to the max number of marbles (red <= 12, blue <= 13, green <= 14)
// AND have no blocking buckets (buckets where one of the colors has more than the max number of marbles)

type Bucket struct {
	red   int
	green int
	blue  int
}

type Game struct {
	num     int
	buckets []Bucket
}

func extractGame(game string) Game {
	game2 := Game{}
	gameBuckets := []Bucket{}

	// extract game number
	num := 0
	gameNumber := game[:strings.IndexByte(game, ':')]
	re := regexp.MustCompile("[0-9]+")
	numInGameNumber := re.FindAllString(gameNumber, -1)
	// convert numInGameNumber (string) to int
	num, err := strconv.Atoi(numInGameNumber[0])
	if err != nil {
		panic(err)
	}

	// extract buckets
	buckets := strings.Split(game[strings.IndexByte(game, ':')+1:], ";")

	// loop through buckets
	for _, bucket := range buckets {
		red := 0
		green := 0
		blue := 0

		// split on ","
		colorsAndNumbers := strings.Split(bucket, ",")
		// loop through colorsAndNumbers
		for _, colorAndNumber := range colorsAndNumbers {
			colorAndNumberSplit := strings.Split(colorAndNumber, " ")
			if colorAndNumberSplit[2] == "red" {
				// convert colorAndNumberSplit[0] (string) to int
				numberForColor, err := strconv.Atoi(colorAndNumberSplit[1])
				if err != nil {
					panic(err)
				}
				red += numberForColor
			}
			if colorAndNumberSplit[2] == "blue" {
				// convert colorAndNumberSplit[0] (string) to int
				numberForColor, err := strconv.Atoi(colorAndNumberSplit[1])
				if err != nil {
					panic(err)
				}
				blue += numberForColor
			}
			if colorAndNumberSplit[2] == "green" {
				// convert colorAndNumberSplit[0] (string) to int
				numberForColor, err := strconv.Atoi(colorAndNumberSplit[1])
				if err != nil {
					panic(err)
				}
				green += numberForColor
			}
		}

		gameBuckets = append(gameBuckets, Bucket{
			red:   red,
			green: green,
			blue:  blue,
		})
	}

	game2.num = num
	game2.buckets = gameBuckets

	return game2
}

func determinePossibleGames(games []Game, maxRed int, maxGreen int, maxBlue int) []int {
	possibleGameIds := []int{}

	// loop through games
	for _, game := range games {
		found := false
		blocking := false
		for _, bucket := range game.buckets {
			if bucket.red <= maxRed && bucket.blue <= maxBlue && bucket.green <= maxGreen {
				found = true
			} else {
				blocking = true
			}
		}
		// if found and not blocking, append game.num to possibleGameIds
		if found && !blocking {
			possibleGameIds = append(possibleGameIds, game.num)
		}
	}
	return possibleGameIds
}

func main() {
	// open file
	file, err := os.Open("../input.txt")
	defer file.Close()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// create scanner
	scanner := bufio.NewScanner(file)

	games := []Game{}

	// while scanner.Scan() is true, keep looping through the file
	for scanner.Scan() {
		// get games with correct format (game number, buckets)
		game := extractGame(scanner.Text())
		fmt.Println("game", game)
		games = append(games, game)
	}

	// determine possible games (red <= 12, blue <= 13, green <= 14)
	possibleGameIds := determinePossibleGames(games, 12, 13, 14)

	// get total
	total := 0
	for _, id := range possibleGameIds {
		total += id
	}
	fmt.Println("total", total)
}
