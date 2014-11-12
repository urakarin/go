package main

import (
	"fmt"
	"math/rand"
	"time"
)

type FutureData struct {
	content chan string
}

type RealData struct {
	count	int
	chara	string
}

func (rd *RealData) heavyWork(seed int64) string {
	rnd := rand.New(rand.NewSource(seed))
	msg := ""
	fmt.Printf("        making RealData(%d, %s) BEGIN\n", rd.count, rd.chara);
	for i := 0;  i < rd.count; i++ {
		msg += rd.chara
		time.Sleep(time.Duration(rnd.Intn(1000)) * time.Millisecond)
	}
	fmt.Printf("        making RealData(%d, %s) END\n", rd.count, rd.chara);
	return msg
}


func FutureThread(ft *FutureData, count int, chara string) {
	//fmt.Printf("  futureThread (%d, %s) BEGIN\n", count, chara)
	rd := &RealData{count, chara}
	msg := rd.heavyWork(123)
	ft.content <- msg
}


func request(count int, chara string) *FutureData {
	fmt.Printf("    request(%d, %s) BEGIN\n", count, chara)
	// (1)
	future := &FutureData{make(chan string, 1)}

	// (2)
	go FutureThread(future, count, chara)
	fmt.Printf("    request(%d, %s) END\n", count, chara)

	// (3)
	return future;
}


func main() {
	fmt.Printf("main BEGIN\n")
    data1 := request(10, "A");
    data2 := request(20, "B");
    data3 := request(30, "C");

    fmt.Printf("main otherJob BEGIN\n")
	time.Sleep(2 * time.Second)
    fmt.Printf("main otherJob END\n")

    fmt.Printf("data1 = %s\n", <-data1.content)
    fmt.Printf("data2 = %s\n", <-data2.content)
    fmt.Printf("data3 = %s\n", <-data3.content)
	fmt.Printf("main END\n")
}

