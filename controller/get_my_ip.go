package controller

import (
	"github.com/Anonymouscn/go-partner/restful"
	restfulmodel "github.com/Anonymouscn/go-partner/restful/model"
	"github.com/Anonymouscn/ip-server/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// GetMyIP 获取我的 IP 地址
func GetMyIP(ctx *gin.Context) {
	ip, isPrivate := service.GetMyIP(ctx)
	ctx.JSON(http.StatusOK, restfulmodel.SuccessWithData(
		restful.Data{
			"ip":         ip,
			"is_private": isPrivate,
			"time":       time.Now().Unix(),
		},
	))
}
