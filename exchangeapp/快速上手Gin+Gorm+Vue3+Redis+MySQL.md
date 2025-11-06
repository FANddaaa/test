# 快速上手Gin+Gorm+Vue3+Redis+MySQL
原视频链接：[零基础Go+Gin+Gorm+Vue3+Redis+MySQL快速上手实战教程](https://www.bilibili.com/video/BV1BY4UefEkM)

## Section1 环境配置及前置重要概念
### 1.1 基本环境配置——MySQL
略

### 1.2 基本概念——前后端分离&RESTful API

#### 1.2.1 什么是前后端分离
将传统的Web应用从一个整体拆分成两个独立的部分：前端(Front-end)和后端(Back-end)

* 前端：主要负责用户界面(UI)的呈现和交互逻辑
* 后端：主要负责业务逻辑、数据处理与数据库的交互

优点：

* 提高开发效率
* 增强系统可维护性
* 提升用户体验
* 技术栈灵活

#### 1.2.2 前后端分离中API的作用
API：Application Programming Interface(译：应用程序接口)，连接前后端的桥梁

#### 1.2.3 RESTful API简述
REST：Representational State Transfer(译：表述性状态转移)，是一种软件架构风格而非标准，RESTful API即为一种REST风格的接口。

很多年前的案例：

```
http://www.example.com/get_rates?name=usdeur
http://www.example.com/update_rates?...
http://www.example.com/delete_rates?...
```

满足REST风格的接口是这样的：

```
GET http://www.example.com/rates
POST http://www.example.com/rates
Data: name=usdeur
```

##### GET、POST、PUT和DELETE
1. GET：取出资源(一项或多项)
2. POST：新建一个资源
3. PUT：更新资源(客户端提供完整资源数据)
4. PATCH：更新资源(客户端提供需要修改的资源数据)
5. DELETE：删除资源

##### RESTful API核心要点
1. 资源(resources)：每一个URL代表一个资源(将互联网的资源抽象为资源，获取资源的方式定义为方法)，如`api/article/3`，以资源为基础，资源可以是图片，更多以JSON为载体，如`{"fromCurrency": "USD", "toCurrency": "KZT", "rate": 479}`
2. 使用HTTP方法(GET、POST、PUT和DELETE等)表示对资源的操作。
3. 使用HTTP状态码(HTTP Status Code)表示请求的结果(响应状态码)。如200(成功)、404(未找到)、500(服务器错误)、401(Unauthorized)：未经授权、201(Created)：请求已被成功处理且创建了新的资源。

##### RESTful API建议规范
URL重要概念：[MDN Web Docs](https://developer.mozilla.org/zh-CN/docs/Learn/Common_questions/Web_mechanics/What_is_a_URL)
![What_is_a_URL](https://developer.mozilla.org/zh-CN/docs/Learn/Common_questions/Web_mechanics/What_is_a_URL/mdn-url-all.png)
需要注意的是：URL是URI的子集  
`/{version}/{resources}/{resources_id}`  
例子：`api/v1/articles/3`

1. 使用名词表示资源
2. 使用层次结构
3. 不使用文件扩展名，正确使用`/`表示层级关系

优势：遵循RESTful API建议规范可以提高API的可读性、可维护性和可拓展性。

#### 其他概念
##### 请求地址、基础地址、接口地址(API)  
`请求地址`在大部分情况即为`基础地址` + `api/接口地址`  
例子：基础地址：`http://www.example.com`  
接口地址：`api/v1/user/inkkaplumchannel`

##### 数据格式
JSON，也是目前最主流的

``` json
{
  "_id": 1,
  "fromCurrency": "USD",
  "toCurrency": "KZT",
  "rate": 479,
  "date": "2024-08-30T22:45:42.8774003+08:00"
}
```

### 1.3 基本概念——前后端分离&MVC&前后端路由&Gin&Gorm
#### 1.3.1 MVC
MVC：全称Model-View-Controller，是一种常用的软件架构模式，旨在将应用程序的关注点分离，提高代码的可维护性。  
建设我们用Go+Gin+Gorm+Vue开发了一个简易的博客网站：

```
哔哩哔哩InkkaPlum频道的个人博客
├── backend/
│   ├── main.go
│   ├── controllers/
│   │   ├── auth.go
│   │   ├── article.go
│   │   └── ...
│   ├── models/
│   │   ├── users.go
│   │   ├── article.go
│   │   └── ...
│   ├── routers/
│   │   └── router.go
│   └── config/
│       └── config.yaml
└── frontend/
    ├── src/
    │   ├── App.vue
    │   ├── components/
    │   ├── views/
    │   └── router/
    ├── index.html
    └── ...
```

##### Model-模型
`models/article.go`

``` go
type Article struct {
    Content string
    CreateAt time.Time
    UpdateAt time.Time
    ...
}
```

##### View-视图
`components/ArticleForm.vue`

``` html
<template>
  <form @submit.prevent="onSubmit">
    <input type="text" v-model="title" placeholder="标题" />
    <textarea v-model="content" placeholder="内容"></textarea>
    <button type="submit">提交</button>
  </form>
</template>

<script>
...
</script>

```
``` html
<h2>{{ article.Title }}</h2>
<p>{{ article.Preview }}</p>
```

##### Controller-控制器
`controllers/article.go`

``` go
func CreateArticleHandler(c *gin.Context) {
    var article models.Article
    if err := c.ShouldBingJSON(&article); err != nil {
        // ... 处理错误
    }
    // ...
    
    c.JSON(http.StatusCreated, article)
}
```

##### MVC的工作流程
1. 用户对界面进行操作，如点击点赞按钮
2. View感知这些事件，通过Controller进行处理(`POST`)
3. Controller处理相关业务，对Model的业务数据进行更新
4. View更新用户界面，赞+1

用户操作->View->Controller->Model->View，MVC的通信过程都是**单向**的。

#### 1.3.2 前端路由和后端路由
略

#### 1.3.3 Gin和Gorm
Gin和Gorm经常一起使用来构建Go Web应用程序
##### Gin框架
使用Go语言开发的Web框架，主要特点：

* 高性能
* 中间件支持
* 路由分组

##### Gorm
Gorm是Golang中最流行的ORM(对象关系映射)库(ORM Library)

* ORM
ORM - Object Relational mapping，优势：

* 简单易用
* 自动迁移
* 支持多种数据库

## Section2 Gin&Gorm基础
### 2.1 活用Viper读取配置文件&yaml文件说明
```
文件树
├── config/ <- 新增
│   ├── config.go <- 新增
│   └── config.yml <- 新增
└── main.go <- 新增
```
``` go
// config.go
package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
    App struct {
        Name string
        Port string
    }
    Database struct {
        Host string
        Port string
        User string
        Password string
        Name string
    }
}

var AppConfig *Config

func InitConfig() {
    viper.SetConfigName("config")
    viper.SetConfigType("yml")
    viper.AddConfigPath("./config")
    
    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("Error reading config file: %v", err)
    }
    
    AppConfig = &Config{}
    
    if err := viper.Unmarshal(AppConfig); err != nil {
        log.Fatalf("Unable to decode into struct: %v", err)
    } 
}
```
``` yaml
# config.yml
app:
  name: CurrencyExchangeApp
  port: :3000

database:
  host: localhost
  port: :3306
  user: your_username
  password: your_password
  name: currency_exchange_db
```
```
go mod init

go get -u github.com/gin-gonic/gin
go get github.com/spf13/viper
```
``` go
// main.go
package main

import (
    "fmt"

    "exchangeapp/config"
)

func main() {
    config.InitConfig()
    fmt.Println(config.AppConfig.App.Port)
}
```
### 2.2 初次体验Gin
```
文件树
├── config/
│   ├── config.go
│   └── config.yml
└── main.go <- 修改
```
``` go
// main.go
package main

import (
    "github.com/gin-gonic/gin"

    "exchangeapp/config"
)

func main() {
    config.InitConfig()
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
    // listen and serve on 0.0.0.0:3000(default: 8080)
    r.Run(config.AppConfig.App.Port)
}
```
### 2.3 Gin路由的基本配置&路由分组&其他概念
```
文件树
├── config/
│   ├── config.go
│   └── config.yml
├── router/ <- 新增
│   └── router.go <- 新增
└── main.go <- 修改
```
``` go
// router.go
package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	auth := r.Group("/api/auth")
	{
		auth.POST("/login", func(ctx *gin.Context) {
			ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
				"message": "login success",
			})
		})
		auth.POST("/register", func(ctx *gin.Context) {
			ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
				"message": "register success",
			})
		})
	}

	return r
}
```
``` go
// main.go
package main

import (
	"exchangeapp/config"
	"exchangeapp/router"
)

func main() {
	config.InitConfig()
	r := router.SetupRouter()

	port := config.AppConfig.App.Port
	if port == "" {
		port = ":8080"
	}
	// listen and serve on 0.0.0.0:8888(default: 8080)
	r.Run(port)
}
```
### 2.4 Gorm基础&Gorm连接数据库初始化&global go文件说明
官方文档(中文)：<https://gorm.io/zh_CN/docs/>

基本配置

```
go get -u gorm.io/gorm
```
```
文件树
├── config/
│   ├── config.go <- 修改
│   ├── config.yml <- 修改
│   └── db.go  <- 新增
├── global/ <- 新增
│   └── global.go <- 新增
├── router/
│   └── router.go
└── main.go
```
``` go
// config.go
package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Name string
		Port string
	}
	Database struct {
		Dsn          string
		MaxIdleConns int
		MaxOpenConns int
	}
}

var AppConfig *Config

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	AppConfig = &Config{}

	if err := viper.Unmarshal(AppConfig); err != nil {
		log.Fatalf("Unable to decode into struct: %v", err)
	}

	initDB()
}
```
``` yaml
# config.yml
app:
  name: CurrencyExchangeApp
  port: :8888

database:
  dsn: root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local
  maxIdleConns: 10
  maxOpenConns: 100
```
``` go
// db.go
package config

import (
	"exchangeapp/global"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDB() {
	dsn := AppConfig.Database.Dsn
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	sqlDB, err := db.DB()

	sqlDB.SetMaxIdleConns(AppConfig.Database.MaxIdleConns)
	sqlDB.SetMaxOpenConns(AppConfig.Database.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err != nil {
		log.Fatalf("Failed to configure database: %v", err)
	}

	global.Db = db
}
```
``` go
// global.go
package global

import "gorm.io/gorm"

var (
	Db *gorm.DB
)
```
## Section3 实现注册登录功能
### 3.1 Gorm模型定义&Gorm约定说明
```
文件树
├── config/
│   ├── config.go
│   ├── config.yml
│   └── db.go
├── global/
│   └── global.go
├── models/ <- 新增
│   └── user.go <- 新增
├── router/
│   └── router.go
└── main.go
```
``` go
// user.go
package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
}
```
### 3.2 模型绑定&HTTP请求报文&概念分析
HTTP请求报文包含三个部分：请求行 + 请求头 + 请求体

**请求行**包含三个内容

1. method：**GET**、**POST**等
2. request-URI：**http://www.example.com/api/v1/test**
3. HTTP-version：**HTTP/1.1**

**请求头**(Request Header)：由键值对组成，每行为一对，key/value之间通过冒号分割

**请求体**：我们的JSON内容，发送给服务器的数据

HTTP响应也有行、头、体的概念

状态行包含三个内容

1. HTTP-version：**HTTP/1.1**
2. 状态码：**200**
3. 状态码的文本描述：**OK**

```
文件树
├── config/
│   ├── config.go
│   ├── config.yml
│   └── db.go
├── controllers/ <- 新增
│   └── auth_controller.go <- 新增
├── global/
│   └── global.go
├── models/
│   └── user.go
├── router/
│   └── router.go
└── main.go
```
``` go
// auth_controller.go
package controllers

import (
	"net/http"
	
	"github.com/gin-gonic/gin"

	"exchangeapp/models"
)

func Register(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

```
### 3.3 密码加密功能实现&Bcrypt概述
#### 加密逻辑 - Bcrypt Package
Bcrypt是一种用于密码哈希的加密算法，它是基于Blowfish算法的加强版，被广泛应用于存储密码和进行身份验证。

优势：

* 安全性高：Bcrypt采用了`Salt`和`Cost`两种机制，可有效地防止彩虹表攻击和暴力破解攻击，从而保证安全性。

```
go get -u golang.org/x/crypto/bcrypt
```
```
文件树
├── config/
│   ├── config.go
│   ├── config.yml
│   └── db.go
├── controllers/
│   └── auth_controller.go <- 修改
├── global/
│   └── global.go
├── models/
│   └── user.go
├── router/
│   └── router.go
├── utils/ <- 新增
│   └── utils.go <- 新增
└── main.go
```
``` go
// auth_controller.go
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"exchangeapp/models"
	"exchangeapp/utils"
)

func Register(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPwd, err := utils.HashPassword(user.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user.Password = hashedPwd
}
```
``` go
// utils.go
package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), 12)
	return string(hash), err
}
```
### 3.4 JWT概述&生成JWT
#### JWT定义
官网：<https://jwt.io/introduction>

JSON Web Token(JWT | JSON网络令牌)是一种开放标准(RFC 7519)，用于在网络应用环境间安全地传递声明(claims)。JWT是一种紧凑且自包含的方式，用于作为JSON对象在各方之间安全地传递信息。由于其信息是经过数字签名的，所以可以确保发送的数据在传输过程中未被篡改。

组成成分：Header + Payload + Signature

可用的库：<https://jwt.io/libraries>

```
go get github.com/golang-jwt/jwt/v5
```
```
文件树
├── config/
│   ├── config.go
│   ├── config.yml
│   └── db.go
├── controllers/
│   └── auth_controller.go <- 修改
├── global/
│   └── global.go
├── models/
│   └── user.go
├── router/
│   └── router.go
├── utils/
│   └── utils.go <- 修改
└── main.go
```
``` go
// auth_controller.go
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"exchangeapp/models"
	"exchangeapp/utils"
)

func Register(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPwd, err := utils.HashPassword(user.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user.Password = hashedPwd

	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
```
``` go
// utils.go
package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func HashPassword(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), 12)
	return string(hash), err
}

func GenerateJWT(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24 * 3).Unix(),
	})

	signedToken, err := token.SignedString([]byte("secret"))
	return "Bearer " + signedToken, err
}
```
### 3.5 CRUD说明&Gorm自动迁移&创建记录
#### CURD简述
代表Create(创建) + Read(读取) + Update(更新) + Delete(删除)

`AutoMigrate`会创建表：`db.AutoMigrate(&User{})`

```
文件树
├── config/
│   ├── config.go
│   ├── config.yml
│   └── db.go
├── controllers/
│   └── auth_controller.go <- 修改
├── global/
│   └── global.go
├── models/
│   └── user.go
├── router/
│   └── router.go
├── utils/
│   └── utils.go
└── main.go
```
``` go
// auth_controller.go
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"exchangeapp/global"
	"exchangeapp/models"
	"exchangeapp/utils"
)

func Register(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPwd, err := utils.HashPassword(user.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user.Password = hashedPwd

	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := global.Db.AutoMigrate(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := global.Db.Create(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
```
### 3.6 实现登陆全部功能&结构体标签概述
```
文件树
├── config/
│   ├── config.go
│   ├── config.yml
│   └── db.go
├── controllers/
│   └── auth_controller.go <- 修改
├── global/
│   └── global.go
├── models/
│   └── user.go
├── router/
│   └── router.go <- 修改
├── utils/
│   └── utils.go
└── main.go
```
``` go
// auth_controller.go
// ...
func Login(ctx *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User

	if err := global.Db.Where(&models.User{Username: input.Username}).First(&user).Error; err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "wrong credentials"})
		return
	}

	if !utils.CheckPassword(input.Password, user.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "wrong credentials"})
		return
	}

	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
```
``` go
// router.go
package router

import (
	"exchangeapp/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	auth := r.Group("/api/auth")
	{
		auth.POST("/login", controllers.Login)
		auth.POST("/register", controllers.Register)
	}

	return r
}
```
## Section4 实现其他功能
### 4.1 实现创建汇率数据的全部功能
```
文件树
├── config/
│   ├── config.go
│   ├── config.yml
│   └── db.go
├── controllers/
│   ├── auth_controller.go
│   └── exchange_rate_controller.go <- 新增
├── global/
│   └── global.go
├── models/
│   ├── exchange_rate.go <- 新增
│   └── user.go
├── router/
│   └── router.go
├── utils/
│   └── utils.go
└── main.go
```
``` go
// exchange_rate_controller.go
package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"exchangeapp/global"
	"exchangeapp/models"
)

func CreateExchangeRate(ctx *gin.Context) {
	var exchangeRate models.ExchangeRate
	if err := ctx.ShouldBindJSON(&exchangeRate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exchangeRate.Date = time.Now()
	if err := global.Db.AutoMigrate(&exchangeRate); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := global.Db.Create(&exchangeRate).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, exchangeRate)
}
```
``` go
// exchange_rate.go
package models

import "time"

type ExchangeRate struct {
	ID           uint64    `gorm:"primarykey" json:"_id"`
	FromCurrency string    `json:"fromCurrency" binding:"required"`
	ToCurrency   string    `json:"toCurrency" binding:"required"`
	Rate         float64   `json:"rate" binding:"required"`
	Date         time.Time `json:"date"`
}
```
### 4.2 实现获取汇率数据的全部功能&Go切片说明
```
文件树
├── config/
│   ├── config.go
│   ├── config.yml
│   └── db.go
├── controllers/
│   ├── auth_controller.go
│   └── exchange_rate_controller.go <- 修改
├── global/
│   └── global.go
├── models/
│   ├── exchange_rate.go
│   └── user.go
├── router/
│   └── router.go
├── utils/
│   └── utils.go
└── main.go
```
``` go
// exchange_rate_controller.go
// ...
func GetExchangeRates(ctx *gin.Context) {
	var exchangeRates []models.ExchangeRate
	if err := global.Db.Find(&exchangeRates).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, exchangeRates)
}
```
### 4.3-4.4 Gin中间件概述&验证JWT
中间件是为了过滤路由而发明的一种机制，也就是http请求来到时先经过中间件，再到具体的处理函数。请求到达我们定义的处理函数之前，拦截请求并进行相应处理(比如：权限验证、数据过滤等)，可以类比为**前置拦截器**或**前置过滤器**。

自定义中间件可参考Gin的官方文档：<https://gin-gonic/zh-cn/docs/examples/custom-middleware/>

```
文件树
├── config/
│   ├── config.go
│   ├── config.yml
│   └── db.go
├── controllers/
│   ├── auth_controller.go
│   └── exchange_rate_controller.go
├── global/
│   └── global.go
├── middlewares/ <- 新增
│   └── auth_middleware.go <- 新增
├── models/
│   ├── exchange_rate.go
│   └── user.go
├── router/
│   └── router.go
├── utils/
│   └── utils.go <- 修改
└── main.go
```
``` go
// auth_middleware.go
package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	
	"exchangeapp/utils"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization Header"})
			ctx.Abort()
			return
		}
		username, err := utils.ParseJWT(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			ctx.Abort()
			return
		}
		
	}
}
```
``` go
// utils.go
// ...
func ParseJWT(tokenString string) (string, error) {
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte("secret"), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username, ok := claims["username"].(string)
		if !ok {
			return "", errors.New("username is not a string")
		}
		return username, nil
	}

	return "", err
}
```
### 4.5 Gin中间件流程控制概述&Set/Get说明&修改路由实现汇率部分全部功能
```
文件树
├── config/
│   ├── config.go
│   ├── config.yml
│   └── db.go
├── controllers/
│   ├── auth_controller.go
│   └── exchange_rate_controller.go
├── global/
│   └── global.go
├── middlewares/
│   └── auth_middleware.go <- 修改
├── models/
│   ├── exchange_rate.go
│   └── user.go
├── router/
│   └── router.go <- 修改
├── utils/
│   └── utils.go
└── main.go
```
``` go
// auth_middleware.go
package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"exchangeapp/utils"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization Header"})
			ctx.Abort()
			return
		}
		username, err := utils.ParseJWT(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			ctx.Abort()
			return
		}

		ctx.Set("username", username)
		ctx.Next()
	}
}
```
``` go
// router.go
package router

import (
	"github.com/gin-gonic/gin"

	"exchangeapp/controllers"
	"exchangeapp/middlewares"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	auth := r.Group("/api/auth")
	{
		auth.POST("/login", controllers.Login)
		auth.POST("/register", controllers.Register)
	}

	api := r.Group("/api")
	api.GET("/exchangeRates", controllers.GetExchangeRates)
	api.Use(middlewares.AuthMiddleWare())
	{
		api.POST("/exchangeRates", controllers.CreateExchangeRate)
	}

	return r
}
```
### 4.6-4.7 路由参数&实现文章的全部功能
```
文件树
├── config/
│   ├── config.go
│   ├── config.yml
│   └── db.go
├── controllers/
│   ├── article_controller.go <- 新增
│   ├── auth_controller.go
│   └── exchange_rate_controller.go
├── global/
│   └── global.go
├── middlewares/
│   └── auth_middleware.go
├── models/
│   ├── article.go <- 新增
│   ├── exchange_rate.go
│   └── user.go
├── router/
│   └── router.go
├── utils/
│   └── utils.go
└── main.go
```
``` go
// article_controller.go
package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"exchangeapp/global"
	"exchangeapp/models"
)

func CreateArticle(ctx *gin.Context) {
	var article models.Article
	if err := ctx.ShouldBindJSON(&article); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := global.Db.AutoMigrate(&article); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := global.Db.Create(&article).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, article)
}

func GetArticles(ctx *gin.Context) {
	var articles []models.Article
	if err := global.Db.Find(&articles).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	ctx.JSON(http.StatusOK, articles)
}

func GetArticlesByID(ctx *gin.Context) {
	id := ctx.Param("id")

	var article models.Article

	if err := global.Db.Where("id = ?", id).First(&article).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, article)
}
```
``` go
// article.go
package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title   string `binding:"required"`
	Content string `binding:"required"`
	Preview string `binding:"required"`
}
```
## Section5 Redis部分
### 5.1 Redis简介&配置go-redis
#### Redis
Redis官网：<https://redis.io/>  
安装Redis：<https://github.com/tporadowski/redis/releases>  
选择`Redis-x64-5.0.14.1.si`即可

定义：Redis是一个高性能NoSQL的、用C实现、可基于内存亦可持久化的Key-Value数据库(键值对存储数据库)，并提供多种语言的API。

与MySQL数据库不同的是：为了实现数据的持久化存储，MySQL将数据存储到了磁盘中，Redis的数据是存储在内存中的，它的读写速度非常快。

一些基础概念：Redis有如字符串(String)、列表(List)、集合(Set)、有序集合(ZSet/Sorted Set)这样的基础数据类型。`localhost:6379`——在安装过程中设置的端口。

#### 持久化
定义：持久化是指将数据写入持久存储(durable storage)，如固态硬盘(SSD)。

Redis提供了一系列选项：

1. RDB(Redis Database)：RDB持久化通过在指定的时间间隔内创建数据集的快照来保存数据。
2. AOF(Append Only File)：AOF持久化记录服务器接收到的每一个写操作，并将这些操作追加到日志文件中。
3. 无持久化：完全禁用Redis的持久化机制，意味着Redis只会在内存中存储数据。
4. AOF+RDB组合：可以在同一个实例中同时启用RDB和AOF持久化。

设置方法：在redis目录下`C:\Program Files\Redis`找到`redis.windows.conf`即可。

RDB 部分：  
默认情况下，Redis会将数据集的快照保存在磁盘上一个名为`dump.rdb`的二进制文件中。  
在`redis.windows.conf`文件内找到：`save 60 1000`

AOF 部分：  
找到`appendonly no`，将`no`改为`yes`开启AOF

#### go-redis
Go-redis是Golang中用于与Redis交互的强大工具，支持Redis Server的Golang客户端。

GitHub：<https://github.com/redis/go-redis>  
文档：<https://redis.uptrace.dev/zh>

```
go get -u github.com/go-redis/redis
```
```
文件树
├── config/
│   ├── config.go
│   ├── config.yml
│   ├── db.go
│   └── redis.go <- 新增
├── controllers/
│   ├── article_controller.go
│   ├── auth_controller.go
│   └── exchange_rate_controller.go
├── global/
│   └── global.go <- 修改
├── middlewares/
│   └── auth_middleware.go
├── models/
│   ├── article.go
│   ├── exchange_rate.go
│   └── user.go
├── router/
│   └── router.go
├── utils/
│   └── utils.go
└── main.go
```
``` go
// redis.go
package config

import (
	"log"

	"github.com/go-redis/redis"

	"exchangeapp/global"
)

func initRedis() {
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		DB:       0,
		Password: "",
	})
	_, err := RedisClient.Ping().Result()
	if err != nil {
		log.Fatalf("Failed to connect to redis: %v", err)
	}
	global.RedisDB = RedisClient
}
```
``` go
// global.go
package global

import (
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

var (
	Db      *gorm.DB
	RedisDB *redis.Client
)
```
### 5.2 Redis实战&实现点赞功能&原子操作概述
#### redis key 命名规范的设计
key单词与单词之间以`:`分割：如`user:userinfo`、`article:1:likes`

#### SET & INCR & DECR & GET 命令
1. `SET`命令用于设置给定key的值
2. `INCR`将key中储存的数字值增一  
若key不存在，那么key的值会先被初始化为0，然后再执行操作：`0`->`1`
3. `DECR`则将key中储存的数字减一
4. `GET`命令用于获取指定key的值

#### 重要概念：原子性
原子性确保一个操作在执行过程中是不可被打断的。对于原子操作，要么所有的步骤都成功完成并对外可见，要么这些步骤都不执行，系统的状态保持不变。

`INCR`和`DECR`操作再Redis中就是原子操作(原子性的)，它们要么完成对一个键值的增加/减小，要么完全不进行，确保没有中间状态。

例子：用户A发起点赞请求，Redis执行`INCR`命令，将点赞数从`10`增加到`11`，用户B几乎同时发起点赞请求，Redis执行`INCR`命令，将点赞数从`11`增加到`12`。

因此，若多个操作只是单纯的对数据进行增加值或减小值，Redis提供的`INCR`和`DECR`命令可以直接帮助我们进行并发控制。

但是，若多个操作不只是单纯的进行数据增减值，还包括更复杂的操作，如：逻辑判断，此时Redis的单命令操作无法保证多个操作互斥执行，可用Lua脚本来解决此问题。

```
文件树
├── config/
│   ├── config.go
│   ├── config.yml
│   ├── db.go
│   └── redis.go
├── controllers/
│   ├── article_controller.go
│   ├── auth_controller.go
│   ├── exchange_rate_controller.go
│   └── like_controller.go <- 新增
├── global/
│   └── global.go
├── middlewares/
│   └── auth_middleware.go
├── models/
│   ├── article.go
│   ├── exchange_rate.go
│   └── user.go
├── router/
│   └── router.go <- 修改
├── utils/
│   └── utils.go
└── main.go
```
``` go
// like_controller.go
package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"

	"exchangeapp/global"
)

func LikeArticle(ctx *gin.Context) {
	articleID := ctx.Param("id")
	likeKey := "article:" + articleID + ":likes"
	if err := global.RedisDB.Incr(likeKey).Err(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully liked the article"})
}

func GetArticleLikes(ctx *gin.Context) {
	articleID := ctx.Param("id")
	likeKey := "article:" + articleID + ":likes"
	likes, err := global.RedisDB.Get(likeKey).Result()
	if errors.Is(err, redis.Nil) {
		likes = "0"
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"likes": likes})
}
```
``` go
// router.go
package router

import (
	"github.com/gin-gonic/gin"

	"exchangeapp/controllers"
	"exchangeapp/middlewares"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	auth := r.Group("/api/auth")
	{
		auth.POST("/login", controllers.Login)
		auth.POST("/register", controllers.Register)
	}

	api := r.Group("/api")
	api.GET("/exchangeRates", controllers.GetExchangeRates)
	api.Use(middlewares.AuthMiddleWare())
	{
		api.POST("/exchangeRates", controllers.CreateExchangeRate)

		api.POST("/articles", controllers.CreateArticle)
		api.GET("/articles", controllers.GetArticles)
		api.GET("/articles/:id", controllers.GetArticlesByID)

		api.POST("/articles/:id/like", controllers.LikeArticle)
		api.GET("articles/:id/like", controllers.GetArticleLikes)
	}

	return r
}
```
### 5.3 Redis实战&缓存&旁路缓存模式
#### 缓存
为什么要有缓存技术？——减轻数据库访问压力，加快请求响应速度：缓存读取速度比从数据库中读取快得多，亦可大幅减少数据库查询次数，提高应用程序的响应速度。

##### 旁路缓存模式(Cache-Aside模式)
应用程序直接与缓存和数据库交互，并负责管理缓存的内容。使用该模式的应用程序会同时操作缓存与数据库，具体流程如下：

先尝试从缓存中读取数据，若缓存命中，直接返回缓存中的数据。  
若缓存未命中，则从数据库中读取数据，并将数据存入缓存中，然后返回给客户端。

代码逻辑：

1. 缓存未命中：
  * 如果Redis中没有找到对应的缓存，代码会从数据库中获取文章数据。
  * 获取到数据后，代码将数据序列化为JSON并存储在Redis缓存中，同时设置一个过期时间。
  * 最后，返回数据库中的数据给客户端。
2. 缓存命中：
  * 如果缓存命中(Redis中找到对应的缓存数据)，代码直接从缓存中获取数据。
  * 然后将缓存中的数据反序列化为文章列表，返回给客户端。

解决下一个问题

若在缓存有效期内，用户又新增了一些文章，此时用户通过缓存得到文章，将看不到变化。

解决方法案例：常见的缓存失效策略

1. 设置过期时间(已经使用过)
2. 主动更新缓存
  * 当新增文章文章时，除了更新数据库，还要同步更新或删除缓存中对应的数据。这样，下一次读取时，缓存中的数据就是最新的。
  * 或者新增文章时，不直接更新缓存，而是删除缓存中的旧数据。下次读取时，由于缓存没有命中，从数据库中读取最新数据并重新写入缓存。Redis`DEL`命令：用于删除已存在的键。

```
文件树
├── config/
│   ├── config.go
│   ├── config.yml
│   ├── db.go
│   └── redis.go
├── controllers/
│   ├── article_controller.go <- 修改
│   ├── auth_controller.go
│   ├── exchange_rate_controller.go
│   └── like_controller.go
├── global/
│   └── global.go
├── middlewares/
│   └── auth_middleware.go
├── models/
│   ├── article.go
│   ├── exchange_rate.go
│   └── user.go
├── router/
│   └── router.go
├── utils/
│   └── utils.go
└── main.go
```
``` go
// article_controller.go
package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"gorm.io/gorm"

	"exchangeapp/global"
	"exchangeapp/models"
)

var cacheKey = "articles"

func CreateArticle(ctx *gin.Context) {
	var article models.Article
	if err := ctx.ShouldBindJSON(&article); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := global.Db.AutoMigrate(&article); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := global.Db.Create(&article).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 新增文章时，删除缓存中的旧数据
	if err := global.RedisDB.Del(cacheKey).Err(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, article)
}

func GetArticles(ctx *gin.Context) {
	cachedData, err := global.RedisDB.Get(cacheKey).Result() // 尝试从缓存中获取数据
	if errors.Is(err, redis.Nil) { // 1. 缓存未命中
		var articles []models.Article
		// 1.1 从数据库中获取文章数据
		if err := global.Db.Find(&articles).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			return
		}
		// 1.2 将数据序列化为JSON并存储在Redis缓存中
		articleJSON, err := json.Marshal(articles)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if err := global.RedisDB.Set(cacheKey, articleJSON, 10*time.Minute).Err(); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		// 1.3 返回数据库中的数据给客户端
		ctx.JSON(http.StatusOK, articles)
	} else if err != nil { // 2. 其他意外错误
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else { // 3. 缓存命中
		var articles []models.Article
		if err := json.Unmarshal([]byte(cachedData), &articles); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, articles)
	}
}

func GetArticlesByID(ctx *gin.Context) {
	id := ctx.Param("id")

	var article models.Article

	if err := global.Db.Where("id = ?", id).First(&article).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, article)
}
```
## Section6 解决其他问题
涉及前端内容

To be continued ...
