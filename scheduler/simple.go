//抢夺并发式调度器，，这个调度器维护了一个通道，

package scheduler

import "reptile/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

//这是一个创建通道的方法，，抢夺式：只有一个通道，所以直接把Run创建好的通道给每个go就行
func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan		//把创建Run创建好的通道交给每一个GO
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {
}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)		//创建一个Go
}
func (s *SimpleScheduler)Submit (r engine.Request)  {
	go func() {
		s.workerChan <- r		//把任务写到workerChan中
	}()
}

