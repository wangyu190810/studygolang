package socket

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
)

func write(data []byte, conn net.Conn) {
	_, err := conn.Write(data)
	if err != nil {
	}
}

func handler(conn net.Conn, ws *sync.WaitGroup) {
	//var read_byte = make([]byte,1024)
	// 退出时，调用
	defer ws.Done()
	for {
		reader := bufio.NewReader(conn)
		line, err := reader.ReadString(byte('\n'))
		if err != nil {
			err := conn.Close()
			if err != nil {
				fmt.Println(err)
			}
			return
		}
		// 读取一行数据。
		inputs := strings.Split(line, " ")
		switch inputs[0] {
		case "ping":
			{
				fmt.Println(inputs[0])
				write([]byte("pong\n"), conn)
			}
		case "echo":
			{
				echoStr := strings.Join(inputs[1:], " ") + "\n"
				write([]byte(echoStr), conn)
			}
		case "quit":
			{
				err := conn.Close()
				if err != nil {
					fmt.Println(err)
				}
				return
			}
		default:
			fmt.Println(inputs)
		}
	}

}

func RunServer() {
	var err error
	var server net.Listener
	server, err = net.Listen("tcp", "0.0.0.0:8999")
	if err != nil {
		fmt.Println(err)
	}
	//for {
	conn, err := server.Accept()
	if err != nil {
		fmt.Println(err)
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go handler(conn, &wg)
	wg.Wait()
	//defer conn.Close()
	//}

}
