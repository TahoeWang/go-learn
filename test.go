package main

import (
	"fmt"
	"runtime"
	"time"
)

func sum(values [] int, resultChan chan int) {
	sum := 0
	for _, value := range values {
		sum += value
	}
	resultChan <- sum // 将计算结果发送到channel中
	resultChan <- 2
}

type Block struct {
	number int
}

type Body struct {
	block *Block
	id int
}

// work_pool
func worker(id int,jobs<-chan int,res chan<- int){
	for job := range jobs{
		fmt.Printf("worker:%d job:%d\n",id,job)
		res <- job*2
		time.Sleep(time.Millisecond*500)
		fmt.Printf("worker:%d job:%d\n",id,job)
	}

}
func main() {
	jobs :=make(chan int,100)
	res := make(chan int,100)

	// 开启3个goroutine
	for j:=0;j<3;j++{
		go worker(j,jobs,res)
	}
	// channel中无值，range时，会阻塞下边语句执行，阻塞到channel关闭，再接着执行下边的语句
	// 虽然目前jobs是空的，但worker不会range完退出，而是会一直等待jobs关闭。

	fmt.Printf("goroutine 数量 %d\n", runtime.NumGoroutine())
	//runtime.NumGoroutine()

	// 发送5个任务
	for i:=0;i<5;i++{
		jobs <- i
	}

	time.Sleep(time.Millisecond*2000)
	// 关闭jobs以后，worker接着执行下面的逻辑，结束goroutine
	fmt.Printf("goroutine 数量 %d\n", runtime.NumGoroutine())

	close(jobs)

	time.Sleep(time.Millisecond*2000)
	// 关闭jobs以后，worker接着执行下面的逻辑，结束goroutine
	fmt.Printf("goroutine 数量 %d\n", runtime.NumGoroutine())

	//for i:=0;i<5;i++{
	//	jobs <- i
	//}

	//输出结果
	for i:=0;i<5;i++{
		ret := <-res
		fmt.Println(ret)
	}
	fmt.Printf("goroutine 数量 %d\n", runtime.NumGoroutine())
}
//func main() {
//	values := [] int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
//	resultChan := make(chan int, 1)
//	go sum(values[:len(values)/2], resultChan)
//	go sum(values[len(values)/2:], resultChan)
//
//	fmt.Print(<-resultChan)
//
//	resultChan <- 10
//
//	fmt.Print(<-resultChan)
//
//	//sum1, sum2 := <-resultChan, <-resultChan // 接收结果
//	//fmt.Println("Result:", sum1, sum2, sum1 + sum2)
//	//body := new(Body)
//	//fmt.Print(body.block)
//}

