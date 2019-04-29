package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type MainController struct{}

func (ctrl MainController) Main(ctx *gin.Context) {
	ctx.Redirect(http.StatusMovedPermanently, "/web")
}
