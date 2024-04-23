package service

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"runtime"
	"stuLook-service/global"
	"stuLook-service/model"
	"stuLook-service/utils"
	"sync"
	"time"
)

var wg sync.WaitGroup

func ReceiveMqtt() {
	wg.Add(1)
	task := &model.Task{
		Handler: func(v ...interface{}) {
			wg.Done()
			MqttClientDTU(global.MqTest, "rt-base-West")
		},
		Params: []interface{}{global.MqTest, "rt-base-West"},
	}
	global.Pool.Put(task)
	wg.Wait()
	//// 安全关闭任务池（保证已加入池中的任务被消费完）
	//pool.Close()
	//// 如果任务池已经关闭, Put() 方法会返回 ErrPoolAlreadyClosed 错误
	//err = pool.Put(&mortar.Task{
	//	Handler: func(v ...interface{}) {},
	//})
	//if err != nil {
	//	fmt.Println(err) // print: pool already closed
	//}
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
