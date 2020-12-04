// Package poolbuilder -- пул строителей своей базы
package poolbuilder

import (
	"aicup2020/model"
	"aicup2020/strat/builder"
	"aicup2020/strat/player"
	"log"
)

// TPoolBuilder -- пул строителей
type TPoolBuilder struct {
	pool         []*builder.TBuilder
	baseX, baseY int32
}

var (
	// PoolBuilder -- пул строителей
	PoolBuilder *TPoolBuilder
)

func New(observe *model.PlayerView, baseX, baseY int32) (pool *TPoolBuilder) {
	pool = &TPoolBuilder{
		pool:  make([]*builder.TBuilder, 0),
		baseX: baseX,
		baseY: baseY,
	}
	pool.getBuilders(observe)
	pool.findFoot(observe)
	return
}

// Ищет всех строителей
func (sf *TPoolBuilder) getBuilders(observe *model.PlayerView) {
	for _, obj := range observe.Entities {
		if obj.PlayerId == nil {
			continue
		}
		objId := *obj.PlayerId
		if objId != player.Player.Id() {
			continue
		}
		// Получить своего строителя
		if obj.EntityType != model.EntityTypeBuilderUnit {
			continue
		}
		builder := builder.New(obj, sf.baseX, sf.baseY)
		sf.addBuilder(builder)
	}
	log.Printf("getBuilders(): len=%v\n", len(sf.Builders()))
}

// Ищет ближайшую еду к строителю
func (sf *TPoolBuilder) findFoot(observe *model.PlayerView) {
	for _, builder := range sf.Builders() {
		builder.ResetFood()
		for _, obj := range observe.Entities {
			if obj.EntityType != model.EntityTypeResource {
				continue
			}
			// Проверить дальность до строителя
			builder.CheckDistFood(&obj)
		}
	}
}

// Aдобавляет в пул строителя
func (sf *TPoolBuilder) addBuilder(builder *builder.TBuilder) {
	if builder == nil {
		log.Panicf("TPoolBuilder.Add(): builder==nil\n")
	}
	sf.pool = append(sf.pool, builder)
}

// GetPoolActions -- опрашивает строителей, возвращает их пул дейтсвий
func (sf *TPoolBuilder) GetPoolActions() map[int32]model.EntityAction {
	poolAct := make(map[int32]model.EntityAction)
	for _, builder := range sf.pool {
		//builder.Report()
		poolAct[builder.Num()] = *builder.GetAction()
	}
	return poolAct
}

// Builders -- возвращает список всех строителей
func (sf *TPoolBuilder) Builders() []*builder.TBuilder {
	return sf.pool
}
