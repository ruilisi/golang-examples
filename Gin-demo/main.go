package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"Gin-demo/sql"

	"github.com/gin-gonic/gin"
)

type User struct {
	id       string
	password string
}

//func LinkRedis() {
//	client := redis.NewClient(&redis.Options{
//		Addr:     "localhost://5600",
//		Password: "", // no password set
//		DB:       0,  // use default DB
//	})
//	return client
//}

func PhoneCheckRule(a string) bool {
	if regexp.MustCompile(`^1[3456789][0-9]{9}$`).MatchString(a) {
		return true
	} else {
		return false
	}
}

func EmailCheckRule(a string) bool {
	if regexp.MustCompile(`^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$`).MatchString(a) {
		return true
	} else {
		return false
	}
}

func PasswordCheck(a string) bool {
	reg1 := regexp.MustCompile(`[a-z]+`)
	reg2 := regexp.MustCompile(`[A-Z]+`)
	reg3 := regexp.MustCompile(`[0-9]+`)
	if reg1.MatchString(a) && reg2.MatchString(a) && reg3.MatchString(a) && len(a) >= 8 {
		return true
	} else {
		return false
	}

}

func signupInit(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", nil)
}

func signinInit(c *gin.Context) {
	c.HTML(http.StatusOK, "signin.html", nil)
}

func profileInit(c *gin.Context) {
	//	client := LinkRedis
	//val, err := client.Get()
	username, _ := c.Cookie("username")
	if username == "" {
		c.Redirect(http.StatusMovedPermanently, "signin")
	} else {
		c.HTML(http.StatusOK, "profile.html", gin.H{
			"text": username,
		})
	}
}

func signupHandler(c *gin.Context) {
	var user User
	user.id = c.PostForm("id")
	user.password = c.PostForm("password")
	fmt.Println(user.id)
	fmt.Println(user.password)
	if !PhoneCheckRule(user.id) && !EmailCheckRule(user.id) {
		func(c *gin.Context) {
			c.HTML(http.StatusOK, "signup.html", gin.H{
				"text": "非手机号或邮箱"})
		}(c)
	} else if !PasswordCheck(user.password) {

		func(c *gin.Context) {
			c.HTML(http.StatusOK, "signup.html", gin.H{
				"text": " 密码不符合规范"})
		}(c)
	} else {
		if PhoneCheckRule(user.id) {
			sql.SqlInsert(user.id, "", user.password)
			c.Redirect(http.StatusMovedPermanently, "signin")
		} else {
			sql.SqlInsert("", user.id, user.password)
			c.Redirect(http.StatusMovedPermanently, "signin")
		}
		//sql.SqlClose()
	}
}

func signinHanlder(c *gin.Context) {
	var user User
	user.id = c.PostForm("id")
	user.password = c.PostForm("password")
	fmt.Println(user.id)
	btn := c.PostForm("flag")
	fmt.Println(btn)
	if btn == "2" {
		c.Redirect(http.StatusMovedPermanently, "signup")
	} else {
		check_pwd := sql.SqlSelect(user.id, user.id)
		if strings.Compare(strings.Trim(check_pwd, " "), user.password) == 0 {
			//			client := LinkRedis()
			//			err := client.Set(user.id, user.password, 0).Err()
			//if err!=nil{
			//				panic(err)
			//			}
			c.SetCookie("username", user.id, 300, "/", "localhost", false, true)
			c.Redirect(http.StatusMovedPermanently, "/profile")
		} else {
			func(c *gin.Context) {
				c.HTML(http.StatusOK, "signin.html", gin.H{
					"text": "Login Failed"})
			}(c)
		}
	}
}

func main() {
	sql.SqlOpen()
	defer sql.SqlClose()
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	r.LoadHTMLGlob("*.html")
	r.GET("/signup", signupInit)
	r.POST("/signup", signupHandler)
	r.GET("/signin", signinInit)
	r.POST("/signin", signinHanlder)
	r.GET("/profile", profileInit)
	r.Run("0.0.0.0:3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
