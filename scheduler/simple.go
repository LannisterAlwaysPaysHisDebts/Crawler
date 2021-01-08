package scheduler

import "Crawler/engine"

// 简单调度器 附带了request通道
type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {

}

// request 写入通道
func (s *SimpleScheduler) Submit(request engine.Request) {
	go func() {
		s.workerChan <- request
	}()
}

func (s *SimpleScheduler) WorkChan() chan engine.Request {
	return s.workerChan
}

// 初始化
func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}
