package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"imooc.com/datasource"
	"imooc.com/model"
	"imooc.com/util"
	"log"
	"net/http"
	"time"
)

func init()  {
	log.Println(">>>> get database connection start <<<<")
	//db = database.GetDataBase()
}

// localhost:9090/user/query?id=2&name=hello
func QueryParam(context *gin.Context) {
	println(">>>> query user by url params action start <<<<")
	u := model.User{}

	if context.BindJSON(&u) == nil{
		log.Printf("%v",u)
	}



	//context.Bind(&u)
	//err := u.Login()
	//if err != nil{
	//	context.JSON(304,gin.H{
	//		"result":model.User{},
	//	})
	//	return
	//}

	//checkError(err)
	context.JSON(200,gin.H{
		"result":u,
	})
}

func Register(c *gin.Context){
	println(">>>> register action start <<<<")
	userName := c.PostForm("name")
	passWord := c.PostForm("password")
	//email := c.PostForm("email")

	_,err := model.GetUserByUname(userName)
	if err == nil{
		c.JSON(http.StatusOK, gin.H{
			"message":"user exist",
			"code": 200,
		})
		return
	}
	if err != nil  && err != gorm.ErrRecordNotFound{
		log.Printf("User controller get user fail, username=%v, err=%v",userName,err)
		c.JSON(http.StatusOK, gin.H{
			"message":"sys err",
			"code": 500,
		})
		return
	}
	newUser := model.User{Name:userName,Password:passWord}
	err = newUser.Insert()
	if err != nil{
		c.JSON(http.StatusOK, gin.H{
			"message":"sys fail",
			"code": 500,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":"success",
		"code": 200,
		"data":"/goods/get_search",
	})
}

// form表单提交
func Login(c *gin.Context) {
	println(">>>> bind form post params action start <<<<")

	userName := c.PostForm("name")
	passWord := c.PostForm("password")

	u,err := model.GetUserByUnameAndPwd(userName,passWord)

	if err != nil{
		c.JSON(http.StatusOK, gin.H{
			"result":"success",
			"code": 500,
		})
		return
	}
	token := util.GetToken(userName)
	log.Printf("login success,ins cookie = %v",token)
	c.SetCookie("token", token, 3600, "/", "localhost", http.SameSiteLaxMode, false,true)
	err = datasource.GetRedis().Set(token,"2333",10*time.Minute).Err()
	if err != nil {
		panic(err)
	}
	log.Printf("user:%v,login",u.Name)
	c.JSON(http.StatusOK, gin.H{
		"result":"success",
		"code": 200,
		"data":"/goods/get_search",
	})

	//	// 重定向
	//	context.Redirect(http.StatusMovedPermanently,"/file/view")
	//}

}

// 跳转html
func RenderForm(context *gin.Context) {
	println(">>>> render to html action start <<<<")

	context.Header("Content-Type", "text/html; charset=utf-8")
	context.HTML(200,"insertUser.html",gin.H{})
}

func checkError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}