package api

import (
	. "awesomeProject/middleware"
	myjwt "awesomeProject/middleware/jwt"
	. "awesomeProject/models"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func LoginForm(c *gin.Context) {
	var form Login
	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if form.User != "manu" || form.Password != "123" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}
	//c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	//val, err := database.Rclient.Get("test").Result()
	//if err != nil{
	//	panic(err)
	//}else {
	//	fmt.Println("key", val)
	//}
	SendTest("go_login")
	generateToken(c, form)
}

// GetDataByTime 一个需要token认证的测试接口
func GetDataByTime(c *gin.Context) {
	claims := c.MustGet("claims").(*myjwt.CustomClaims)
	//val, err := database.Rclient.Get("test").Result()
	//if err != nil{
	//	panic(err)
	//}
	//fmt.Println("key", val)

	if claims != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 0,
			"msg":    "token有效",
			"data":   claims,
		})
	}
}

// 生成令牌
func generateToken(c *gin.Context, user Login) {
	j := &myjwt.JWT{
		[]byte("newtrekWang"),
	}
	claims := myjwt.CustomClaims{
		//user.Id,
		user.User,
		//user.Phone,
		jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), // 过期时间 一小时
			Issuer:    "newtrekWang",                   //签名的发行者
		},
	}

	token, err := j.CreateToken(claims)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    err.Error(),
		})
		return
	}

	log.Println(token)

	//data := LoginResult{
	//	User:  user,
	//	Token: token,
	//}
	//val, err := database.Rclient.Get("test").Result()
	//if err != nil{
	//	panic(err)
	//}else {
	//	log.Println("key", val)
	//}
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "登录成功！",
		"data":   token,
	})
	return
}
