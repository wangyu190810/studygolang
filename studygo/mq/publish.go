package mq

type Publish struct {

}

func (self *Subscribe)Publish(topic string,msg Message){
	for _ , subscribers := range self.subscribers{
		if topic == subscribers.topic{
			subscribers.ch <- msg
		}
	}

}



