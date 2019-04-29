package l1v

import (
	"gitlab.com/akita/akita"
	"gitlab.com/akita/util"
	"gitlab.com/akita/util/akitaext"
)

// A Cache is a customized L1 cache the for R9nano GPUs.
type Cache struct {
	*akitaext.TickingComponent

	TopPort     akita.Port
	BottomPort  akita.Port
	ControlPort akita.Port

	dirBuf   util.Buffer
	bankBufs []util.Buffer

	coalesceStage    *coalescer
	directoryStage   *directory
	bankStages       []*bankStage
	parseBottomStage *bottomParser
	respondStage     *respondStage

	transactions             []*transaction
	postCoalesceTransactions []*transaction
}

// Tick update the state of the cache
func (c *Cache) Tick(now akita.VTimeInSec) bool {
	madeProgress := false

	madeProgress = c.respondStage.Tick(now) || madeProgress
	madeProgress = c.parseBottomStage.Tick(now) || madeProgress
	for _, bs := range c.bankStages {
		madeProgress = bs.Tick(now) || madeProgress
	}
	madeProgress = c.directoryStage.Tick(now) || madeProgress
	madeProgress = c.coalesceStage.Tick(now) || madeProgress

	return madeProgress
}
