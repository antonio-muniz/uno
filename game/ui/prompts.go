package ui

import (
	"fmt"
	"strings"

	"github.com/antonio-muniz/uno/game/card"
	"github.com/antonio-muniz/uno/game/card/color"
)

func promptString(message string) string {
	for {
		Println(message)
		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			Println("Invalid input")
			continue
		}
		return input
	}
}

func promptLowercaseString(message string) string {
	input := promptString(message)
	return strings.ToLower(input)
}

func promptUppercaseString(message string) string {
	input := promptString(message)
	return strings.ToUpper(input)
}

func PromptCardSelection(cards []card.Card) card.Card {
	runeSequence := runeSequence{}
	cardOptions := make(map[string]card.Card)
	for _, card := range cards {
		label := string(runeSequence.next())
		cardOptions[label] = card
	}

	cardSelectionLines := []string{"Select a card to play:"}
	for label, card := range cardOptions {
		cardSelectionLines = append(cardSelectionLines, fmt.Sprintf("%s (enter %s)", card, label))
	}
	cardSelectionMessage := strings.Join(cardSelectionLines, "\n")

	for {
		selectedLabel := promptUppercaseString(cardSelectionMessage)
		selectedCard, found := cardOptions[selectedLabel]
		if !found {
			Printfln("No card assigned to '%s'", selectedLabel)
			continue
		}
		return selectedCard
	}
}

func PromptColor() color.Color {
	colorMessage := fmt.Sprintf(
		"Select a color: '%s', '%s', '%s' or '%s'?",
		color.Red,
		color.Yellow,
		color.Green,
		color.Blue,
	)
	for {
		colorName := promptLowercaseString(colorMessage)
		chosenColor, err := color.ByName(colorName)
		if err != nil {
			Printfln("Unknown color '%s'", colorName)
			continue
		}
		return chosenColor
	}
}
