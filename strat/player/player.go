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

var (
	// Player -- ссылка на себя самого игрока
	Player *TPlayer
)

func New(obj *model.Player) *TPlayer {
	return &TPlayer{
		obj: obj,
	}
}

// Id -- возвращает собственный номер
func (sf *TPlayer) Id() int32 {
	return sf.obj.Id
}

// Score -- возвращает собственные очки
func (sf *TPlayer) Score() int32 {
	return sf.obj.Score
}

// Resource -- возвращает собственные ресурсы
func (sf *TPlayer) Resource() int32 {
	return sf.obj.Resource
}
