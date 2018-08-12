package scheduler

import "crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	s.workerChan <- r
}

func (s *SimpleScheduler) configureMasterWorkerChan(c chan engine.Request) {
	s.workerChan = c
}