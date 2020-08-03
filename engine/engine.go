package engine

import (
	"fmt"
	"sync"
	"time"

	"webterren.com/demo/downloader"
	"webterren.com/demo/log"
)

type Scheduler interface {
	Submit(request *Request)   //将新的request 放入 请求对象的 管道
	WorkChan() chan *Request   // 给work 创建一个专属的 channel
	WorkReady(w chan *Request) // 将worker 产生的request 写入 worker专属的channel
	Run()
}

type Engine struct {
	Scheduler   Scheduler
	WorkerCount int
}

func (G_engine *Engine) Run(requests ...Request) {

	out := make(chan *ParseResult)

	G_engine.Scheduler.Run()

	go func() {
		var wg sync.WaitGroup
		for i := 0; i < G_engine.WorkerCount; i++ {

			creatWorker(G_engine.Scheduler.WorkChan(), out, G_engine.Scheduler, &wg)
		}
		wg.Wait()
		close(out)
	}()

	for _, request := range requests {
		G_engine.Scheduler.Submit(&request)
	}

	for {

		result, ok := <-out
		if !ok {
			break
		}
		for _, item := range result.Items {
			fmt.Printf("Got Item : %v \n", item)
		}

		for _, resq := range result.Requests {

			G_engine.Scheduler.Submit(&resq)
		}

	}
	fmt.Printf(">>>>>最后的到的具体信息的数量:%v<<<<<", log.Count)
	// for {

	// 	result := <-out

	// 	for _, item := range result.Items {
	// 		fmt.Printf("Got Item : %v \n", item)
	// 	}

	// 	for _, resq := range result.Requests {

	// 		G_engine.Scheduler.Submit(&resq)
	// 	}
	// }
}

/*
	in 是worker的专属channel 存放 Request
	out 作为输出channel 存储解析的结果
	scheduler 作为调度中心 会将新的request对象提交
*/

func creatWorker(in chan *Request, out chan *ParseResult, schduler Scheduler, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {

		for {
			var timer = time.NewTicker(5 * time.Second)
			schduler.WorkReady(in)
			select {
			case request := <-in:

				fmt.Print(">>>><<<<<")
				node, err := downloader.DownLoad(request.Url)

				if nil != err {

					fmt.Printf("fetch url %s error,error : %v", request.Url, err)
					continue
				}
				parseResult := request.Parsefunc(node)

				out <- &parseResult

			case <-timer.C:
				wg.Done()
				return
			}

		}

	}()

}
