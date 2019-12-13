package scheduler

import "go-craler.com/engine"

/**
简单的调度器
*/
type SimpleScheduler struct {
	WorkChan chan engine.Request
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.WorkChan
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {
	panic("implement me")
}

func (s *SimpleScheduler) Run() {
	s.WorkChan = make(chan engine.Request)
}

/**
  提交数据给workChan
*/
func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() {
		s.WorkChan <- r
	}()
}
