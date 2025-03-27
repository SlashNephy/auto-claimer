package entity

import (
	"fmt"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type Game string

const (
	GameHonkaiStarrail Game = "honkai_starrail"
)

func (g Game) LocalizeMessage() *i18n.Message {
	switch g {
	case GameHonkaiStarrail:
		return &i18n.Message{
			ID:    "Game.HonkaiStarrail",
			Other: "Honkai: Star Rail",
		}
	default:
		panic(fmt.Errorf("unexpected game: %s", g))
	}
}
