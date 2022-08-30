package mock

import (
	"math/rand"
	"time"

	"github.com/silverswords/dielectric/server/model"
	"github.com/silverswords/dielectric/server/util"
)

func GenerateAcidRecord(t time.Time) *model.AcidRecord {
	var items [6]float64

	for i := 0; i < 6; i++ {
		items[i] = util.Round(rand.Float64() * 9 + 1, 1)
	}

	return &model.AcidRecord{
		DetectionTime: t,
		Items:         items,
	}
}
