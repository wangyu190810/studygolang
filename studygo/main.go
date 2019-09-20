package main

import (
	"studygolang/studygo/zeromq"
)

// "studygolang/studygo/socket"

func twoSum(nums []int, target int) []int {
	var num []int
	for key, row := range nums {
		end := target - row
		// if end >= 0 {
		for key2, row2 := range nums[key:] {
			if row2 == end && key != key2 {
				num = append(num, key)
				num = append(num, key2+key)
				return num
			}
		}
		// }
	}
	return num
}

func main() {
	//base.Hello()
	//base.Run()
	//base.InterRun()
	//base.RunChannel()
	//base.RunReflect()
	//mq.Start()
	// socket.RunServer()
	// nums := []int{-1, -2, -3, -4, -5}
	// num_sum := twoSum(nums, -8)
	// fmt.Println("sum", num_sum)
	// leetcode.Run()
	zeromq.RunZero()
	// data.SyncLog3()
}
