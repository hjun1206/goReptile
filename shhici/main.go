package main

import (
	"reptile/engine"
	"reptile/porsist"
	"reptile/scheduler"
	"reptile/shhici/parser"
)

func main() {
	 e:=engine.ConcurrentEngine{
		 Scheduler:   &scheduler.QueuedScheduler{},
		 WorkerCount: 10,
		 ItemChan:porsist.ItemSaver(),
	 }
	 e.Run(engine.Request{
		 Url:        "https://so.gushiwen.org/authors/",
		 ParserFunc: parser.ParsePoetList,
	 })
}