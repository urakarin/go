package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Request struct {
	name	string
	number	int
}

func (r *Request) Execute() {
	fmt.Printf("executes [ Request from %s No.%d ]\n", r.name, r.number)
}


type Worker struct {
	id    int
	request_queue chan *Request
}

func WorkerThread(wk *Worker, seed int64) {
	rd := rand.New(rand.NewSource(seed))
	for {
		rq := <-wk.request_queue
		fmt.Printf("Worker-%d ", wk.id)
		rq.Execute()
		time.Sleep(time.Duration(rd.Intn(1000)) * time.Millisecond)
		//fmt.Printf("Worker-%d <<<\n", wk.id)
	}
}


type ClientThread struct {
	name   string
	request_queue chan *Request
}

func (c ClientThread) Start(seed int64) {
	rd := rand.New(rand.NewSource(seed))
	i := 0
	for {
		rq := &Request{c.name, i}
		c.request_queue <-rq
		i++
		//time.Sleep(1 * time.Second)
		time.Sleep(time.Duration(rd.Intn(1000)) * time.Millisecond)
	}
}


func main() {
	request_queue := make(chan *Request, 100)
	for i := 0; i < 5; i++ {
		go WorkerThread(&Worker{i, request_queue}, 000 + int64(i))
	}

	go ClientThread{"Alice", request_queue}.Start(111)
	go ClientThread{"Bobby", request_queue}.Start(222)
	go ClientThread{"Chris", request_queue}.Start(333)

	select {}
}

