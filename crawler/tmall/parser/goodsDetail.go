package parser

import (
	"imooc.com/crawler/engine"
	"imooc.com/model"
	"math"
	"regexp"
	"strconv"
)

var (
	// <p class="productPrice"><em title="34.99"><b>&yen;</b>34.99</em></p>
	priceReg = regexp.MustCompile(`<p class="productPrice">[\s]*<em title="([\d]+.[\d]+)"><b>&yen;</b>[\d]+.[\d]+</em>[\s]*</p>`)

	//<p class="productTitle"><a href="//detail.tmall.com/item.htm?id=1718" target="_blank" title="天堂伞超大雨伞折叠晴雨两用伞三折防晒防紫外线遮阳伞太阳伞男女" data-p="1-11" > 天堂<span class=H>伞</span>超大<span class=H>雨伞</span>折叠晴雨两用<span class=H>伞</span>三折防晒防紫外线遮阳<span class=H>伞</span>太阳<span class=H>伞</span>男女 </a> </p>
	linkReg = regexp.MustCompile(`<p class="productTitle">[\s]*<a href="([\S]+)" target="_blank" title="([\S]+)"`)

	//<p class="productStatus" ><span>月成交 <em>4629笔</em></span>
	ordersReg = regexp.MustCompile(`<p class="productStatus" >[\s]*<span>月成交 <em>([\d]+笔)</em></span>`)
	//<img src="//img.alicdn.com/bao/uploaded/i2/2207264430349/O1CN01cuvPyz1ERsokj3JL0_!!2207264430349.jpg">
	photoReg = regexp.MustCompile(`<img[\s]*src=[\s]*"(//img.alicdn.com/bao[\S]+jpg)" />`)
	)

//入参： []byte 爬下来的网页源代码
func ParseGoodList(contents []byte) engine.ParseResult {

	prices := getFirstMatchs(contents, priceReg)
	links,titles := getTwoMatchs(contents,linkReg)
	orders := getFirstMatchs(contents,ordersReg)
	photos:= getFirstMatchs(contents,photoReg)

	//for _,p := range photos{
	//	fmt.Printf("get photo = %v", p)
	//}
	result := engine.ParseResult{
	}

	n := math.Min(float64(len(prices)), float64(len(orders)))
	n = math.Min(n,float64(len(links)))
	n = math.Min(n,float64(len(photos)))


	for i:=0;i< int(n);i++{
		profile := model.Goods{}

		profile.Eid = 1
		id,err := getGID(links[i])
		if err != nil{
			continue
		}
		profile.GoodID = id
		profile.Price = prices[i]
		profile.Url = links[i]
		profile.Title = titles[i]
		profile.OrderCount = orders[i]
		profile.PhotoUrl = photos[i]
		result.Items = append(result.Items, profile)
	}
	return result
}



func getGID(url string)(int64,error){
	goodIdReg := regexp.MustCompile(`id=([\d]+)`)
	m := goodIdReg.FindString(url)

	retString := m[3:]
	res,err := strconv.ParseInt(retString, 10, 64)

	//log.Printf("get id = %d, m = %v,err=%v",res,m,err)
	return res,err

}

//get value by reg from contents
func getFirstMatchs(contents []byte, re *regexp.Regexp) []string {

	m := re.FindAllSubmatch(contents,-1)

	res := make([]string,0)

	for i := 0;i<len(m);i++{
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
		res = append(res, string(m[i][1]))
		res2 = append(res2, string(m[i][2]))
	}
	return res,res2
}