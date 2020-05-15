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
)

func CompareGoods(goods model.GoodSlice,num int)(resGoods []model.Goods){
	sort.Sort(goods)
	return goods
}

func SaveGoods(goods []model.Goods)(error){
	for _,g := range goods{
		datasource.GetRedis().Set(util.GetGoodPrex(g.GoodID,g.Eid),"2333",3600*12)
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