package entity

type Code interface {
	GetGame() Game
	GetCode() string
	GetRewards() []string
}
