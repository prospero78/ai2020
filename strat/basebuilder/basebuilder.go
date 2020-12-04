// Package basebuilder -- тип для базы строительных юнитов
package basebuilder

import (
	"aicup2020/model"
	"aicup2020/strat/player"
	"aicup2020/strat/poolbuilder"
	"log"
	"math/rand"
)

// TBaseBuilder -- операции со зданием-базой
type TBaseBuilder struct {
	base model.Entity
}

var (
	// Base -- объект базы
	Base *TBaseBuilder
)

// New -- возвращает новый объект базы
func New(obj model.Entity) (base *TBaseBuilder) {
	base = &TBaseBuilder{
		base: obj,
	}
	return base
}

// GetCoord -- возвращает координаты базы
func (sf *TBaseBuilder) GetCoord() (x int32, y int32) {
	return sf.base.Position.X, sf.base.Position.Y
}

// GetAct -- возвращает действие
func (sf *TBaseBuilder) GetAct() (act model.EntityAction) {
	x := sf.base.Position.X - 10 + rand.Int31n(50)
	if x < 0 {
		x = 2
	}
	y := sf.base.Position.Y - 10 + rand.Int31n(50)
	if y < 0 {
		y = 2
	}
	if player.Player.Resource() > 20 && len(poolbuilder.PoolBuilder.Builders()) <= 10 {
		log.Printf("TBaseBuilder.GetAct(): build unit.builder baseX=%v baseY=%v x=%v y=%v\n", sf.base.Position.X, sf.base.Position.Y, x, y)
		act = model.EntityAction{
			BuildAction: &model.BuildAction{
				EntityType: model.EntityTypeBuilderUnit,
				Position: model.Vec2Int32{
					X: x,
					Y: y,
				},
			},
		}
	}
	return act
}

// Id -- возвращает свой Id
func (sf *TBaseBuilder) Id() int32 {
	return sf.base.Id
}


// GetX -- возвращает координату Х базы строительства строителей
func (sf *TBaseBuilder) GetX() int32 {
	return sf.base.Position.X
}

// GetY -- возвращает координату Y базы строительства строителей
func (sf *TBaseBuilder) GetY() int32 {
	return sf.base.Position.Y
}