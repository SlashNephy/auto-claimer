package entity

import (
	"fmt"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type Game string

const (
	GameHonkaiStarrail  Game = "honkai_starrail"
	GameGenshinImpact   Game = "genshin_impact"
	GameZenlessZoneZero Game = "zenless_zone_zero"
)

func (g Game) LocalizeMessage() *i18n.Message {
	switch g {
	case GameHonkaiStarrail:
		return &i18n.Message{
			ID:    "Game.HonkaiStarrail",
			Other: "Honkai: Star Rail",
		}
	case GameGenshinImpact:
		return &i18n.Message{
			ID:    "Game.GenshinImpact",
			Other: "Genshin Impact",
		}
	case GameZenlessZoneZero:
		return &i18n.Message{
			ID:    "Game.ZenlessZoneZero",
			Other: "Zenless Zone Zero",
		}
	default:
		panic(fmt.Errorf("unexpected game: %s", g))
	}
}
