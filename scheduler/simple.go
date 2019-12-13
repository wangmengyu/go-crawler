package scheduler

import "go-craler.com/engine"

/**
简单的调度器
*/
type SimpleScheduler struct {
	WorkChan chan engine.Request
}

/**
提交数据给workChan
*/
func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() {
		s.WorkChan <- r
	}()
}

/**
设置workChan
*/
func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.WorkChan = c

}
