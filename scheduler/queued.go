package scheduler

import "Crawler/engine"

// 调度队列
type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (q *QueuedScheduler) WorkerReady(r chan engine.Request) {
	q.workerChan <- r
}

// 将需要采集的request写入通道
func (q *QueuedScheduler) Submit(r engine.Request) {
	q.requestChan <- r
}

func (q *QueuedScheduler) WorkChan() chan engine.Request {
	return make(chan engine.Request)
}

// 初始化
func (q *QueuedScheduler) Run() {
	q.workerChan = make(chan chan engine.Request)
	q.requestChan = make(chan engine.Request)

	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request

		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}

			select {
			case w := <-q.workerChan:
				workerQ = append(workerQ, w)
			case r := <-q.requestChan:
				requestQ = append(requestQ, r)
			case activeWorker <- activeRequest:
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()

}
