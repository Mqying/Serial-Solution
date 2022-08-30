package mock

import (
	"math/rand"
	"time"

	"github.com/silverswords/dielectric/server/model"
	"github.com/silverswords/dielectric/server/util"
)

func GenerateWaterRecord(t time.Time) *model.WaterRecord {
	randFlag := rand.Float64()

	quantity := util.Round(rand.Float64() * 900000 + 100000, 1)
	ratio1 := util.Round(rand.Float64() * 900000 + 100000, 1)
	ratio2 := util.Round(rand.Float64() * 9000 + 1000, 1)

	if(randFlag >= 0.5) {
		return &model.WaterRecord{
			DetectionTime: t,
			Quantity:      quantity,
			Ratio1:        0,
			Ratio2:        ratio2,
		}
	}else{
		return &model.WaterRecord{
			DetectionTime: t,
			Quantity:      quantity,
			Ratio1:        ratio1,
			Ratio2:        0,
		}
	}
}