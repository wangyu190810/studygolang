package mq

import (
	"fmt"
	"time"
)

type Message struct {
	id int
	data string
	timestamp time.Time

}

type Subscriber struct {
	ch chan Message
	topic string
}

type Subscribe struct {
	subscribers []Subscriber
}

func (self *Subscribe) recived(ch chan Message ){
	data := <- ch
	fmt.Println(data)
}

func (self *Subscribe) register(subscriber  Subscriber) {
	self.subscribers = append(self.subscribers,subscriber)
}

func (self *Subscribe)distribute(){
	for _,subscriber := range self.subscribers{
		 self.recived(subscriber.ch)
	}
}

func  Start(){
	subscribe := Subscribe{}
	// 注册订阅消息
	subscribe.register(Subscriber{make(chan Message, 10),"hello"})
	subscribe.register(Subscriber{make(chan Message, 10),"hello_world"})

	// 数据信息
	msg := Message{1,"hello",time.Now()}
	// 分发到不同的主题
	subscribe.Publish("hello",msg)
	//subscribe.Publish("hello_world",msg)
	subscribe.Publish("hello_world",Message{2,"hello_world",time.Now()})
	// 订阅结果展示
	subscribe.distribute()
}




