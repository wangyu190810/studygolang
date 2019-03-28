package mq

type Publish struct {
}

func (self *Subscribe) Publish(topic string, msg Message) {
	subscribes, err := self.topic[topic]
	if err != false {
		for _, subscribe := range subscribes {
			subscribe.ch <- msg
		}
	}
}
