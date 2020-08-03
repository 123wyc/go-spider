package scheduler

import (
	"webterren.com/demo/engine"
)

type Scheduler struct {
	requestChan chan *engine.Request
	workerChan  chan chan *engine.Request
}

func (Schedu *Scheduler) WorkChan() chan *engine.Request {

	return make(chan *engine.Request)
}

func (Schedu *Scheduler) WorkReady(w chan *engine.Request) {

	Schedu.workerChan <- w
}

func (Schedu *Scheduler) Submit(request *engine.Request) {

	Schedu.requestChan <- request
}

func (scheduler *Scheduler) Run() {

	//初始化scheduler
	scheduler.workerChan = make(chan chan *engine.Request)
	scheduler.requestChan = make(chan *engine.Request)

	go func() {

		var workerQ []chan *engine.Request
		var requestQ []*engine.Request

		for {
			var activeReq *engine.Request
			var activeWorker chan *engine.Request

			if len(requestQ) > 0 && len(workerQ) > 0 {

				activeReq = requestQ[0]
				activeWorker = workerQ[0]
			}

			select {

			case r := <-scheduler.requestChan:
				// 有新请求存入 requestChan

				requestQ = append(requestQ, r)

			case w := <-scheduler.workerChan: //说明生成了新的worker
				workerQ = append(workerQ, w)

			case activeWorker <- activeReq:
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}
		}

	}()

}
