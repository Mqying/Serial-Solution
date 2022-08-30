package controller

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/silverswords/dielectric/server/model"
	"github.com/silverswords/dielectric/server/zlog"
)

// Controller external service interface
type RecordController struct {
	db *sql.DB
}

// New create an external service interface
func NewrecordController(db *sql.DB) *RecordController {
	return &RecordController{
		db: db,
	}
}

//RegisterRouter register router and from now on, every API would check if valid on current AdminID.
// Should init the permission to the API.
func (c *RecordController) RegisterRouter(r gin.IRouter) {
	err := model.CreateRecordTable(c.db)
	if err != nil {
		zlog.Error(err)
	}

	r.GET("/get", c.getAllRecords)
}

func (c *RecordController) getAllRecords(ctx *gin.Context) {
	value := ctx.Query("type")
	if value == "" {
		zlog.Error(errors.New("[DEVICE_TYPE] without device type"))
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	deviceType, err := strconv.Atoi(value)
	if err != nil {
		zlog.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	result, err := model.GetAllRecords(c.db, deviceType)
	if err != nil {
		zlog.Error(err)
		ctx.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway})
		return
	}
	
	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "records": result})
}