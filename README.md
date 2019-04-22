# Felix
[![Build Status](https://travis-ci.org/dejavuzhou/felix.svg?branch=master)](https://travis-ci.org/dejavuzhou/felix)

## For Who

- BackEnd Engineer
- DevOps Engineer
- People Heavily engage with SSH

## Do What

- manage massive SSH login configuration
- generate a RESTful app from SQL database with [gin-gonic/gin](https://github.com/gin-gonic/gin) and [GORM](https://github.com/jinzhu/gorm) in GO
- start TCP and SOCK proxy with ssh with speedily
- terminal task list
- Pewdiepie's brofit command to subscribe the Youtube channel

## Overview
commands:

```bash
$ felix

Usage:
  felix [command]

Available Commands:
  brofist     Pewdiepie needs your help.Do your part to subscribe Pewdiepie's Yo                                                                                                              utube Channel.
  clean       purge all felix configuration
  ginbro      generate a RESTful code project from SQL database
  goDoc       golang.google.cn/pkg
  help        Help about any command
  json        open a tab in browser to convert json to golang struct
  ssh         open a ssh terminal
  sshadd      add a ssh connection
  sshdl       scp download file or folder
  sshdu       duplicate a ssh connection
  sshedit     update a ssh connection config
  sshexport   export all ssh connection info to a csv file
  sshimport   import massive ssh server info from a csv file
  sshinfo     view a ssh info
  sshls       list all ssh info or search by hostname
  sshproxy    ssh port proxy
  sshrm       delete a ssh info
  sshsocks    start a socks4/5 proxy
  sshup       scp upload
  task        show all TODO tasks
  taskad      add a task
  taskok      set a task done
  taskrm      remove a task

Flags:
  -h, --help      help for felix
      --verbose   verbose

Use "felix [command] --help" for more information about a command.

```


## Installation Or Compiling

### 

## ScreenShot

### command: `felix ginbro`
```bash
$ felix ginbro
Error: required flag(s) "appDir", "dbAddr" not set
Usage:
  felix ginbro [flags]

Examples:
felix rest -u root -p password -a "127.0.0.1:3306" -d dbname -c utf8 --authTable=users --authColumn=pw_column -o=FelixRestOut"

Flags:
  -o, --appDir string       app's code output directory
  -l, --appListen string    app's listening addr (default "127.0.0.1:5555")
      --authColumn string   bcrypt password column (default "password")
      --authTable string    login user table (default "users")
  -a, --dbAddr string       datatbase connection addr (default "127.0.0.1:3306")
  -c, --dbCharset string    database charset (default "utf8")
  -n, --dbName string       database name
  -p, --dbPassword string   database user password (default "password")
  -t, --dbType string       database type: mysql/postgres/mssql/sqlite (default "mysql")
  -u, --dbUser string       database username (default "root")
  -h, --help                help for ginbro

Global Flags:
      --verbose   verbose

required flag(s) "appDir", "dbAddr" not set
```

### command: `felix sshls`

![felix sshls](iamges/sshls)

### command: `felix ssh 2`

![felix sshls](iamges/sshIn)

### command: `felix sshsocks 34 -l 1080`

![felix sshls](iamges/sshsocks)

### command: `felix taskad`

![felix sshls](iamges/taskad)
