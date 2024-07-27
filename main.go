package main

import (
	"flag"
	"order-service/shared/config"
	"order-service/shared/db"
	"order-service/shared/env"
	"order-service/shared/log"
	"order-service/shared/middleware"
	"order-service/transport"
)

var conf *config.Main

func init() {
	level := flag.String("env", "local", "a string")
	flag.Parse()

	conf = env.InitEnv(*level)
}

func main() {
	rest := middleware.InitRest()
	database, err := db.InitDB(conf)
	if err != nil {
		return
	}
	zapLog, err := log.InitLog(conf)
	if err != nil {
		return
	}
	transport.Setup(rest.Group("/v1/admin"), database, zapLog)

	rest.Logger.Fatal(rest.Start(":" + conf.Rest.Port))

}
