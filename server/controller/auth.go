package controller

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/silverswords/dielectric/server/serial"
	"github.com/silverswords/dielectric/server/zlog"
)

// Controller external service interface
type AuthController struct {
	db *sql.DB
}

// New create an external service interface
func NewAuthController(db *sql.DB) *AuthController {
	return &AuthController{
		db: db,
	}
}

// RegisterRouter register router. It fatal because there is no service if register failed.
func (c *AuthController) RegisterRouter(r gin.IRouter) {
	if r == nil {
		zlog.Info("[InitRouter]: server is nil")
	}

	// userAuth crud API
	r.GET("/login", c.Login)
}

//login with name and password
func (c *AuthController) Login(ctx *gin.Context) {
	err := serial.Login()

	if err != nil {
		zlog.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}
