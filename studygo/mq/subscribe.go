package mq

import (
	"fmt"
	"time"
)

// user info 接受者的信息
type User struct {
	id   int
	name string
}

// 发送的消息结构
type Message struct {
	id        int
	data      string
	timestamp time.Time
}

// 订阅者的订阅信息
type Subscriber struct {
	ch    chan Message
	user  User
	topic string
	task  func(msg Message)
}

func defaultCall(msg Message) {
	fmt.Println("default call ", msg)
}

// 订阅的数据操作容器
type Subscribe struct {
	topic map[string][]Subscriber
}

// 注册订阅者的订阅信息，订阅消息主题，用户，订阅者消息chan 大小
func (self *Subscribe) register(subscriber Subscriber) {
	self.topic[subscriber.topic] = append(self.topic[subscriber.topic], subscriber)
}

// 展示订阅者收到的信息
func (self *Subscribe) recived(ch chan Message) Message {
	data := <-ch

	//fmt.Println(data)
	return data
}

func (self *Subscribe) registerAllTopic(user User) {
	for topic, _ := range self.topic {
		subscriber := Subscriber{make(chan Message, 10), user, topic, defaultCall}
		self.topic[topic] = append(self.topic[topic], subscriber)
	}
}

// 订阅者列表，订阅的数据进行展示
func (self *Subscribe) distribute(topic string) {
	subscribes, err := self.topic[topic]
	if err != false {
		for _, subscribe := range subscribes {
			msg := self.recived(subscribe.ch)
			fmt.Println(subscribe.user)
			subscribe.task(msg)
		}
	}
}

func Start() {
	// 订阅者容器初始化
	subscribe := Subscribe{}
	// topic map make 初始化
	subscribe.topic = make(map[string][]Subscriber)
	// 注册订阅消息
	user := User{2, "22too"}
	// 注册订阅主题和订阅的用户
	//subscribe.register(Subscriber{make(chan Message, 10),user,"hello"})
	subscribe.register(Subscriber{make(chan Message, 10), User{3, "33too"}, "hello", defaultCall})
	subscribe.register(Subscriber{make(chan Message, 10), user, "hello_world", defaultCall})
	subscribe.registerAllTopic(User{4, "44too"})
	// 数据信息
	msg := Message{1, "hello", time.Now()}
	// 分发到不同的主题
	subscribe.Publish("hello", msg)
	//subscribe.Publish("hello_world",msg)
	subscribe.Publish("hello_world", Message{2, "hello_world", time.Now()})
	// 订阅结果展示
	subscribe.distribute("hello")
	subscribe.distribute("hello_world")

}
