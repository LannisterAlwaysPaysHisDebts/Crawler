// 入口文件
package main

import (
	"Crawler/engine"
	"Crawler/persist"
	"Crawler/scheduler"
	"Crawler/source/zhenAi"
)

func main() {
	// 获取一个es数据保存的channel
	itemChan, err := persist.ItemSaver(zhenAi.Index)
	if err != nil {
		panic(err)
	}

	// 初始化引擎类型： SimpleEngine简单引擎， ConcurrentEngine并发引擎
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
		ItemChan:    itemChan,
	}

	e.Run(zhenAi.IndexRequest())
}
