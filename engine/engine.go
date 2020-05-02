package engine

import (
	"log"
	"reptile/fetcher"
)

type SimpleEngine struct {}
//主引擎，主要用来发布任务，传入一个或多个Request
func (e SimpleEngine)Run(seeds ...Request)  {
	var requests []Request	//定义一个Request切片，用来分发任务

	//遍历seeds，把他放入队列中（因为seeds可以传入多个）
	for _,r:=range seeds {
		requests = append(requests,r)
	}

	//分发任务，直到requests队列为零
	var count = 0
	for len(requests) >0{
		//使用第0个requests，然后截取切片1到最后在重新赋值
		count++
		r:=requests[0]
		requests = requests[1:]

		parseResult, err := worker(r)
		if err != nil {
			continue
		}
		//把获取到的request在添加到requests队列中
		requests=append(requests, parseResult.Requests...)
		//这里是打印下获取到的数据
		for _,item:=range parseResult.Items  {
			log.Printf("Got item %s",item)
		}
	}
}

func worker(r Request) (ParseResult, error) {
	//fmt.Println("Url:"+r.Url)
	//调用Fetch来获得一个网页的数据
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("fetching url: %s, Fetcher error: %v",r.Url,err)
		return ParseResult{}, err
	}
	//使用解析器来解析返回的网页内容，获得一个ParseResult
	return r.ParserFunc(body),nil
}