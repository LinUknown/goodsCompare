package controller

import (
	"github.com/gin-gonic/gin"
	"imooc.com/crawler"
	"imooc.com/datasource"
	"imooc.com/model"
	"imooc.com/service"
	"imooc.com/util"
	"net/http"
)

// localhost:9090/user/query?id=2&name=hello
func Search(c *gin.Context) {
	println(">>>> Search action start <<<<")
	key := c.Query("search")
	result := crawler.Search(key)
	goods := service.CompareGoods(result,10)
	go func() {
		//过滤掉1天内更新的数据
		resGood := make([]model.Goods,0)
		for _,g := range goods{
			_,err := datasource.GetRedis().Get(util.GetGoodPrex(g.GoodID,g.Eid)).Result()
			if err != nil{
				//redis内不存在
				resGood = append(resGood, g)
			}
		}
		//service.SaveGoods(resGood)
	}()
	c.JSON(http.StatusOK,gin.H{
		"result":goods,
	})
}

//get 请求，获取历史价格
func PriceHistory(c *gin.Context) {

	//id := c.PostForm("id")
	//
	//eID,_ := strconv.Atoi(id[0:1])
	//gID,_ :=strconv.ParseInt(id[1:], 10, 64)

	eID := 1
	gID := 111


	goods,err:= model.GetGoodsByIDAndPID(int64(gID),eID)
	//通过id和pid，查出不同时间点下的商品记录，然后构图
	if err != nil{
		c.JSON(http.StatusOK, gin.H{
			"message":"sys fail",
			"code": 500,
		})
		return
	}

	if goods==nil || len(goods) == 0{
		c.JSON(http.StatusOK, gin.H{
			"message":"get data fail",
			"code": 500,
		})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"code": 200,
		"result":goods,
	})
	return


}