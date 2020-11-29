package builder

import (
	"aicup2020/model"
	"math"
)

/*
	Пакет предоставляет тип строителя.
*/

// TBuilder -- операции со строителем
type TBuilder struct {
	obj      *model.Entity
	food     *model.Entity // Еда, к которой надо идти
	foodDist float64       // Расстояние до еды
}

// New -- созлаёт нового строителя
func New(obj *model.Entity) *TBuilder {
	return &TBuilder{
		obj: obj,
	}
}

// SetFood -- устанавливает объект еды для поглощения
func (sf *TBuilder) SetFood(food *model.Entity) {
	sf.food = food
	deltaX := float64(food.Position.X - sf.obj.Position.X)
	deltaY := float64(food.Position.Y - sf.obj.Position.Y)
	sf.foodDist = math.Sqrt(math.Pow(deltaX, 2) + math.Pow(deltaY, 2))
}

// CheckDistFood -- проверяет дистанцию до еды (если ближе, чем есть -- сохраняет)
func (sf *TBuilder) CheckDistFood(obj *model.Entity) (res bool) {
	deltaX := float64(obj.Position.X - sf.obj.Position.X)
	deltaY := float64(obj.Position.Y - sf.obj.Position.Y)
	foodDist := math.Sqrt(math.Pow(deltaX, 2) + math.Pow(deltaY, 2))
	if foodDist < 1 {
		return true
	}
	if foodDist < sf.foodDist {
		sf.food = obj
		sf.foodDist = foodDist
	}
	return false
}

// Num -- возвращает номер строителя
func (sf *TBuilder) Num() int32 {
	return sf.obj.Id
}

// GetAction -- возвращает действие строителя для текущего тика
func (sf *TBuilder) GetAction() *model.EntityAction {
	act := &model.EntityAction{}
	if sf.food != nil {
		if sf.foodDist < 1 {

		}
		act.MoveAction = &model.MoveAction{
			Target: model.Vec2Int32{
				X: sf.food.Position.X,
				Y: sf.food.Position.Y,
			},
		}
	}
	return act
}
