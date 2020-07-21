package event_test

import (
	"testing"

	"github.com/antonio-muniz/uno/pkg/card"
	"github.com/antonio-muniz/uno/pkg/card/color"
	"github.com/antonio-muniz/uno/pkg/event"
	"github.com/stretchr/testify/require"
)

func TestCardPlayed(t *testing.T) {
	listenerOne := &listener{}
	listenerTwo := &listener{}

	event.CardPlayed.AddListener(listenerOne)
	event.CardPlayed.AddListener(listenerTwo)

	payloads := []event.CardPlayedPayload{
		{
			PlayerName: "Someone",
			Card:       card.NewWildCard(),
		},
		{
			PlayerName: "Somebody",
			Card:       card.NewDrawTwoCard(color.Green),
		},
	}

	for _, payload := range payloads {
		event.CardPlayed.Emit(payload)
	}

	require.Equal(t, payloads, listenerOne.receivedPayloads)
	require.Equal(t, payloads, listenerTwo.receivedPayloads)
}

type listener struct {
	receivedPayloads []event.CardPlayedPayload
}

func (l *listener) OnCardPlayed(payload event.CardPlayedPayload) {
	l.receivedPayloads = append(l.receivedPayloads, payload)
}
