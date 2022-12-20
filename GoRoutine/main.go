package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 10; i++ {
		fmt.Println(s, "***", i)
	}
}

// http://golang.site/go/article/21-Go-%EB%A3%A8%ED%8B%B4-goroutine
// 동시성과 병렬성의 차이
func main() {
	// 함수를 동기적으로 실행
	say("Sync")

	// 함수를 비동기적으로 실행
	// Go루틴은 Go 런타임이 관리하는 Lightweight 논리적 (혹은 가상적) 쓰레드
	// Go에서 "go" 키워드를 사용하여 함수를 호출하면, 런타임시 새로운 goroutine을 실행한다.
	// goroutine은 비동기적으로(asynchronously) 함수루틴을 실행하므로,
	// 여러 코드를 동시에(Concurrently) 실행하는데 사용된다.

	go say("Async1")
	go say("Async2")
	go say("Async3")

	// 3초 대기
	time.Sleep(time.Second * 3)
}
