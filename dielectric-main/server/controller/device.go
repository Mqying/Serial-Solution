package controller

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/silverswords/dielectric/server/model"
	sensor "github.com/silverswords/dielectric/server/serial"
	"go.uber.org/zap"

	"github.com/silverswords/dielectric/server/global"
	"github.com/silverswords/dielectric/server/zlog"
)

// Controller external service interface
type DeviceController struct {
	s  *sensor.VoltageSensor
	db *sql.DB
}

// New create an external service interface
func NewDeviceController(db *sql.DB) *DeviceController {
	return &DeviceController{
		db: db,
	}
}

func (c *DeviceController) openSerialPort(deviceType int) error {
	if global.IsDebugMode() {
		return nil
	}

	if c.s == nil {
		v := &sensor.VoltageSensor{}
		if err := v.OpenSerial(deviceType); err != nil {
			zlog.Error(err)
			return err
		}

		c.s = v
	}

	return nil
}

func (c *DeviceController) RegisterLoginRouter(r gin.IRouter) {
	r.GET("/login", c.checkDeviceID)
}

//RegisterRouter register router
func (c *DeviceController) RegisterRouter(r gin.IRouter) {
	r.GET("/frontPage", c.frontPage)
	r.GET("/nextPage", c.nextPage)
	r.GET("/previousPage", c.previousPage)
	r.GET("/print", c.print)
}

func (c *DeviceController) checkDeviceID(ctx *gin.Context) {
	if global.IsDebugMode() {
		ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
		return
	}

	deviceType, err := model.GetDeviceType(c.db)

	if err != nil {
		zlog.Error(err, zap.String("init", "device type failed"))
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	if err = c.openSerialPort(deviceType); err != nil {
		zlog.Error(err, zap.String("init", "open serail failed"))
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}
	
	err = sensor.Login()

	if err != nil {
		zlog.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}

//send command to device and send data received
func (c *DeviceController) frontPage(ctx *gin.Context) {
	var (
		gbkData []byte
	)

	deviceType, err := GetDeviceType(ctx)

	if err != nil {
		zlog.Error(err)
		ctx.JSON(http.StatusOK, gin.H{"status": http.StatusBadRequest, "errorCode":0})
		return
	}

	if global.IsDebugMode() {
		switch (deviceType) {
		case model.DielectronType:
			gbkData, err = global.DielectronMockData(1)
		default:
		}
	} else {
		if err = c.openSerialPort(deviceType); err != nil {
			zlog.Error(err)
			ctx.JSON(http.StatusOK, gin.H{"status": http.StatusBadRequest, "errorCode":100})
			return
		}

		gbkData, err = c.s.SendFrontPage(deviceType)
	}

	if err != nil {
		zlog.Error(err)
		ctx.JSON(http.StatusOK, gin.H{"status": http.StatusBadGateway,"errorCode":1})
		return
	}

	record, err := c.Parse(gbkData, deviceType)
	if err != nil {
		zlog.Error(err)
		ctx.JSON(http.StatusOK, gin.H{"status": http.StatusBadGateway, "errorCode":11})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "record": record})
}

func (c *DeviceController) nextPage(ctx *gin.Context) {
	var (
		gbkData []byte
	)

	deviceType, err := GetDeviceType(ctx)
	if err != nil {
		zlog.Error(err)
		ctx.JSON(http.StatusOK, gin.H{"status": http.StatusBadRequest, "errorCode":0})
		return
	}

	if global.IsDebugMode() {
		switch (deviceType) {
		case model.DielectronType:
			gbkData, err = global.DielectronMockData(1)
		default:
		}
	} else {
		if err = c.openSerialPort(deviceType); err != nil {
			zlog.Error(err)
			ctx.JSON(http.StatusOK, gin.H{"status": http.StatusBadRequest, "errorCode":100})
			return
		}

		gbkData, err = c.s.SendNextPage(deviceType)
	}

	if err != nil {
		zlog.Error(err)
		ctx.JSON(http.StatusOK, gin.H{"status": http.StatusBadGateway, "errorCode":2})
		return
	}

	record, err := c.Parse(gbkData, deviceType)
	if err != nil {
		zlog.Error(err)
		ctx.JSON(http.StatusOK, gin.H{"status": http.StatusBadGateway, "errorCode":10})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "record": record})
}

func (c *DeviceController) previousPage(ctx *gin.Context) {
	var (
		gbkData []byte
	)

	deviceType, err := GetDeviceType(ctx)
	if err != nil {
		zlog.Error(err)
		ctx.JSON(http.StatusOK, gin.H{"status": http.StatusBadRequest, "errorCode":0})
		return
	}

	if global.IsDebugMode() {
		switch (deviceType) {
		case model.DielectronType:
			gbkData, err = global.DielectronMockData(-1)
		default:
		}
	} else {
		if err = c.openSerialPort(deviceType); err != nil {
			zlog.Error(err)
			ctx.JSON(http.StatusOK, gin.H{"status": http.StatusBadRequest, "errorCode":100})
			return
		}

		gbkData, err = c.s.SendPreviousPage(deviceType)
	}

	if err != nil {
		zlog.Error(err)
		ctx.JSON(http.StatusOK, gin.H{"status": http.StatusBadGateway, "errorCode":3})
		return
	}

	record, err := c.Parse(gbkData, deviceType)
	if err != nil {
		zlog.Error(err)
		ctx.JSON(http.StatusOK, gin.H{"status": http.StatusBadGateway, "errorCode":9})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "record": record})
}

//send command to device,then it print
func (c *DeviceController) print(ctx *gin.Context) {
	err := c.s.SendPrint()
	if err != nil {
		zlog.Error(err)
		ctx.JSON(http.StatusOK, gin.H{"status": http.StatusBadGateway, "errorCode":7})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}