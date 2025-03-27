package hoyoverse

import "github.com/SlashNephy/auto-claimer/domain/entity"

type Code struct {
	Game    entity.Game
	Code    string
	Rewards []string
}

func (c *Code) GetGame() entity.Game {
	return c.Game
}

func (c *Code) GetCode() string {
	return c.Code
}

func (c *Code) GetRewards() []string {
	return c.Rewards
}

var _ entity.Code = new(Code)
