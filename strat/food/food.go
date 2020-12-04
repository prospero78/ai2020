package food

import (
	"aicup2020/model"
	"math"
)

/*
	Пакет предоставляет тип еды. Хранит в себе свойства еды.
*/

// TFood -- операции с едой
type TFood struct {
	food *model.Entity // Объект еды
	dist float64       // Дистанция от юнита до еды
}

func New(food *model.Entity, baseX, baseY int32) *TFood {
	deltaX := float64(food.Position.X - baseX)
	deltaY := float64(food.Position.Y - baseY)
	foodDist := math.Sqrt(deltaX*deltaX + deltaY*deltaY)
	fd := &TFood{
		food: food,
		dist: foodDist,
	}
	return fd
}

// Dist -- возвращает хранимую дистанцию до еды
func (sf *TFood) Dist() float64 {
	return sf.dist
}

// PosX -- возвращает хранимую позицию X
func (sf *TFood) PosX() int32 {
	return sf.food.Position.X
}

// PosY -- возвращает хранимую позицию Y
func (sf *TFood) PosY() int32 {
	return sf.food.Position.Y
}

// GetType -- возвращает объект еды
func (sf *TFood) GetType() model.EntityType {
	return sf.food.EntityType
}

// IsEmpty -- не установлена еда для робота
func (sf *TFood) IsEmpty() bool {
	return sf.food == nil
}

// Id -- возвращает свой id
func (sf *TFood) Id() *int32 {
	return &sf.food.Id
}
