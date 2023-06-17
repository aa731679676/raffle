package controller

import (
	"net/http"

	"cn.a2490/auth"
	"cn.a2490/common"
	"cn.a2490/config"
	"github.com/gin-gonic/gin"
)

type BaseController struct {
}

func (base *BaseController) Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get(config.Config.Token.TokenKey)
		if token == "" {
			c.Abort()
			c.JSON(http.StatusUnauthorized, base.RespF())
			return
		}
		loginId := auth.GetLoginId(token)
		if loginId == "" {
			c.Abort()
			c.JSON(http.StatusUnauthorized, base.RespF())
			return
		}
		c.Set(config.Config.Token.TokenKey, loginId)
		c.Next()
	}
}

func (*BaseController) RespS() common.Resp {
	return common.Resp{
		Code: http.StatusOK,
	}
}

func (c *BaseController) RespSuccess(data interface{}) common.Resp {
	resp := c.RespS()
	resp.Data = data
	return resp
}

func (c *BaseController) RespSuccessMsg(message string) common.Resp {
	resp := c.RespS()
	resp.Message = message
	return resp
}

func (*BaseController) RespF() common.Resp {
	return common.Resp{
		Code: http.StatusBadRequest,
	}
}

func (c *BaseController) RespFail(data interface{}) common.Resp {
	resp := c.RespF()
	resp.Data = data
	return resp
}

func (c *BaseController) RespFailMsg(message string) common.Resp {
	resp := c.RespF()
	resp.Message = message
	return resp
}

func (*BaseController) RespE() common.Resp {
	return common.Resp{
		Code: http.StatusInternalServerError,
	}
}

func (c *BaseController) RespError(data interface{}) common.Resp {
	resp := c.RespE()
	resp.Data = data
	return resp
}

func (c *BaseController) RespErrorMsg(message string) common.Resp {
	resp := c.RespE()
	resp.Message = message
	return resp
}
