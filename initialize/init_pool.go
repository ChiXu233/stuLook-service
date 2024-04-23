package initialize

import (
	"stuLook-service/global"
	"stuLook-service/model"
)

func InitPool() {
	var err error
	if global.Pool == nil {
		global.Pool, err = model.NewPool(10)
		if err != nil {
			panic(err)
		}
	}
}
