package scheduler

import "Crawler/engine"

// 调度队列
// Run方法创建好两个通道request chan 与worker chan， 同时异步调用doWork
// WorkChan方法创建一个request chan，在engine代码中这个chan被称为in,同时engine代码中还有一个chan被称为out
// 具体流程是：
// 0. 初始化Run，初始化doWork
// 1. 根据配置文件调用WorkChan方法创建若干个in给外部
// 2. 外部调用WorkerReady将in全部写入workerChan
// 3. 外部调用Submit方法将request写入requestChan
// 4. doWork将workerChan与requestChan的所有数据分别写入两个队列
// 5. 如果两个队列都不为空，提取出一个request与一个in，将request写入in
// 6. 外部从in拿到对应的request，处理完之后重新调用WorkerReady将in投入WorkChan
// 这样写的优点：1. 根据配置同时创建n个worker,比较灵活；2. 如果写入大量request，有queue可以进行缓冲；
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
	// 创建两个chan， 一个传递请求，一个传递请求通道
	q.workerChan = make(chan chan engine.Request)
	q.requestChan = make(chan engine.Request)

	go q.doWork()
}

//
func (q *QueuedScheduler) doWork() {
	// 请求列表/工作公道 列表
	var requestQ []engine.Request
	var workerQ []chan engine.Request

	for {
		// 当前活跃的请求/工作通道
		// 会将请求写入worker通道里面
		var activeRequest engine.Request
		var activeWorker chan engine.Request

		// 当列表大于0，则写入
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
			// 必须将裁剪queue的逻辑放在这里， 因为select不一定会走入当前case
			// 如果未走进当前case，仅上面active的赋值操作会重复进行
			workerQ = workerQ[1:]
			requestQ = requestQ[1:]
		}
	}
}
