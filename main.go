package main

import (
	"crawler/engine"
	"crawler/zhenai/parser"
	"crawler/scheduler"
)

func main()  {
	engine.ConcurrentEngine{
		Scheduler: &scheduler.SimpleScheduler{},
	}.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}