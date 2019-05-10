package ginbro

import (
	"fmt"
	"github.com/dejavuzhou/felix/utils"
	"github.com/fatih/color"
	"os/exec"
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

func Run(dbUser, dbPassword, dbAddr, dbName, dbCharset, dbType, appPath, appListen, authTable, authColumn, pkg string) error {
	app, err := newApp(dbType, dbAddr, dbUser, dbPassword, dbName, dbCharset, authTable, authColumn, appPath, appListen, pkg)
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

func newApp(dbType, dbAddr, dbUser, dbPassword, dbName, dbCharset, authTable, authColumn, appPath, appListen, pkgName string) (*app, error) {
	cols, err := FetchDbColumn(dbType, dbAddr, dbUser, dbPassword, dbName, dbCharset)
	if err != nil {
		return nil, err
	}
	resources, err := transformToResources(cols, authTable, authColumn)
	if err != nil {
		return nil, err
	}
	return &app{
		ProjectPackage: pkgName,
		AppListen:      appListen,
		AuthTable:      authTable,
		AuthPassword:   authColumn,
		AppSecret:      utils.RandomString(32),
		Resources:      resources,
		ProjectPath:    appPath,
		DbType:         dbType,
		DbAddr:         dbAddr,
		DbName:         dbName,
		DbUser:         dbUser,
		DbCharset:      dbCharset,
		DbPassword:     dbPassword,
	}, nil
}
func (app *app) generateCodeBase() error {

	for _, tplNode := range parseOneList {
		err := tplNode.ParseExecute(app.ProjectPath, "", app)
		if err != nil {
			return fmt.Errorf("parse [%s] template failed with error : %s", tplNode.NameFormat, err)
		}
	}

	for _, resource := range app.Resources {
		resource.ProjectPackage = app.ProjectPackage
		tableName := resource.TableName
		//generate model from resource
		for _, tplNode := range parseObjList {
			err := tplNode.ParseExecute(app.ProjectPath, tableName, resource)
			if err != nil {
				return fmt.Errorf("parse [%s] template failed with error : %s", tplNode.NameFormat, err)
			}
		}
		if resource.IsAuthTable {
			err := tHandlerLogin.ParseExecute(app.ProjectPath, "", resource)
			if err != nil {
				return fmt.Errorf("parse [%s] template failed with error : %s", tModelJwt.NameFormat, err)
			}
			err = tModelJwt.ParseExecute(app.ProjectPath, "", resource)
			if err != nil {
				return fmt.Errorf("parse [%s] template failed with error : %s", tModelJwt.NameFormat, err)
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
