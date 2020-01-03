package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/antonio-muniz/uno/game"
	"github.com/antonio-muniz/uno/game/player"
	"github.com/antonio-muniz/uno/game/ui"
)

var botNames = []string{
	"Ashe",
	"Braum",
	"Caitlyn",
	"Darius",
	"Elise",
	"Fizz",
	"Irelia",
}

func main() {
	setupRandomizer()

	ui.Message.Welcome()

	uno := createGame()
	winner := uno.Play()

	ui.Message.WinnerFound(winner.Name())
}

func createGame() *game.Game {
	var (
		numberOfPlayers int
		humanPlayerName string
	)

	ui.Println("How many players in the game? (2-8)")
	_, err := fmt.Scanln(&numberOfPlayers)
	if err != nil || numberOfPlayers < 2 || numberOfPlayers > 8 {
		panic("number of players must be between 2 and 8")
	}

	ui.Println("What's your name?")
	_, err = fmt.Scanln(&humanPlayerName)
	if err != nil || humanPlayerName == "" {
		panic("your name cannot be empty")
	}

	players := createPlayers(numberOfPlayers, humanPlayerName)

	return game.New(players)
}

func createPlayers(numberOfPlayers int, humanPlayerName string) []game.Player {
	players := make([]game.Player, 0, numberOfPlayers)
	players = append(players, player.NewHumanPlayer(humanPlayerName))
	players = append(players, generateBots(numberOfPlayers-1)...)
	return players
}

func generateBots(amount int) []game.Player {
	rand.Shuffle(len(botNames), func(i int, j int) { botNames[i], botNames[j] = botNames[j], botNames[i] })
	bots := make([]game.Player, 0, amount)
	for _, botName := range botNames[:amount] {
		bots = append(bots, player.NewGoodPlayer(botName))
	}
	return bots
}

func setupRandomizer() {
	rand.Seed(time.Now().UnixNano())
}
