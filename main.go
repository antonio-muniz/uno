package main

import (
	"math/rand"
	"time"

	"github.com/antonio-muniz/uno/game"
	"github.com/antonio-muniz/uno/game/player"
	"github.com/antonio-muniz/uno/game/ui"
)

func main() {
	setupRandomizer()

	ui.Message.Welcome()

	uno := createGame()
	winner := uno.Play()

	ui.Message.WinnerFound(winner.Name())
}

func createGame() *game.Game {
	numberOfPlayers := ui.PromptIntegerInRange(2, 8, "How many players in the game? (2-8)")
	humanPlayerName := ui.PromptString("What's your name?")

	players := player.CreatePlayers(numberOfPlayers, humanPlayerName)

	return game.New(players)
}

func setupRandomizer() {
	rand.Seed(time.Now().UnixNano())
}
