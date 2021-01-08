package engine

// 并发引擎类型
type ConcurrentEngine struct {
	Scheduler        Scheduler // 调度队列
	WorkerCount      int       // 启动进程数量
	ItemChan         chan Item // 数据保存通道
	RequestProcessor Processor //
}

type Processor func(request Request) (ParserResult, error)

// 数据调度Scheduler的接口(duck typing)
// 当前存在simple(单通道)，queued(队列)两个版本
// 详细代码见/scheduler
type Scheduler interface {
	ReadNotifier
	Submit(Request)         // 提交request进通道
	WorkChan() chan Request // 创建一个工作通道
	Run()                   // 初始化scheduler
}

type ReadNotifier interface {
	WorkerReady(chan Request)
}

// 开始执行程序
func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParserResult)

	// 创建调度队列
	e.Scheduler.Run()

	// 根据workerCount配置创建对应数量的工作通道
	for i := 0; i < e.WorkerCount; i++ {
		e.createWorker(e.Scheduler.WorkChan(), out, e.Scheduler)
	}

	// 将请求提交到队列
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

			// 阻塞等待work chan的数据
			request := <-in

			// 对请求进行处理
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
