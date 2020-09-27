package main

import (
	"gotmp/internal/conf"
	"log"
)

func main()  {
	conf.InitConfigs()
	log.Printf("App Name: %s, Version: %s", conf.GetString("app.name"), conf.GetString("app.version"))
}