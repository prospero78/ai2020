package strat

/*
	Пакет предоставляет стратегию для игры
*/

import (
	"aicup2020/model"
	"aicup2020/strat/builder"
	"log"
)

var (
	Builder *builder.TBuilder
	myId    int32
)

// SetBuilder -- устанавливает своего строителя на каждом тике
func SetBuilder(observe *model.PlayerView) (res bool) {
	myId = observe.MyId
	log.Printf("self=%v ", myId)
	getPlayer(observe)
	for _, obj := range observe.Entities {
		if obj.PlayerId == nil {
			continue
		}
		if *obj.PlayerId != myId {
			continue
		}
		// Получить своего строителя
		switch obj.EntityType {
		case model.EntityTypeBuilderUnit:
			Builder = builder.New(&obj)
		default:
			log.Printf("SetBuilder(): unknown type obj(%v)\n", obj.EntityType)
		}
	}
	if Builder == nil {
		log.Panicf("SetBuilder(): builder==nil\n")
	}
	return Builder != nil
}

// FindFood -- ищет ближайшую еду к строителю
func FindFoot(observe *model.PlayerView) {
	var foodNear *model.Entity
	for _, obj := range observe.Entities {
		if obj.EntityType == model.EntityTypeResource {
			if foodNear == nil {
				foodNear = &obj
				Builder.SetFood(&obj)
			}
		}
		// Проверить дальность до строителя
		if Builder.CheckDistFood(&obj) {
			break
		}
	}
}

// GetAction -- возвращает, что надо делать
func GetAction() *model.Action {
	poolAct := make(map[int32]model.EntityAction)
	poolAct[Builder.Num()] = *Builder.GetAction()
	act := &model.Action{
		EntityActions: poolAct,
	}
	return act
}

// Выводит статистику по игрокам
func getPlayer(observe *model.PlayerView) {
	for _, pl := range observe.Players {
		log.Printf("getPlayer(): id=%v score=%v resource=%v", pl.Id, pl.Score, pl.Resource)
	}
}
