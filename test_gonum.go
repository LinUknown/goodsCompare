package main

import (
	"fmt"
	"imooc.com/model"
	"imooc.com/util"
	"time"
)

func main(){

	goods := make([]model.Goods,0)

	t := time.Now()

	for i:=0;i<50;i++{
		t2:= t.AddDate(0,0,1).Format(util.TIME_TEMPLATE)
		goods = append(goods, model.Goods{Eid:1,GoodID:111,CreateTime:t2,Price:fmt.Sprintf("%d",i),
		PhotoUrl:" //img12.360buyimg.com/n7/jfs/t1/108743/27/13348/119857/5e9ff842E654999aa/59ffe4ad57fb1015.jpg",
		Title:"京东·急速达立即抢购(此商品不参加上述活动)"})
		}
	model.Save(goods)
	//service.GetPhoto(goods)

}
