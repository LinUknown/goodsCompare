package model

import (
	"fmt"
	"imooc.com/datasource"
	"imooc.com/util"
	"strconv"
	"time"
)

type Goods struct {
	Eid 	int //tmall:1  JD：2 亚马逊:3
	GoodID	int64 //商品Id，和ECId作为唯一标识符号
	Title       string
	OrderCount    string
	Price        string
	Url 		string
	CreateTime string
	PhotoUrl string
}
type GoodSlice []Goods

func (s GoodSlice) Len() int { return len(s) }
func (s GoodSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s GoodSlice) Less(i, j int) bool {
	p1,_ := strconv.ParseFloat(s[i].Price, 64)
	p2,_ := strconv.ParseFloat(s[j].Price, 64)
	return p1 < p2
}

func Save(goods []Goods)error{
	db := datasource.GetDB()
	for _,g := range goods{
		g.CreateTime = fmt.Sprintf(time.Now().Format(util.TIME_TEMPLATE))
		err := db.Save(g).Error
		if err != nil{
			continue
		}
	}
	return nil
}

func GetGoodsByIDAndPID(id int64,eid int) ([]Goods,error){
	goods := make([]Goods,0)
	db := datasource.GetDB()
	err := db.Where("good_id=? AND eid=?",id,eid).Find(&goods).Error
	return goods,err
}
