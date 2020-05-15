package model

import "imooc.com/datasource"

type User struct {
	Id	string `json:"id"`
	Name       string `json:"name"`
	Password	string `json:"password"`
}
//增
func (u *User) Insert() error{
	db := datasource.GetDB()
	err  := db.Create(u).Error
	return err
}

//查
func GetUserByUnameAndPwd(name string,password string) (*User,error){
	u := &User{}
	db := datasource.GetDB()
	err := db.Where("name = ? and password = ?",name,password).Find(u).Error
	return u,err
}
//查
func GetUserByUname(name string) (*User,error){
	u := &User{}
	db := datasource.GetDB()
	err := db.Where("name = ?",name).Find(u).Error
	return u,err
}