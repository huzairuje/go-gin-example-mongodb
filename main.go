package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/go-gin-example-mongodb/config"
	"github.com/go-gin-example-mongodb/database"
	loanDomain "github.com/go-gin-example-mongodb/loan"
	"github.com/go-gin-example-mongodb/response"
	"github.com/go-gin-example-mongodb/utils"
)

func main() {
	//load config from config.yml
	loadCfg := config.LoadConfig()
	//initiate database connection
	db := database.ConnectDB(loadCfg)
	//initiate gin instance
	router := gin.Default()
	//custom response for not matching routes
	router.NoRoute(func(c *gin.Context) {
		response.NotFound(c, utils.NotMatchingAnyRoute, utils.NotFound)
	})
	//initiate routes by domain (this domain just loan domain)
	loanDomain.Routes(router, db)
	//start web (gin) service
	serverPort := loadCfg.ServerPort
	if serverPort != "" {
		logrus.Info(utils.ServerPortIsSet, serverPort)
		serverPort = loadCfg.ServerPort
	} else {
		logrus.Errorf(utils.ServerPortIsNotSet)
		serverPort = utils.DefaultServerPort
	}
	err := router.Run(serverPort)
	if err != nil {
		logrus.Info(err)
		panic(err)
	}

}
