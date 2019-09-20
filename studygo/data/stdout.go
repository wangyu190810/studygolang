package data

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
)

func SyncLog2() {
	cmdStr := `
#!/bin/bash
for var in {1..2}
do
     echo "Hello, Welcome ${var} times "
done`
	cmd := exec.Command("bash", "-c", cmdStr)
	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
	}
	var cmdStdoutPipe io.ReadCloser
	var cmdStderrPipe io.ReadCloser
	for {
		cmdStdoutPipe, err = cmd.StdoutPipe()
		cmdStderrPipe, err = cmd.StderrPipe()
		if cmdStdoutPipe != nil || cmdStderrPipe != nil {
			break
		} else {
			fmt.Println(err)
		}

	}

	f, _ := os.OpenFile("test1231.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	//if ferr != nil {
	//	panic(ferr)
	//}
	defer f.Close()
	logger := log.New(f, "", log.LstdFlags)
	logger.Print("start print log:")
	oldFlags := logger.Flags()
	logger.SetFlags(0)
	var wg sync.WaitGroup
	wg.Add(2)
	go OutsyncLog(logger, cmdStdoutPipe, &wg)
	go OutsyncLog(logger, cmdStderrPipe, &wg)
	//wg.Wait()
	err = cmd.Wait()
	//执行完后再打开log输出的格式
	logger.SetFlags(oldFlags)
	logger.Print("log print done")

}

func OutsyncLog(logger *log.Logger, reader io.ReadCloser, ws *sync.WaitGroup) {
	//因为logger的print方法会自动添加一个换行，所以我们需要一个cache暂存不满一行的
	cache := ""
	buf := make([]byte, 1024)
	defer ws.Done()
	for {
		//logger.Print(buf)
		if reader == nil {
			continue
		}
		strNum, err := reader.Read(buf)
		if strNum > 0 {
			outputByte := buf[:strNum]
			//这里的切分是为了将整行的log提取出来，然后将不满整行和下次一同打印
			outputSlice := strings.Split(string(outputByte), "\n")
			logText := strings.Join(outputSlice[:len(outputSlice)-1], "\n")
			logger.Printf("%s%s", cache, logText)
			cache = outputSlice[len(outputSlice)-1]
		}
		if err != nil {
			if err == io.EOF || strings.Contains(err.Error(), "file already closed") {
				err = nil
			}
		}
	}
}

func SyncLog3() {
	cmdStr := `
#!/bin/bash
for var in {1..10}
do	
	sleep 1
     echo "Hello, Welcome ${var} times "
done`
	cmd := exec.Command("bash", "-c", cmdStr)

	f, _ := os.OpenFile("test1231.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer f.Close()
	//if ferr != nil {
	//	panic(ferr)
	//}
	cmd.Stderr = f
	cmd.Stdout = f
	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
	}
	cmd.Wait()
}
