package mock

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/silverswords/dielectric/server/model"
	"github.com/silverswords/dielectric/server/zlog"
)

type MockController struct {
	db *sql.DB

	index        int
	acidRecords  []*model.AcidRecord
	waterRecords []*model.WaterRecord
	dieleRecords []*model.DielectronRecord
	flashRecords []*model.FlashRecord
}

func NewMockController(db *sql.DB) *MockController {
	controller := &MockController{
		db: db,

		index:        0,
		acidRecords:  make([]*model.AcidRecord, 10),
		waterRecords: make([]*model.WaterRecord, 10),
		dieleRecords: make([]*model.DielectronRecord, 10),
		flashRecords: make([]*model.FlashRecord, 10),
	}

	temp := time.Now()
	for i := 0; i < 10; i++ {
		controller.acidRecords[i] = GenerateAcidRecord(temp)
		controller.acidRecords[i].Index = i + 1

		controller.waterRecords[i] = GenerateWaterRecord(temp)
		controller.waterRecords[i].Index = i + 1

		controller.dieleRecords[i] = GenerateDielectricRecord(temp)
		controller.dieleRecords[i].Index = i + 1

		controller.flashRecords[i] = GenerateFlashRecord(temp)

		temp = temp.Add(-time.Hour * 24)
	}

	return controller
}

func (c *MockController) RegisterRouter(r gin.IRouter) {
	r.GET("/frontPage", c.frontPage)
	r.GET("/nextPage", c.nextPage)
	r.GET("/previousPage", c.previousPage)
	r.GET("/print", c.print)

	r.GET("/getId", c.getId)
}

func (c *MockController) frontPage(ctx *gin.Context) {
	c.index = 0

	deviceType, err := getType(ctx)
	if err != nil {
		zlog.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	c.choice(deviceType, ctx)

}

func (c *MockController) nextPage(ctx *gin.Context) {
	if c.index+1 < 10 {
		c.index++
	}

	deviceType, err := getType(ctx)
	if err != nil {
		zlog.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	c.choice(deviceType, ctx)
}

func (c *MockController) previousPage(ctx *gin.Context) {
	if c.index > 0 {
		c.index--
	}
	deviceType, err := getType(ctx)
	if err != nil {
		zlog.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	c.choice(deviceType, ctx)
}

func (c *MockController) print(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}

func (c *MockController) dieleCheckExistAndInsert() error {
	exist, err := model.DielectronRecordIsExist(c.db, c.dieleRecords[c.index].DetectionTime)
	if err != nil {
		return err
	}

	if !exist {
		if err := model.CheckAndInsertDielectronRecord(c.db, c.dieleRecords[c.index]); err != nil {
			return err
		}
	}

	return nil
}

func (c *MockController) acidCheckExistAndInsert() error {
	exist, err := model.AcidRecordIsExist(c.db, c.acidRecords[c.index].DetectionTime)
	if err != nil {
		return err
	}

	if !exist {
		if err := model.CheckAndInsertAcidRecord(c.db, c.acidRecords[c.index]); err != nil {
			return err
		}
	}

	return nil
}

func (c *MockController) flashCheckExistAndInsert() error {
	exist, err := model.FlashRecordIsExist(c.db, c.flashRecords[c.index].DetectionTime)
	if err != nil {
		return err
	}

	if !exist {
		if err := model.CheckAndInsertFlashRecord(c.db, c.flashRecords[c.index]); err != nil {
			return err
		}
	}

	return nil
}

func getType(ctx *gin.Context) (int, error) {
	value := ctx.Query("type")
	if value == "" {
		return -1, errors.New("without device type")
	}

	deviceType, err := strconv.Atoi(value)
	if err != nil {
		return -1, err
	}
	return deviceType, nil
}

func (c *MockController) waterCheckExistAndInsert() error {
	exist, err := model.WaterRecordIsExist(c.db, c.waterRecords[c.index].DetectionTime)
	if err != nil {
		return err
	}

	if !exist {
		if err := model.CheckAndInsertWaterRecord(c.db, c.waterRecords[c.index]); err != nil {
			return err
		}
	}

	return nil
}

func (c *MockController) choice(deviceType int, ctx *gin.Context) {
	switch deviceType {
	case model.DielectronType:
		if err := c.dieleCheckExistAndInsert(); err != nil {
			zlog.Error(err)
			ctx.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "record": c.dieleRecords[c.index]})

	case model.WaterType:
		if err := c.waterCheckExistAndInsert(); err != nil {
			zlog.Error(err)
			ctx.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "record": c.waterRecords[c.index]})

	case model.AcidType:
		if err := c.acidCheckExistAndInsert(); err != nil {
			zlog.Error(err)
			ctx.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "record": c.acidRecords[c.index]})
	
	case model.FlashType:
		if err := c.flashCheckExistAndInsert(); err != nil {
			zlog.Error(err)
			ctx.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "record": c.flashRecords[c.index]})


	}
}

func (c *MockController) getId(ctx *gin.Context) {
	data := []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d}
	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "record": data})
}
