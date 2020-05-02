package engine

import (
	"log"
	"reptile/porsist"
)

//定义一个结构体，里面有调度器和要开worker数量
type ConcurrentEngine struct {
	Scheduler	Scheduler
	WorkerCount int
	ItemChan chan porsist.Profile
}
//定义调度器这个接口
type Scheduler interface {
	ReadyNotifier		//
	Submit(Request)		//提交Request
	WorkerChan()chan Request	//返回Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine)Run(seeds ...Request)  {
	//建立两个通道
	//in := make(chan Request)
	out := make(chan ParseResult)	//这是worker工作完后用来返回的通道
	//e.Scheduler.ConfigureMastWorkerChan(in)
	//启动调度器
	e.Scheduler.Run()
	//根据传入的WorkerCount，来开go
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(),out,e.Scheduler)//主要的工作方法
	}
	//遍历初始传入的种子页面，
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}
	//做一个计数
	for  {
		result := <-out	//从通道中把值取出
		//遍历得到的值打印下
		for _, item := range result.Items {
			//fmt.Println(item)
			e.ItemChan <- item
		}
		//把得到的Requests，填入到通道
		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
		
	}
}

func createWorker(in chan Request,out chan ParseResult,ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)  //把通道加到工作通道
			request := <-in			//从通道中读出Request
			//开始工作
			result, err := worker(request)
			if err != nil {
				//有错误打印下出错的URL,然后跳出此次工作
				log.Printf("出错的URL：%s",request.Url)
				continue
			}
			//工作完后把得到的值放到返回管道
			out <- result
		}
	}()
}