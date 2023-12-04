package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

var input =`Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`

func inputMap() []string	{
	return strings.Split(input, "\n");
}

type Card struct {
	left []int
	right []int
}

func toIntList(lstStr []string) []int {
	lst := []int{}
	for _, str := range lstStr {
		nr, _ := strconv.Atoi(str)
		lst = append(lst, nr)
	}
	return lst
}

func pow(x int) int {
	return int(math.Pow(float64(2), float64(x)))
}


var numberRegx, _ = regexp.Compile("\\d+")

func mapToStruct(inputStr []string) []Card {
	cards := []Card{}
	for i := 0; i < len(inputStr); i++ {
		lineInfo := strings.Split(inputStr[i], ": ")
		cardsInfo := strings.Split(lineInfo[1], " | ")

		winningNumbers := numberRegx.FindAllString(cardsInfo[0],-1)
		numbers := numberRegx.FindAllString(cardsInfo[1],-1)

		cards = append(cards, Card{left: toIntList(winningNumbers), right: toIntList(numbers)})
	}

	return cards;
}

func partFunc1(cards []Card) int {
	var totalSum = 0;

	for _, card := range cards {
		var cardSum = 0;
		for _, winningNumber := range card.left {
			for _, number := range card.right {
				if(winningNumber == number) {
					cardSum ++;
				}
			}
		}
		if(cardSum > 0) {
			totalSum += pow(cardSum-1);
		}
	}

	return totalSum
}

func part2Func(cards []Card) int {

	cardsTotal := make([]int, len(cards));

	for index := range cardsTotal {
		cardsTotal[index] = 1;
	}


	for cardNr, card := range cards {
		var cardSum = 0;

		for _, winningNumber := range card.left {
			for _, number := range card.right {
				if(winningNumber == number) {
					cardSum++;
				}
			}
		}

		if(cardSum > 0) {
			for i := 1; i < cardSum+1; i++ {
				cardsTotal[cardNr + i] = cardsTotal[cardNr] + cardsTotal[cardNr + i];
			}
			
		}
	}
	

	var totalSum = 0;
	for _, cardTotal := range cardsTotal {
		totalSum += cardTotal;
	}
	return totalSum;
}

func readFile() []string {
	absfile, _ := filepath.Abs("./day4/input.txt");

	file, err := os.Open(absfile)
	
	if err != nil {
		log.Fatal(err)
	}
	
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var unparsedGame []string

	for scanner.Scan() {
		unparsedGame = append(unparsedGame, scanner.Text())
	}

	return unparsedGame
}

func main() {
	var unparsedGame = readFile()
	// var unparsedGame = inputMap()

	game := mapToStruct(unparsedGame)

	part1 := partFunc1(game);
	part2 := part2Func(game);
	fmt.Println("part1:", part1, "part2:", part2)
}