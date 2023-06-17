package controller

import (
	"net/http"

	"cn.a2490/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	*BaseController
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{&BaseController{}, userService}
}

func (controller *UserController) Router(r *gin.Engine) {
	userRouter := r.Group("/user")
	userRouter.POST("/getToken", controller.GetToken)

	userRouter.Use(controller.Auth())
	{
		userRouter.POST("/createUser", controller.CreateUser)
	}
}

// GetToken
// @Tags 获取token
// @Summary 获取token
// @Produce json
// @param phone formData string true "手机号码"
// @Success 200 {object} common.Resp
// @Router /user/getToken [post]
func (controller *UserController) GetToken(c *gin.Context) {
	phone := c.PostForm("phone")
	if phone == "" {
		c.JSON(http.StatusBadRequest, controller.RespFailMsg("phone is request, but it is empty "))
		return
	}
	token, err := controller.userService.GetToken(phone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.RespErrorMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, controller.RespSuccess(token))
}

// CreateUser
// @Tags 人员管理
// @Summary 创建人员
// @Produce json
// @param raffleKey header string true "token"
// @param phone formData string true "手机号码"
// @param name formData string true "姓名"
// @Success 200 {object} common.Resp
// @Router /user/createUser [post]
func (controller *UserController) CreateUser(c *gin.Context) {
	phone := c.PostForm("phone")
	name := c.PostForm("name")
	if phone == "" {
		c.JSON(http.StatusBadRequest, controller.RespFailMsg("phone is request, but it is empty "))
		return
	}
	if name == "" {
		c.JSON(http.StatusBadRequest, controller.RespFailMsg("phone is request, but it is empty "))
		return
	}
	err := controller.userService.CreateUser(phone, name)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, controller.RespErrorMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, controller.RespS())
}
