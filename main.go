package main

import (
	"webterren.com/demo/engine"
	"webterren.com/demo/pprof"
	"webterren.com/demo/scheduler"
	"webterren.com/demo/spiders/zhenai"
)

func main() {
	var url = "http://www.zhenai.com/zhenghun"
	pprof.Init()
	(&engine.Engine{WorkerCount: 50, Scheduler: &scheduler.Scheduler{}}).Run(engine.Request{
		Url:       url,
		Parsefunc: zhenai.ParseCityList,
	})
}
