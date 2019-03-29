package socket

import (
	"fmt"
	"net"
	"sync"
)

func client_read(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	line, flag := read(conn)
	if flag == false {
		return
	}
	fmt.Println(line)
}

func clinet_handler(conn net.Conn, wg *sync.WaitGroup) {

	defer wg.Done()
	write([]byte("ping \n"), conn)
	go client_read(conn, wg)
	write([]byte("echo data \n"), conn)
	go client_read(conn, wg)
	write([]byte("quit \n"), conn)
	go client_read(conn, wg)
}

func Error(err error) bool {
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func RunClient() {
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if !Error(err) {
		return
	}
	var wg sync.WaitGroup
	wg.Add(4)
	go clinet_handler(conn, &wg)

	wg.Wait()
}
