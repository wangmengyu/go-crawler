1. 启动distributed-server
cd /Users/wangmengyu/work/code/go-crawler/distributed/persist/server/
go run itemsaver.go --port=1234

2. 启动worker（抓取+解析）
cd /Users/wangmengyu/work/code/go-crawler/distributed/worker/server
go run workder.go --port=9000
go run workder.go --port=9001

3. 启动根目录的main方法
cd /Users/wangmengyu/work/code/go-crawler/
go run main.go --itemsaver_host=":1234" --worker_hosts=":9000,:9001"


