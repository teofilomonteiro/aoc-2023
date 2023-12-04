package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

var input =`467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

func ourStrMap() []string	{
	return strings.Split(input, "\n");
}

type Number struct {
	value int
	pos []int
}

type Symbol struct {
	value string
	pos int
}

type Line struct {
	numbers []Number
	symbols []Symbol
}


var number, _ = regexp.Compile("\\d+")
var anySymbol, _ = regexp.Compile("[^0-9.]{1}")

func toGame(gameStr []string) []Line {
	line := []Line{}
	for i := 0; i < len(gameStr); i++ {
		foundNumbers := number.FindAllString(gameStr[i], -1)
		numbersPos := number.FindAllStringIndex(gameStr[i], -1)
		numbers := []Number{}
		for index := range foundNumbers {
			nr, _ := strconv.Atoi(foundNumbers[index])

			numbers = append(numbers, Number{value: nr, pos: numbersPos[index]})
		}

		foundSymbols := anySymbol.FindAllString(gameStr[i], -1)

		symbolsPos := anySymbol.FindAllStringIndex(gameStr[i], -1)
		
		symbols := []Symbol{}
		for index := range foundSymbols {
			symbols = append(symbols, Symbol{value: foundSymbols[index], pos: symbolsPos[index][0]})
		}

		line = append(line, Line{numbers: numbers, symbols: symbols})
	}

	return line;
}

func partFunc1(lines []Line) int {
	var sumID = 0;

	for i := 0; i < len(lines); i++ {
		if(len(lines[i].numbers) > 0) {
			for j := 0; j < len(lines[i].numbers); j++ {
				currentNr := lines[i].numbers[j]

				for k := i-1; k <= i+1; k++ {
					
					if(k < 0 || k >= len(lines)) {
						continue;
					}

					lineSymbols := lines[k].symbols

					found := false
					for l := 0; l < len(lineSymbols); l++ {
						if(lineSymbols[l].pos >= currentNr.pos[0]-1 && lineSymbols[l].pos <= currentNr.pos[1]) {
							found = true
							break;
						}
					}

					if(found) {
						sumID = sumID + currentNr.value
						break;
					}
				}
			}
		}
	}

	return sumID
}

func part2Func(lines []Line) int {
	var sumID = 0;

	for i := 0; i < len(lines); i++ {
		if(len(lines[i].symbols) > 0) {
			for j := 0; j < len(lines[i].symbols); j++ {
				currentSymbol := lines[i].symbols[j]

				if(currentSymbol.value != "*") {
					continue;
				}
				numbers := []int{}
				// Search line indexes
				for k := i-1; k <= i+1; k++ {
					if(k < 0 || k >= len(lines)) {
						continue;
					}

					lineNumbers := lines[k].numbers



					for l := 0; l < len(lineNumbers); l++ {
						currentNr := lineNumbers[l]

						if(currentSymbol.pos >= currentNr.pos[0]-1 && currentSymbol.pos <= currentNr.pos[1]) {
							numbers = append(numbers, currentNr.value)
						}
					}
				}
				
				if(len(numbers) == 2) {
					sumID = sumID + numbers[0] * numbers[1]
				}
			}
		}
	}

	return sumID
}

func readFile() []string {
	absfile, _ := filepath.Abs("./day3/input.txt");

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
	// var unparsedGame = ourStrMap()

	game := toGame(unparsedGame)

	part1 := partFunc1(game);
	part2 := part2Func(game);
	fmt.Println("part1:", part1, "part2:", part2)
}