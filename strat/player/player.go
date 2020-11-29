package player

/*
	Тип представляет статистику игрока
*/

import (
	"aicup2020/model"
)

// TPlayer -- операции с игроком
type TPlayer struct {
	obj *model.Player
}

func New(obj *model.Player) *TPlayer {
	return &TPlayer{
		obj: obj,
	}
}
