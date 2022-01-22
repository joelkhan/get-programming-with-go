package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
 * 如果seed固定，那么每次程序重启后，重新生成de随机数会重复上一次的随机数
 * 为了尽量随机性，那么我们可以每次使用不同的seed来启动程序，就可以保证每次启动都产生新的随机数，
 * 聪明的你肯定想到了使用时间戳
 */
func main() {
	var count = 10

	for count > 0 {
		fmt.Println(count)
		time.Sleep(time.Second)
		rand.Seed(time.Now().UnixNano())
		randNum := rand.Intn(100)
		fmt.Printf("your rand number: %v\n", randNum)
		if randNum == 0 {
			break
		}
		count--
	}

	if count == 0 {
		fmt.Println("Liftoff!")
	} else {
		fmt.Println("Launch failed!")
	}
}
