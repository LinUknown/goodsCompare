package service

import (
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"imooc.com/datasource"
	"imooc.com/model"
	"imooc.com/util"
	"sort"
	"strconv"
	"time"
)

func CompareGoods(goods model.GoodSlice,num int)(resGoods []model.Goods){
	sort.Sort(goods)
	return goods
}

func SaveGoods(goods []model.Goods)(error){
	for _,g := range goods{
		datasource.GetRedis().Set(util.GetGoodPrex(g.GoodID,g.Eid),"2333",24*time.Hour)
	}
	return model.Save(goods)
}

func GetPhoto(goods[] model.Goods)string{

	println("send to me goods is ", len(goods))

	p, _ := plot.New()
	p.Title.Text = "商品：" //+ goods[0].Title + "价格波动图"
	p.X.Label.Text = "时间"
	p.Y.Label.Text = "价格"
	points := make(plotter.XYs, len(goods))
	//tmp := 0.00
	for i:=0;i< len(goods);i++{
		points[i].X ,_= strconv.ParseFloat(goods[i].CreateTime,64)
		fmt.Printf("photo price show =%v", points[i].X)
		points[i].Y ,_= strconv.ParseFloat(goods[i].Price,64)
	}
	plotutil.AddLinePoints(p,"y = x * x", points)
	fileName := fmt.Sprintf("price_history_%v_%v.png",goods[0].Eid,goods[0].GoodID)
	p.Save(8*vg.Inch, 4*vg.Inch, fileName)

	return fileName
}

const ZSET_KEY = "GOODS_LIKE"

func Like(id string)error  {
	newScore, err := datasource.GetRedis().ZIncrBy(ZSET_KEY, 10, id).Result()
	if err != nil {
		fmt.Printf("zincrby failed, err:%v\n", err)
		return err
	}
	fmt.Printf("Golang's score is %f now.\n", newScore)
	return nil
}

func GetTop()([]model.Goods,[]float64,error){
	// 取分数最高的3个
	ret, err := datasource.GetRedis().ZRevRangeWithScores(ZSET_KEY, 0, 10).Result()
	if err != nil {
		fmt.Printf("zrevrange failed, err:%v\n", err)
		return nil,nil,err
	}
	scores := make([]float64,0)
	goods := make([]model.Goods,0)
	for _, z := range ret {
		id := fmt.Sprintf("%v",z.Member)
		eID,_ := strconv.Atoi(id[0:1])
		gID,_ :=strconv.ParseInt(id[1:], 10, 64)
		//fmt.Printf("i get %v %v %v\n", id,eID,gID)
		good,err:= model.GetGoodsByIDAndPID(gID,eID)
		if err != nil || len(good) == 0{
			continue
		}
		scores = append(scores, z.Score)
		goods = append(goods, good[0])
	}
	return goods,scores,nil

}