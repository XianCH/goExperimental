package main

import (
	g "github.com/x14n/goExperimental/gin-crud-api/v3/global"
	"github.com/x14n/goExperimental/gin-crud-api/v3/server"
)

func main() {

	g.ZapInit()
	server.SCore()
	// g.GLogger.Info("fuck u")
	// _, err := g.InitMysql()
	// if err != nil {
	// 	g.GLogger.Info(err.Error())
	// }
	//
}
