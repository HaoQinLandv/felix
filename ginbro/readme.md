# [Converting a MySQL database'schema to a RESTful golang APIs app in the fastest way](https://github.com/dejavuzhou/felix/ginbro)


Ginbro is a scaffold tool for Gin-Gorm-MySQL which you just need to input one command to create a mighty RESTful APIs App.

## Demo and Translated Document
- [中文文档](readme_zh.md)            
- [Video-Demo-Youtube](https://www.youtube.com/watch?v=TvWQhNKfmCo&feature=youtu.be)
- [Video-Demo-Bilibili](https://www.bilibili.com/video/av36804258/)

## Feature
- [fastest way to generate a RESTful APIs application with MySQL in Go](/boilerplate)
- support [JWT Authorization Bearer](_boilerplate/handlers/middleware_jwt.go) [Auth](_boilerplate/handlers/handler_auth.go) and [JWT middleware](_boilerplate/models/jwt.go)
- [support brute-force-login firewall](_boilerplate/models/model_users.go)
- [build in swift golang-memory cache](https://github.com/dejavuzhou/ginbro/blob/master/boilerplate/models/db_memory.go)
- [generate GORM model from MySQL database schema](_boilerplate/models)
- [powered with Swagger document and SwaggerUI](_boilerplate/static/swagger)
- [capable of serve VueJs app's static files](_boilerplate/static)
- [configurable CORS middleware](_boilerplate/handlers/gin_helper.go)
- [user friendly configuration](tpl/config.toml)
- [fully build-in cron task support](_boilerplate/tasks)
- [travis CI/CD](https://travis-ci.org/dejavuzhou/felix)
    


## Usage

### 1. `felix ginbro` generate a new Gin+Gorm+MySQL RESTful APIs Application with JWT middleware and auth
example 

`felix ginbro -a dev.wordpress.com:3306 -P go_package_name -n db_name -u db_username -p 'my_db_password' -d '~/thisDir' `

```bash
$ felix ginbro -h
generate a RESTful APIs app with gin and gorm for gophers

Usage:
  felix ginbro [flags]

Examples:
felix ginbro -u root -p password -a "127.0.0.1:3306" -d dbname -c utf8 --authTable=users --authColumn=pw_column -P=APP_PACKAGE_NAME" -d '~/thisDir' 

Flags:
  -l, --appListen string    app's listening addr (default "127.0.0.1:5555")
      --authColumn string   bcrypt password column (default "password")
      --authTable string    login user table (default "users")
  -a, --dbAddr string       database connection addr (default "127.0.0.1:3306")
  -c, --dbCharset string    database charset (default "utf8")
  -n, --dbName string       database name
  -p, --dbPassword string   database user password (default "password")
  -t, --dbType string       database type: mysql/postgres/mssql/sqlite (default"mysql")
  -u, --dbUser string       database username (default "root")
  -d, --dir string          code project output directory,default is current working dir (default ".")
  -h, --help                help for ginbro
  -P, --pkg string          eg1: github.com/dejavuzhou/ginSon, eg2: ginbroSon

Global Flags:
      --verbose   verbose
```

#### the generated project directory [ginbro DEMO-code-repository](https://github.com/dejavuzhou/ginbro-son)



## environment
- my development environment
    - Windows 10 pro 64
    - go version go1.11.1 windows/amd64
    - mysql version <= 5.7

## Info
- resource table'schema which has no "ID","id","Id'" or "iD" will not generate model or route.
- the column which type is json value must be a string which is able to decode into a JSON, when resource is called POST or PATCH from the swaggerUI.
## Thanks
- [swagger Specification](https://swagger.io/specification/)
- [gin-gonic/gin](https://github.com/gin-gonic/gin)
- [GORM](http://gorm.io/)
- [viper](https://github.com/spf13/viper)
- [cobra](https://github.com/spf13/cobra#getting-started)
- [dgrijalva/jwt-go](https://github.com/dgrijalva/jwt-go)
- [base64captcha](https://github.com/mojocn/base64Captcha)
## Please feedback your [`issue`](https://github.com/dejavuzhou/felix/issues) with database schema file