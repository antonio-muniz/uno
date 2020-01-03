package game

import (
	"fmt"
	"strings"

	"github.com/antonio-muniz/uno/game/card"
	"github.com/antonio-muniz/uno/game/card/color"
)

type GameState struct {
	currentColor      color.Color
	lastPlayedCard    card.Card
	currentPlayerHand []card.Card
	playerSequence    []string
	playerHandCounts  map[string]int
}

func (s GameState) CurrentColor() color.Color {
	return s.currentColor
}

func (s GameState) LastPlayedCard() card.Card {
	return s.lastPlayedCard
}

func (s GameState) CurrentPlayerHand() []card.Card {
	return s.currentPlayerHand
}

func (s GameState) PlayerSequence() []string {
	return s.playerSequence
}

func (s GameState) PlayerHandCounts() map[string]int {
	return s.playerHandCounts
}

func (s GameState) String() string {
	var lines []string
	lines = append(lines, fmt.Sprintf("Current color: %s", s.currentColor))
	lines = append(lines, fmt.Sprintf("Last played card: %s", s.lastPlayedCard))

	var playerStatuses []string
	for _, playerName := range s.playerSequence {
		playerStatus := fmt.Sprintf("%s (%d card(s))", playerName, s.playerHandCounts[playerName])
		playerStatuses = append(playerStatuses, playerStatus)
	}
	lines = append(lines, fmt.Sprintf("Turn order: %s", strings.Join(playerStatuses, ", ")))

	lines = append(lines, fmt.Sprintf("Your hand: %s", s.currentPlayerHand))

	return strings.Join(lines, "\n")
}
