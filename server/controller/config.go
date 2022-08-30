package controller

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/silverswords/dielectric/server/model"
	"github.com/silverswords/dielectric/server/zlog"
)

// Controller external service interface
type ConfigController struct {
	db *sql.DB
}

// New create an external service interface
func NewconfigController(db *sql.DB) *ConfigController {
	return &ConfigController{
		db: db,
	}
}

// RegisterRouter register router.
func (c *ConfigController) RegisterRouter(r gin.IRouter) {
	r.GET("/get", c.getConfig)
}

func (c *ConfigController) getConfig(ctx *gin.Context) {
	config, err := model.GetConfig(c.db)
	if err != nil {
		zlog.Error(err)
		ctx.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "config": config})
}
