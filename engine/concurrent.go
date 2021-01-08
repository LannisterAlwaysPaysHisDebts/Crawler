package engine

// 并发引擎类型
type ConcurrentEngine struct {
	Scheduler        Scheduler // 调度队列
	WorkerCount      int       // 启动进程数量
	ItemChan         chan Item // 数据保存通道
	RequestProcessor Processor //
}

type Processor func(request Request) (ParserResult, error)

// 调度队列Scheduler的duck typing
type Scheduler interface {
	ReadNotifier
	Submit(Request)
	WorkChan() chan Request
	Run()
}

type ReadNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParserResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		e.createWorker(e.Scheduler.WorkChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out

		// 保存数据
		for _, item := range result.Items {
			go func() { e.ItemChan <- item }()
		}

		// 采集其他资源
		for _, request := range result.Requests {
			if !isDuplicate(request.Url) {
				e.Scheduler.Submit(request)
			}
		}
	}
}

func (e *ConcurrentEngine) createWorker(
	in chan Request,
	out chan ParserResult,
	ready ReadNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)
			request := <-in
			result, err := e.RequestProcessor(request)
			if err != nil {
				continue
			}

			out <- result
		}
	}()
}

var visitedUrls = make(map[string]bool)

// url 去重，防止重复采集 todo：待优化
func isDuplicate(url string) bool {
	if visitedUrls[url] {
		return true
	}

	visitedUrls[url] = true
	return false
}
