package scheduler

import "go-craler.com/engine"

/**
  队列的调度器
*/
type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request // 并发请求的的通道全部放入的总workerChan

}

func (q *QueuedScheduler) Submit(r engine.Request) {
	q.requestChan <- r
}

/**
一旦有准备好的work就送入workerChan
*/
func (q *QueuedScheduler) WorkerReady(w chan engine.Request) {
	q.workerChan <- w
}

func (q *QueuedScheduler) ConfigureMasterWorkerChan(chan engine.Request) {
	panic("implement me")
}

/**
总控
*/
func (s *QueuedScheduler) Run() {
	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)

	go func() {
		var requestQueue []engine.Request
		var workerQueue []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQueue) > 0 && len(workerQueue) > 0 {
				activeRequest = requestQueue[0]
				activeWorker = workerQueue[0]
			}
			select {
			case r := <-s.requestChan:
				requestQueue = append(requestQueue, r)
			case w := <-s.workerChan:
				workerQueue = append(workerQueue, w)
			case activeWorker <- activeRequest:
				//将当前激活的请求送给当前激活的worker
				requestQueue = requestQueue[1:]
				workerQueue = workerQueue[1:]

			}
		}
	}()
}
