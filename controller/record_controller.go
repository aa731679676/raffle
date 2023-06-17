package controller

import (
	"net/http"
	"strconv"

	"cn.a2490/config"
	"cn.a2490/service"
	"github.com/gin-gonic/gin"
)

type RecordController struct {
	*BaseController
	recordService *service.RecordService
}

func NewRecordController(recordService *service.RecordService) *RecordController {
	return &RecordController{&BaseController{}, recordService}
}

func (controller *RecordController) Router(r *gin.Engine) {
	userRouter := r.Group("/raffle")

	userRouter.Use(controller.Auth())
	{
		userRouter.POST("/doDraw", controller.DoDraw)
	}
}

// DoDraw
// @Tags 抽奖管理
// @Summary 抽奖
// @Produce json
// @param raffleKey header string true "token"
// @Success 200 {object} common.Resp
// @Router /raffle/doDraw [post]
func (controller *RecordController) DoDraw(c *gin.Context) {
	userId, _ := c.Get(config.Config.Token.TokenKey)
	id, _ := strconv.ParseUint(userId.(string), 10, 64)
	prize, err := controller.recordService.DoDraw(uint(id))
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, controller.RespErrorMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, controller.RespSuccess(prize))
}
