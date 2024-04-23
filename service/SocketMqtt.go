package service

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"runtime"
	"stuLook-service/global"
	"stuLook-service/model"
	"stuLook-service/utils"
	"time"
)

func ReceiveMqtt() {
	global.Pool.Put(&model.Task{
		Handler: func(v ...interface{}) {
			MqttClientDTU(global.MqTest, "rt-base-West")
		},
		Params: []interface{}{global.MqTest, "rt-base-West"},
	})
}

func MqttClientDTU(mqttClient mqtt.Client, topic string) {
	mqttClient.Subscribe(topic, 0, func(client mqtt.Client, message mqtt.Message) {
		fmt.Println(message.Topic(), message.Payload())
		utils.Try(func() {
			fmt.Println("现有goroutine数量", runtime.NumGoroutine())
			//模拟接收设备并处理
			ParseDTU(message.Payload(), len(message.Payload()))
		})
	})
}

// ParseDTU 模拟解析函数
func ParseDTU(payload []byte, n int) {
	fmt.Println("解析函数收到数据")
	time.Sleep(time.Nanosecond * 200)
	fmt.Println("解析函数结束")
	return
}
