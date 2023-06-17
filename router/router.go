package router

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"cn.a2490/common"
	"cn.a2490/config"
	"cn.a2490/controller"
	"cn.a2490/dao"
	"cn.a2490/docs"
	"cn.a2490/service"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitServer(funcs ...func()) {
	port := config.Config.Port
	log.Printf("start Server: %s\n", port)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: Router(),
	}
	serverStartAndShutDown(srv, funcs)
}

func Router() *gin.Engine {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Use(recoverAdvice)

	userDao := dao.NewUserDao()
	userService := service.NewUserService(userDao)
	userController := controller.NewUserController(userService)
	userController.Router(r)

	recordDao := dao.NewRecordDao()
	prizeDao := dao.NewPrizeDao()
	recordService := service.NewRecordService(recordDao, prizeDao)
	recordController := controller.NewRecordController(recordService)
	recordController.Router(r)

	remarkDao := dao.NewRemarkDao()
	remarkService := service.NewRemarkService(remarkDao)
	remarkController := controller.NewRemarkController(remarkService)
	remarkController.Router(r)
	return r
}

// recoverAdvice 全局异常捕获
func recoverAdvice(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("panic: %v\n", r)
			c.JSON(http.StatusInternalServerError,
				common.Resp{
					Code:    http.StatusInternalServerError,
					Message: errorToString(r),
				})
			c.Abort()
		}
	}()
	c.Next()
}

func errorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}

func serverStartAndShutDown(srv *http.Server, funcs []func()) {
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server start error: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	log.Println("shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("server Shutdown Error: %s\n", err)
	}

	if len(funcs) > 0 {
		for _, funcItem := range funcs {
			funcItem()
		}
	}

	log.Println("server exiting ...")
}
