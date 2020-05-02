//队列并发式调度器，这个调度器维护了两个通道，一个requestChan通道，和一个工作通道的 通道，
package scheduler

import "reptile/engine"


//
type QueuedScheduler struct {
	requestChan chan engine.Request				//这是任务通道
	workerChan	chan chan engine.Request		//这是工作通道的 通道
}

//这是一个创建通道的方法，，队列式：每个go，创建一个通道
func (q *QueuedScheduler) WorkerChan()chan engine.Request {
	return make(chan engine.Request)
}

func (q *QueuedScheduler) Submit(r engine.Request) {
	q.requestChan <- r		//把任务给requestChan通道
}

func (q *QueuedScheduler)WorkerReady(w chan engine.Request)  {
	q.workerChan <- w	//把要工作的工作通道加到
}

func (q *QueuedScheduler)Run()  {
	//把两个通道创建出来
	q.requestChan = make(chan engine.Request)
	q.workerChan = make(chan chan engine.Request)

	//创建一个go，一直在进行队列资源的分发
	go func() {
		//定义两个队列：任务队列，和工作通道队列
		var requestQ  []engine.Request
		var workerQ   []chan engine.Request

		for {
			var activeRequest  engine.Request
			var activeWorker chan engine.Request

			//如果任务队列和工作队列都大于零，则分配任务
			if len(requestQ)>0&&len(workerQ)>0{
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}

			//监听两个两个通道
			select {
			case r:=<-q.requestChan:			//如果requestChan发生了写出的操作
				requestQ = append(requestQ,r)	//则把requestChan加入到requestQ这个队列中
			case w:= <-q.workerChan:			//如果workerChan发生了写出的操作
				workerQ = append(workerQ,w)		//则把workerChan加入到workerQ这个队列中
			case activeWorker <- activeRequest:	//如果任务写进任务通道的操作
				workerQ = workerQ[1:]			//则把该任务和该通道拿掉
				requestQ =requestQ[1:]
			}


		}
	}()
}

