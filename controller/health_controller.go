package controller

import (
	restfulmodel "github.com/Anonymouscn/go-partner/restful/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Health 健康检查接口
func Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, restfulmodel.Success())
}
