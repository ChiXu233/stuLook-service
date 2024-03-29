package main

import (
	"stuLook-service/initialize"
	"stuLook-service/router"
)

func init() {
	initialize.Init()
}
func main() {
	engine := router.GetEngine()
	if err := engine.Run(":8081"); err != nil {
		panic(err)
	}

}
