package main

import (
	"stuLook-service/initialize"
	"stuLook-service/router"
	"stuLook-service/service"
)

func init() {
	initialize.Init()
}
func main() {
	service.ReceiveMqtt()
	engine := router.GetEngine()
	if err := engine.Run(":8081"); err != nil {
		panic(err)
	}

}
