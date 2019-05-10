# [MySQL数据库生成RESTful APIs APP](https://github.com/dejavuzhou/felix)
[![Build Status](https://travis-ci.org/dejavuzhou/felix.svg?branch=master)](https://travis-ci.org/dejavuzhou/felix) 
[![GoDoc](http://godoc.org/github.com/dejavuzhou/felix?status.svg)](http://godoc.org/github.com/dejavuzhou/felix) 
[![Go Report Card](https://goreportcard.com/badge/github.com/dejavuzhou/felix)](https://goreportcard.com/report/github.com/dejavuzhou/felix)
![stability-stable](https://img.shields.io/badge/stability-stable-brightgreen.svg)
[![codebeat badge](https://codebeat.co/badges/650029a5-fcea-4416-925e-277e2f178e96)](https://codebeat.co/projects/github-com-dejavuzhou-ginbro-master)
[![codecov](https://codecov.io/gh/dejavuzhou/felix/branch/master/graph/badge.svg)](https://codecov.io/gh/dejavuzhou/felix)

## 一个命令行工具:快速生成go语言RESTful APIs应用 
## 文档和DEMO
- [在线DEMO](http://ginbro.mojotv.cn/swagger/)
- [中文文档](readme_zh.md)            
- [Video-Demo-Youtube](https://www.youtube.com/watch?v=TvWQhNKfmCo&feature=youtu.be)
- [Video-Demo-Bilibili](https://www.bilibili.com/video/av36804258/)

## Feature
- [生成完善RESTful APIs 应用](/boilerplate)
- [自动生成完善的Swagger文档](_boilerplate/static/swagger)
- [自动生成数据库表的模型和标注](_boilerplate/models)
- 支持 [JWT Authorization Bearer](_boilerplate/handlers/middleware_jwt.go) [身份验证](_boilerplate/handlers/handler_auth.go) and [JWT 中间件](_boilerplate/models/jwt.go)
- [支持登陆防火墙](_boilerplate/models/model_users.go)
- [支持静态资源替代nginx](_boilerplate/static)
- [可配置的跨域cors中间件](_boilerplate/handlers/gin_helper.go)
- [用户友好的自定义配置](tpl/config.toml)
- [支持定时任务](_boilerplate/tasks)
- [内置高效率的内存数据库](https://github.com/dejavuzhou/felix/blob/master/boilerplate/models/db_memory.go)


## 使用
`felix ginbro -u root -p ginbro -c utf8 -a 127.0.0.1:3306 --authTable=users --authPassword=password  --outPackage=github.com/dejavuzhou/felix_demo -d=ginbro`
- cd 到生成的项目
- go build  和run
- 访问[`http://127.0.0.1:5555/swagger`](http://127.0.0.1:5555/swagger)

### 生成新project目录树 [ginbro-son DEMO代码](https://github.com/dejavuzhou/felix-son)
```shell
C:\Users\zhouqing1\go\src\github.com\mojocn\apiapp>tree /f /a
|   config.toml
|   main.go
|   readme.md
|
+---config
|       viper.go
+---handlers
|       gin.go
|       handler_wp_posts.go
|       handler_wp_users.go
|
+---models
|       db.go
|       model_wp_posts.go
|       model_wp_users.go
|
+---static
|   |   .gitignore
|   |   index.html
|   |   readme.md
|   |
|   \---index_files
|           style.css
|           syntax.css
|
\---swagger
        .gitignore
        doc.yml
        favicon-16x16.png
        favicon-32x32.png
        index.html
        oauth2-redirect.html
        readme.md

```
### 命令参数说明
```shell
felix ginbro -h
generate a RESTful APIs app with gin and gorm for gophers. For example:
        felix ginbro -u eric -p password -a "127.0.0.1:3306" -d "mydb"

Usage:
  create gen [flags]

Flags:
  -a, --Mysql IP PORT    mysql host:port (default "dev.mojotv.com:3306")
  -l, --应用地址端口    app listen Address eg:mojotv.cn, use domain will support gin-TLS (default "127.0.0.1:5555")
  -c, --数据库字符集    database charset (default "utf8")
  -d, --数据库名称   database name (default "dbname")
  -h, --help              help for gen
  -o, --输出地址      输出地址相对于$GOPATH/src
  -p, --数据库密码   database password (default "Password")
  -u, --数据库用户     database user name (default "root")
  --authTable 登陆用户表名  default users
  --authPassword 登陆用户密码字段 default password
```


## 注意
- mysql表中没有id/ID/Id/iD字段将不会生成路由和模型
- json字段 在update/create的时候 必须使可以序列号的json字符串(`eg0:"{}" eg1:"[]"`),否则mysql会报错

## 致谢
- [swagger规范](https://swagger.io/specification/)
- [gin-gonic/gin框架](https://github.com/gin-gonic/gin)
- [GORM数据库ORM](http://gorm.io/)
- [viper配置文件读取](https://github.com/spf13/viper)
- [cobra命令行工具](https://github.com/spf13/cobra#getting-started)
- [我的另外一个go图像验证码开源项目](https://github.com/mojocn/base64Captcha)

## 请各位大神不要吝惜提[`issue`](https://github.com/dejavuzhou/felix/issues)同时附上数据库表结构文件