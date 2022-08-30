package main

import (
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"

	controller "github.com/silverswords/dielectric/server/controller"

	"github.com/silverswords/dielectric/server/middleware"
	"github.com/silverswords/dielectric/server/zlog"
)

const (
	maxOpenConns = 10000
	maxIdleConns = 1000
	maxLifetime  = time.Duration(30) * time.Second

	adminRouterGroup  = "admin"
	deviceRouterGroup = "device"
	recordRouterGroup = "record"
	configRouterGroup = "config"

	serverAddr = "127.0.0.1:9090"
)

func main() {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())
	routerBasicGroup := router.Group("/api/v1")

	db, err := sql.Open("sqlite3", "config/sqlite.db")
	if err != nil {
		zlog.Fatal(err.Error())
	}

	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxLifetime(maxLifetime)
	defer db.Close()

	recordController := controller.NewrecordController(db)
	configController := controller.NewconfigController(db)
	//deviceController := mock.NewMockController(db)
	deviceController := controller.NewDeviceController(db)
	
	recordController.RegisterRouter(routerBasicGroup.Group(recordRouterGroup))
	deviceController.RegisterLoginRouter(routerBasicGroup.Group(adminRouterGroup))
	deviceController.RegisterRouter(routerBasicGroup.Group(deviceRouterGroup))
	configController.RegisterRouter(routerBasicGroup.Group(configRouterGroup))

	zlog.Fatal(router.Run(serverAddr).Error())
}