package builder

import (
	"aicup2020/model"
	"aicup2020/strat/food"
	"aicup2020/strat/player"
	"log"
	"os"
)

/*
	Пакет предоставляет тип строителя.
*/

// TBuilder -- операции со строителем
type TBuilder struct {
	unit         model.Entity
	food         *food.TFood // Еда, к которой надо идти
	player       *player.TPlayer
	gamerId      int32
	baseX, baseY int32
}

// New -- создаёт нового строителя
func New(unit model.Entity, baseX, baseY int32) *TBuilder {
	id := *unit.PlayerId
	return &TBuilder{
		unit:    unit,
		player:  player.Player,
		gamerId: id,
		baseX:   baseX,
		baseY:   baseY,
	}
}

// ResetFood -- сбрасывает объект еды для поглощения
func (sf *TBuilder) ResetFood() {
	sf.food = nil
}

// CheckDistFood -- проверяет дистанцию до еды (если ближе, чем есть -- сохраняет)
func (sf *TBuilder) CheckDistFood(nam *model.Entity) {
	if sf.food == nil {
		sf.food = food.New(nam, sf.baseX, sf.baseY)
	}
	foodNext := food.New(nam, sf.baseX, sf.baseY)
	if foodNext.Dist() < sf.food.Dist() {
		sf.food = foodNext
	}
}

// Num -- возвращает номер строителя
func (sf *TBuilder) Num() int32 {
	return sf.unit.Id
}

// GetAction -- возвращает действие строителя для текущего тика
func (sf *TBuilder) GetAction() *model.EntityAction {
	act := &model.EntityAction{}
	switch sf.food == nil {
	case true: // Нет еды
		act.MoveAction = &model.MoveAction{
			Target: model.Vec2Int32{
				X: sf.unit.Position.X + 1,
				Y: sf.unit.Position.Y + 1,
			},
		}
	case false: // Есть еда
		act.MoveAction = &model.MoveAction{
			Target: model.Vec2Int32{
				X: sf.food.PosX(),
				Y: sf.food.PosY(),
			},
		}
		if !sf.food.IsEmpty() {
			act.AttackAction = &model.AttackAction{
				AutoAttack: &model.AutoAttack{
					ValidTargets: make([]model.EntityType, 0),
				},
				Target: sf.food.Id(),
			}
			act.AttackAction.AutoAttack.ValidTargets = append(act.AttackAction.AutoAttack.ValidTargets, sf.food.GetType())
		}
	}
	return act
}

func (sf *TBuilder) Report() {
	if os.Getenv("LOCAL_DEBUG") == "" {
		return
	}
	if sf.food == nil {
		log.Printf("TBuilder.CheckDistFood(): id=%v type=%v gamerId=%+v player=%v food=nil\n", sf.unit.Id, sf.unit.EntityType, sf.gamerId, sf.player.Id())
		return
	}
	log.Printf("TBuilder.CheckDistFood():  id=%v type=%v gamerId=%+v player=%v dist=%0.1f unit.pos=%v:%v food.pos=%v:%v\n",
		sf.unit.Id,
		sf.unit.EntityType,
		sf.gamerId,
		sf.player.Id(),
		sf.food.Dist(),
		sf.unit.Position.X, sf.unit.Position.Y,
		sf.food.PosX(), sf.food.PosY())
}
