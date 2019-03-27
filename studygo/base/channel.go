package base

import (
	"fmt"
	"time"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func sum_channel(){
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c
	fmt.Println(x, y, x+y)
}

func channel_for_range(){
	go func() {
		time.Sleep(1 * time.Hour)
	}()
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i = i + 1 {
			c <- i
		}
		close(c)
	}()
	for i := range c {
		fmt.Println(i)
	}
	fmt.Println("Finished")
}
func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
func channel_fibonacci(){
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}

type ChannelData struct {

}

type Msg struct {
	id int
	msg string
}

func (channel_data *ChannelData)send(ch chan Msg, msg2 Msg){
	ch <- msg2
}

func (channel_data *ChannelData)receive(ch chan Msg){
	data := <- ch
	fmt.Println(data)
}

func channel_msg()  {
	msg := Msg{1000, "start"}
	ch := make(chan Msg,1)
	channel_data := ChannelData{}
	channel_data.send(ch,msg)
	channel_data.receive(ch)
}

func RunChannel() {
	sum_channel()
	channel_for_range()
	channel_fibonacci()
	channel_msg()
	for{
		time.Sleep(100*time.Second)
	}
}
