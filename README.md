# Crawler

# todo:
## 爬虫
1. 爬取抽象, css选择器/xpath分析数据；
2. 对抗反爬技术 / 遵循对方的robots协议；
3. 模拟登陆，爬取动态网页； cookie 

## 查询
1. 优化es查询质量，支持中文；
2. 优化查询体验，前端界面友好；
3. 爬取照片
4. 大数据 AI 分析

## 部署
1. 脚本部署
2. Docker + k8s  
3. 集成服务发现框架 consul
4. Logstash 汇总与日志分析

# NOTE
## 流程
整个系统分为3个模块：
1. persist(Saver): 数据保存模块，这里用的es
2. scheduler: 任务调度队列；
3. worker: 分为
    1. fetcher: 网页数据抓取
    2. parser: 网页数据解析

流程是：
1. 初始化persist(engine外，作为engine的参数传入一个chan)；
2. 初始化scheduler
3. 根据参数创建若干个scheduler的work chan
4. 执行scheduler.Submit，传入需要采集的request




