package mock

import (
	"math/rand"
	"time"

	"github.com/silverswords/dielectric/server/model"
	"github.com/silverswords/dielectric/server/util"
)

func GenerateFlashRecord(t time.Time) *model.FlashRecord {	
	id := util.GenerateInt(100000)
	
	return &model.FlashRecord{
		DetectionTime:	t,
		Id:				id,
		Pretemp: 		util.Round(rand.Float64() * 900 + 100, 0),
		Pressure: 		util.Round(rand.Float64() * 900 + 100, 1),
		Pointtemp: 		util.Round(rand.Float64() * 900 + 100, 1),
	}
}