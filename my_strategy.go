package main

import (
	. "aicup2020/model"
	"aicup2020/strat"
	//"log"
)

type MyStrategy struct{}

// NewMyStrategy -- возвращает новый объект стратегии
func NewMyStrategy() MyStrategy {
	return MyStrategy{}
}

//  Возвращает действие на каждый тик. Здесь надо реализовать стратегию.
func (strategy MyStrategy) getAction(observe PlayerView, debugInterface *DebugInterface) Action {
	act := Action{
		EntityActions: strat.MakeTik(&observe),
	}
	return act

	// return Action{
	// 	EntityActions: make(map[int32]EntityAction),
	// }
}

// Позволяет при каждом тике отлаживать стратегию.
func (strategy MyStrategy) debugUpdate(playerView PlayerView, debugInterface DebugInterface) {
	//debugInterface.Send(DebugCommandClear{})
	//debugInterface.GetState()
}
