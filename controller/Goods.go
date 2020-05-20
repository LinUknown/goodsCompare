package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"imooc.com/crawler"
	"imooc.com/datasource"
	"imooc.com/model"
	"imooc.com/service"
	"imooc.com/util"
	"net/http"
	"strconv"
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
			if err == redis.Nil {
				resGood = append(resGood, g)
			} else if err != nil {
				panic(err)
			}
		}
		service.SaveGoods(resGood)
	}()
	c.JSON(http.StatusOK,gin.H{
		"result":goods,
	})
}

//get 请求，获取历史价格
func PriceHistory(c *gin.Context) {

	eid := c.Query("eid")
	gid := c.Query("gid")
	eID,_ := strconv.Atoi(eid)
	gID,_ :=strconv.ParseInt(gid, 10, 64)


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

//get 请求，获取历史价格
func Like(c *gin.Context) {

	eid := c.Query("eid")
	gid := c.Query("gid")

	id := eid+gid
	err := service.Like(id)
	if err != nil{
		c.JSON(http.StatusOK, gin.H{
			"message":"sys fail",
			"code": 500,
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"code": 200,
	})
	return
}

func GetTop(c *gin.Context){
	goods,scores,err := service.GetTop()
	if err != nil{
		c.JSON(http.StatusOK, gin.H{
			"message":"sys fail",
			"code": 500,
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"code": 200,
		"goods":goods,
		"scores":scores,
	})
}