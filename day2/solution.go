package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type RoundCubes struct {
	blue int
	red int
	green int
}

type Game struct {
	nr int
	rounds []RoundCubes
}
// 12 red cubes, 13 green cubes, and 14 blue cubes
var bag = map[string]int{
	"blue": 14,
	"red": 12,
	"green": 13,
}

var input =`Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

func simulated() []string	{
	return strings.Split(input, "\n");
}

func toGame(gameStr []string) []Game {
	games := []Game{}
	for i := 0; i < len(gameStr); i++ {
		
		gameInfo := strings.Split(gameStr[i], ": ")
		gameId, _ := strconv.Atoi(strings.Replace(gameInfo[0], "Game ", "", 1));
		rounds := []RoundCubes{}
		game := Game{nr: gameId, rounds: rounds}
		
		roundsStr := strings.Split(gameInfo[1], "; ")

		for j := 0; j < len(roundsStr); j++ {
			roundStr := strings.Split(roundsStr[j], ", ")
			round := RoundCubes{blue: 0, red: 0, green: 0}
			for k := 0; k < len(roundStr); k++ {

				move := strings.Split(roundStr[k], " ")
				if(move[1] == "blue") {
					round.blue, _ = strconv.Atoi(move[0])
				} else if(move[1] == "red") {
					round.red, _ = strconv.Atoi(move[0])
				} else if(move[1] == "green") {
					round.green, _ = strconv.Atoi(move[0])
				}
			}
			game.rounds = append(game.rounds, round)
		}

		games = append(games, game)
	}

	return games
}

func sol1(game []Game) int {
	var sumID = 0;
	for i := 0; i < len(game); i++ {
		isPossible := true;
		for j := 0; j < len(game[i].rounds); j++ {
			if(bag["blue"] < game[i].rounds[j].blue) || (bag["red"] < game[i].rounds[j].red) || (bag["green"] < game[i].rounds[j].green) {
				isPossible = false;
				break;
			}
		}

		if(isPossible) {
			sumID=sumID+game[i].nr;
		}
	}
	return sumID
}

func sol2(game []Game) int {
	var sumID = 0;

	for i := 0; i < len(game); i++ {
		minRound := RoundCubes{blue: game[i].rounds[0].blue, red: game[i].rounds[0].red, green: game[i].rounds[0].green}
		
		for j := 1; j < len(game[i].rounds); j++ {
			minRound.blue = max(minRound.blue, game[i].rounds[j].blue)
			minRound.red = max(minRound.red, game[i].rounds[j].red)
			minRound.green = max(minRound.green, game[i].rounds[j].green)
		}
		sumID=sumID+(minRound.blue*minRound.red*minRound.green);
	}
	return sumID
}

func readFile() []string {
	absfile, _ := filepath.Abs("./day2/input.txt");

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
	// var unparsedGame = simulated()

	game := toGame(unparsedGame)

	solution1 := sol1(game);
	solution2 := sol2(game);
	fmt.Println("solution1:", solution1, "solution2:", solution2)
}