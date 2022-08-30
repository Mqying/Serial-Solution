package mock

import (
	"math/rand"
	"time"

	"github.com/silverswords/dielectric/server/model"
	"github.com/silverswords/dielectric/server/util"
)

func GenerateDielectricRecord(t time.Time) *model.DielectronRecord {
	var items [10]float64
	sum := 0.0
	for i := 0; i < 10; i++ {
		items[i] = util.Round(rand.Float64() * 90 + 10, 1)
		sum += items[i]
	}

	return &model.DielectronRecord{
		DetectionTime: t,
		Average:       util.FixPrecision(sum / 10),
		Items:         items,
	}
}
