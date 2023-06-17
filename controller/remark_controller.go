package controller

import (
	"net/http"

	"cn.a2490/service"
	"github.com/gin-gonic/gin"
)

type RemarkController struct {
	*BaseController
	remarkService *service.RemarkService
}

func NewRemarkController(remarkService *service.RemarkService) *RemarkController {
	return &RemarkController{&BaseController{}, remarkService}
}

func (controller *RemarkController) Router(r *gin.Engine) {
	remarkRouter := r.Group("/remark")

	remarkRouter.Use(controller.Auth())
	{
		remarkRouter.POST("/list", controller.List)
	}
}

// List
// @Tags 说明管理
// @Summary 说明
// @Produce json
// @param raffleKey header string true "token"
// @Success 200 {object} common.Resp
// @Router /remark/list [post]
func (controller *RemarkController) List(c *gin.Context) {
	remarks := controller.remarkService.List()
	c.JSON(http.StatusOK, controller.RespSuccess(remarks))
}
