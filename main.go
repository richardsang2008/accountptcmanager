package main

import (
	"os/signal"

	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/richardsang2008/accountptcmanager/controller"
	"github.com/richardsang2008/accountptcmanager/utility"
	"os"

	"time"

	"github.com/richardsang2008/accountptcmanager/services"

	"context"
	"net/http"
)

func setupRouter() *gin.Engine {
	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.POST("/account/add", services.AddAccount)
	router.GET("/account/add", services.AddAccount)
	router.POST("/account/update", services.UpdateAccountBySpecificFields)
	router.GET("/account/request", services.GetAccountBySystemIdAndLevelAndMark)
	router.POST("/account/release", services.ReleaseAccount)
	//end of meet the old one

	router.POST("/account", services.AddAccount)
	return router
}

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {
	utility.MCache.New(5*time.Minute, 10*time.Minute)
	config := utility.LoadConfiguration("appconfig.json")
	utility.MCache.Set("configuration", config, 5*time.Minute)
	//s,found:=utility.MCache.Get("configuration")
	//f,_:=utility.MLog.New(config.LogFile, config.LogLevel)
	utility.MLog.New(config.LogFile, config.LogLevel)
	//defer f.Close()
	controller.Data.New(config.MysqlDatabase.Username, config.MysqlDatabase.Password, config.MysqlDatabase.Host, config.MysqlDatabase.DBName)
	defer controller.Data.Close()
	router := setupRouter()
	address := fmt.Sprintf("%v:%s", config.Host, config.Port)
	srv := &http.Server{
		Addr:    address,
		Handler: router,
	}
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil {
			utility.MLog.Info("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	utility.MLog.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		utility.MLog.Panic("Server Shutdown:", err)
	}
	utility.MLog.Info("Server exiting")
}
