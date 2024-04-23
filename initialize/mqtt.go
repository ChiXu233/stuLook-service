package initialize

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"time"
)

// MqttInit 使用Mqtt接收瑞通新服务器的CEMS消息
func MqttInit(id string, mqttClient *mqtt.Client) {
	//配置
	address := "tcp://119.45.136.206:1883"
	//设置地址用户名和密码
	clientOptions := mqtt.NewClientOptions().AddBroker(address).SetUsername("stuLook").SetPassword("test111")
	//设置客户端id
	clientOptions.SetClientID(id)
	//掉线后不清除session
	clientOptions.SetCleanSession(false)
	//设置自动重连
	clientOptions.SetAutoReconnect(true)
	//设置处理函数
	clientOptions.OnConnect = func(client mqtt.Client) {
		log.Println("mqttClient," + id + "与服务端建立连接成功！")
	}
	clientOptions.OnConnectionLost = func(client mqtt.Client, err error) {
		log.Printf("mqttClient,"+id+"与服务端断开连接: %v", err)
	}
	*mqttClient = mqtt.NewClient(clientOptions) //客户端建立
	//客户端连接判断
	if token := (*mqttClient).Connect(); token.WaitTimeout(time.Duration(60)*time.Second) && token.Wait() && token.Error() != nil {
		log.Println(token.Error(), "mqtt")
	}
}
