package ginbro

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"os/exec"
	"path"
	"path/filepath"
)

type app struct {
	ProjectPath    string
	ProjectPackage string
	AppSecret      string
	AppListen      string
	Resources      []Resource
	AuthTable      string
	AuthPassword   string
	DbType         string
	DbAddr         string
	DbName         string
	DbUser         string
	DbPassword     string
	DbCharset      string
}

func Run(dbUser, dbPassword, dbAddr, dbName, dbCharset, dbType, appPackage, appListen, authTable, authColumn string) error {
	app, err := newApp(dbType, dbAddr, dbUser, dbPassword, dbName, dbCharset, authTable, authColumn, appPackage, appListen)
	if err != nil {
		return err
	}
	err = app.makeProjectDir()
	if err != nil {
		return err
	}
	err = app.copyStaticAndSwagger()
	if err != nil {
		return err
	}
	err = app.copyTimerTaskDir()
	if err != nil {
		return err
	}
	err = app.copyConfigPackage()
	if err != nil {
		return err
	}
	err = app.generateCodeBase()
	if err != nil {
		return err
	}
	//go fmt codebase
	err = app.goFmtCodeBase()
	if err != nil {
		return err
	}
	return nil
}

func (app *app) makeProjectDir() error {
	if err := os.RemoveAll(app.ProjectPath); err != nil {
		return err
	}
	for _, name := range []string{"handlers", "models"} {
		modulePath := path.Join(app.ProjectPath, name)
		modulePath = filepath.Clean(modulePath)
		err := os.MkdirAll(modulePath, 0777)
		if err != nil {
			return err
		}
	}
	return nil
}
func (app *app) copyStaticAndSwagger() error {
	srcStatic := filepath.Join(goPath, "src", felixGinbroPackage, "_boilerplate/static")
	dstStatic := filepath.Join(app.ProjectPath, "static")
	srcSwagger := filepath.Join(goPath, "src", felixGinbroPackage, "_boilerplate/swagger")
	dstSwagger := filepath.Join(app.ProjectPath, "swagger")
	err := CopyDir(srcSwagger, dstSwagger)
	if err != nil {
		return err
	}
	return CopyDir(srcStatic, dstStatic)
}
func (app *app) copyTimerTaskDir() error {
	src := filepath.Join(goPath, "src", felixGinbroPackage, "_boilerplate/tasks")
	dst := filepath.Join(app.ProjectPath, "tasks")
	return CopyDir(src, dst)
}
func (app *app) copyConfigPackage() error {
	src := filepath.Join(goPath, "src", felixGinbroPackage, "_boilerplate/config")
	dst := filepath.Join(app.ProjectPath, "config")
	return CopyDir(src, dst)
}
func newApp(dbType, dbAddr, dbUser, dbPassword, dbName, dbCharset, authTable, authColumn, projectPackage, appListen string) (*app, error) {
	cols, err := FetchDbColumn(dbType, dbAddr, dbUser, dbPassword, dbName, dbCharset)
	if err != nil {
		return nil, err
	}
	resources, err := transformToResources(cols, authTable, authColumn)
	if err != nil {
		return nil, err
	}
	return &app{
		ProjectPackage: projectPackage,
		AppListen:      appListen,
		AuthTable:      authTable,
		AuthPassword:   authColumn,
		AppSecret:      randomString(32),
		Resources:      resources,
		ProjectPath:    filepath.Join(goPath, "src", projectPackage),
		DbType:         dbType,
		DbAddr:         dbAddr,
		DbName:         dbName,
		DbUser:         dbUser,
		DbCharset:      dbCharset,
		DbPassword:     dbPassword,
	}, nil
}
func (app *app) generateCodeBase() error {
	jobs := map[string]string{
		"tpl/single/models.db.go.tpl":               "models/db.go", //TODO::数据库三种连接
		"tpl/single/models.db_memory.go.tpl":        "models/db_memory.go",
		"tpl/single/models.db_helper.go.tpl":        "models/db_helper.go",
		"tpl/single/main.go.tpl":                    "main.go",
		"tpl/swagger.yaml":                          "swagger/doc.yaml",
		"tpl/single/handlers.gin.go.tpl":            "handlers/gin.go",
		"tpl/single/handlers.gin_helper.go.tpl":     "handlers/gin_helper.go",
		"tpl/single/handlers.middleware_jwt.go.tpl": "handlers/middleware_jwt.go",

		"tpl/config.toml": "config.toml",
	}
	for source, destination := range jobs {
		err := parseTemplate(source, app.ProjectPackage, destination, app)
		if err != nil {
			return fmt.Errorf("parse [%s] template into [%s] failed with error : %s", source, destination, err)
		}
	}

	for _, resouce := range app.Resources {
		resouce.ProjectPackage = app.ProjectPackage
		uri := fmt.Sprintf("models/model_%s.go", resouce.TableName)
		//generate model from resource
		err := parseTemplate("tpl/models.template.tpl", app.ProjectPackage, uri, resouce)
		if err != nil {
			return err
		}
		uri = fmt.Sprintf("handlers/handler_%s.go", resouce.TableName)
		err = parseTemplate("tpl/handlers.template.tpl", app.ProjectPackage, uri, resouce)
		if err != nil {
			return err
		}
		if resouce.IsAuthTable {
			err = parseTemplate("tpl/models.jwt.go.tpl", app.ProjectPackage, "models/jwt.go", resouce)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (app *app) goFmtCodeBase() error {
	cmd := exec.Command("go", "fmt", app.ProjectPackage+"/...")
	_, err := cmd.Output()
	if err != nil {
		return err
	}
	color.Red("your code base is at %s", app.ProjectPath)
	return nil
}
