package strat

/*
	Пакет предоставляет стратегию для игры
*/

import (
	"aicup2020/model"
	"aicup2020/strat/basebuilder"
	"aicup2020/strat/player"
	"aicup2020/strat/poolbuilder"
	"log"
)

var (
	myId        int32
)

// MakeTik -- выполняет всю работу в тике
func MakeTik(observe *model.PlayerView) map[int32]model.EntityAction {
	getSelf(observe)
	getBase(observe)
	poolbuilder.PoolBuilder = poolbuilder.New(observe, basebuilder.Base.GetX(), basebuilder.Base.GetY())
	infoPlayer(observe)
	return getAction()
}

func getBase(observe *model.PlayerView) {
	for _, obj := range observe.Entities {
		if obj.PlayerId == nil {
			continue
		}
		if *obj.PlayerId != myId {
			continue
		}
		if obj.EntityType != model.EntityTypeBuilderBase {
			continue
		}
		// Получить ссылку на свою базу
		basebuilder.Base = basebuilder.New(obj)
		//log.Printf("baseBuilder=%v ", basebuilder.Base.Id())
		return
	}
}

func getSelf(observe *model.PlayerView) {
	myId = observe.MyId
	for _, obj := range observe.Players {
		if obj.Id != myId {
			continue
		}
		// Получить ссылку на себя
		player.Player = player.New(&obj)
		return
	}
}

// GetAction -- возвращает, что надо делать
func getAction() map[int32]model.EntityAction {
	poolAct := make(map[int32]model.EntityAction)
	poolActBuilder := poolbuilder.PoolBuilder.GetPoolActions()
	for key, act := range poolActBuilder {
		poolAct[key] = act
	}
	poolActBuilder[basebuilder.Base.Id()] = basebuilder.Base.GetAct()
	return poolAct
}

// Выводит статистику по игрокам
func infoPlayer(observe *model.PlayerView) {
	log.Printf("infoPlayer(): id=%v score=%v resource=%v", myId, player.Player.Score(), player.Player.Resource())
	// for _, pl := range observe.Players {
	// 	log.Printf("infoPlayer(): id=%v score=%v resource=%v", pl.Id, pl.Score, pl.Resource)
	// }
}
