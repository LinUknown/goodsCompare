package parser

import (
	"imooc.com/crawler/engine"
	"imooc.com/model"
	"math"
	"regexp"
	"strconv"
)

var (
	// <em>￥</em> <i>49.90</i>
	priceReg = regexp.MustCompile(`<em>￥</em>[\s]*<i>([\d]+.[\d]+)</i>`)

	//3.5万+</a>
	//ordersReg = regexp.MustCompile(`<p class="productStatus" >[\s]*<span>月成交 <em>([\d]+笔)</em></span>`)

	//<a target="_blank" title="天堂伞 加大加固强效拒水三折晴雨商务伞3311E碰藏青色" href="//item.jd.com/4865226.html" onclick="searchlog(1,4865226,0,2,'','flagsClk=4199040')">
	titleReg = regexp.MustCompile(`<a target="_blank" title="([\S]+)"[\s]*href="([\S]+)"`)

	//src="//img10.360buyimg.com/n7/jfs/t1[\S]+jpg"
	photoReg = regexp.MustCompile(`<img width="220" height="220" data-img="1" src="(//img[\S]+jpg)" data-lazy-img="done" />`)
)

//入参： []byte 爬下来的网页源代码
func ParseGoodList(contents []byte) engine.ParseResult {
	//TODO 商品的价格、url、title不匹配
	titles,links := getTwoMatchs(contents,titleReg)
	prices := getFirstMatchs(contents, priceReg)
	//orders := getFirstMatchs(contents,ordersReg)

	photos := getFirstMatchs(contents,photoReg)
	//for _,p := range titles{
	//	println("titles" + p)
	//}
	//for _,p := range prices{
	//	println("prices" + p)
	//}

	result := engine.ParseResult{
	}

	n := math.Min(float64(len(prices)), float64(len(titles)/2))
	n = math.Min(n,float64(len(photos)))

	for i:=0;i< int(n);i++{
		//log.Printf("id get after reg %v %v %v",prices[i],titles[i*2],links[i*2])

		profile := model.Goods{}

		profile.Eid = 2
		id,err := getGID(links[i*2])
		if err != nil{
			continue
		}
		profile.GoodID = id
		profile.Price = prices[i]
		profile.Url = links[i*2]
		profile.Title = titles[i*2]
		profile.PhotoUrl = photos[i]
		profile.OrderCount = "1321笔"
		result.Items = append(result.Items, profile)
	}
	return result
}



func getGID(url string)(int64,error){
	goodIdReg := regexp.MustCompile(`([\d]+)`)
	m := goodIdReg.FindString(url)

	retString := m
	res,err := strconv.ParseInt(retString, 10, 64)

	return res,err

}

//get value by reg from contents
func getFirstMatchs(contents []byte, re *regexp.Regexp) []string {

	m := re.FindAllSubmatch(contents,-1)

	res := make([]string,0)

	for i := 0;i<len(m);i++{
		//println("i get :" + string(m[i][1]))
		res = append(res, string(m[i][1]))
 	}

	return res
}

//get value by reg from contents
func getTwoMatchs(contents []byte, re *regexp.Regexp) ([]string ,[]string){

	m := re.FindAllSubmatch(contents,-1)

	res := make([]string,0)
	res2 := make([]string,0)

	for i := 0;i<len(m);i++{
		//println("i get :" + string(m[i][1]))
		res = append(res, string(m[i][1]))
		res2 = append(res2, string(m[i][2]))
	}
	return res,res2
}