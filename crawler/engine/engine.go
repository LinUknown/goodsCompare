package engine

import (
	"imooc.com/crawler/fetcher"
	"imooc.com/model"
	"log"
)

func Run(seeds ...Request)([]model.Goods){

	//这里维持一个队列
	var requestsQueue []Request
	requestsQueue = append(requestsQueue, seeds...)

	var resultList =  make([]model.Goods,0)

	for len(requestsQueue) > 0 {
		//取第一个
		r := requestsQueue[0]
		//只保留没处理的request
		requestsQueue = requestsQueue[1:]

		log.Printf("fetching url:%s\n", r.URL)
		//爬取数据
		body, err := fetcher.Fetch(r.URL)

		//log.Printf("%s",body)
		if err != nil {
			log.Printf("fetch url: %s; err: %v\n", r.URL, err)
			//发生错误继续爬取下一个url
			continue
		}

		//解析爬取到的结果
		result := r.ParserFunc(body)

		//把爬取结果里的request继续加到request队列
		requestsQueue = append(requestsQueue, result.Requests...)

		//打印每个结果里的item
 		for _, item := range result.Items {
			resultList = append(resultList, item)
		}
	}
	return resultList
}