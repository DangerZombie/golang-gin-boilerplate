//	Gin Boilerplate API:
//	 version: 1.0
//	 title: Gin Framework Boilerplate
//	Schemes: http, https
//	BasePath: /api/v1
//	  Consumes:
//	  - application/json
//	Produces:
//	- application/json
//	- text/html; charset=utf-8
//	securityDefinitions:
//	 Bearer:
//	  type: apiKey
//	  name: Authorization
//	  in: header
//
// swagger:meta
package main

import (
	"strings"

	"github.com/DangerZombie/golang-gin-boilerplate/initialization"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	viper.SetConfigFile("yaml")
	profile := "dev"

	var configFileName []string
	configFileName = append(configFileName, "config-")
	configFileName = append(configFileName, profile)
	viper.SetConfigName(strings.Join(configFileName, ""))
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	viper.AutomaticEnv()

	var logger *zap.Logger
	logOutput := viper.GetString("server.log-output")
	if logOutput == "file" {
		filename := viper.GetString("server.output-file-path")
		logger, _ = initialization.NewZapLogger(filename)
	} else {
		logger, _ = initialization.NewZapLogger("")
	}

	// init connection to DB
	db, err := initialization.DbInit()
	if err != nil {
		panic(err.Error())
	}

	logger.Info("connection db success")

	initialization.ServerInit(logger, db)
}
