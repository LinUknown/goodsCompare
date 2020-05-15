package crawler

import (
	"imooc.com/crawler/engine"
	jdParser "imooc.com/crawler/jd/parser"
	tmParser "imooc.com/crawler/tmall/parser"

	"imooc.com/model"
	"net/url"
)

//爬虫包的作用： 根据关键字， 爬取商品信息， 以model的形式返回

func Search(keyWord string)[]model.Goods{

	// q= 需要进行 gbk url加密
	data := url.QueryEscape(keyWord)

	tmallRequest := engine.Request{
		URL: "https://list.tmall.com/search_product.htm?q=" + data,// %D3%EA%C9%A1",
		ParserFunc: tmParser.ParseGoodList,
	}
	tmResult := engine.Run(tmallRequest)

	jdRequest := engine.Request{
		URL: "https://search.jd.com/Search?keyword=" + data + "&enc=utf-8",// %D3%EA%C9%A1",
		ParserFunc: jdParser.ParseGoodList,
	}
	jdResult := engine.Run(jdRequest)

	totResult := make([]model.Goods,0)
	for _,item := range tmResult{
		totResult = append(totResult, item)
	}

	for _,item := range jdResult{
		totResult = append(totResult, item)
	}
	return totResult
	//return jdResult
}

