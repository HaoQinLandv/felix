# a web terminal to ssh

## [Vuejs Admin Demo UserName:`admin` Password:`admin`](http://home.mojotv.cn:2222)

## command: `felix sshw`
```bash
$ felix sshw -h
the demo website is http://home.mojotv.cn:2222

Usage:
  felix sshw [flags]

Flags:
  -a, --addr string       listening addr (default ":2222")
  -x, --expire uint       token expire in * minute (default 1440)
  -h, --help              help for sshw
  -p, --password string   auth password (default "admin")
  -s, --secret string     jwt secret string
  -u, --user string       auth user (default "admin")

Global Flags:
      --verbose   verbose
```

```bash
$ felix sshw
use random string as jwt secret: @Ubr)Vrp~Zoo6Rvrk1PP1*ZXPYby_Z)s
login user: admin
login password: admin
login expire in 1440 minutes
[GIN-debug] [WARNING] Now Gin requires Go 1.6 or later and Go 1.7 will be required soon.

[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "_release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] POST   /api/login                --> github.com/dejavuzhou/felix/ssh2ws/controllers.GetLoginHandler.func1 (4 handlers)
[GIN-debug] GET    /api/ws/:id               --> github.com/dejavuzhou/felix/ssh2ws/controllers.WsSsh (5 handlers)
[GIN-debug] GET    /api/ssh                  --> github.com/dejavuzhou/felix/ssh2ws/controllers.SshAll (5 handlers)
[GIN-debug] POST   /api/ssh                  --> github.com/dejavuzhou/felix/ssh2ws/controllers.SshCreate (5 handlers)
[GIN-debug] GET    /api/ssh/:id              --> github.com/dejavuzhou/felix/ssh2ws/controllers.SshOne (5 handlers)
[GIN-debug] PATCH  /api/ssh/:id              --> github.com/dejavuzhou/felix/ssh2ws/controllers.SshUpdate (5 handlers)
[GIN-debug] DELETE /api/ssh/:id              --> github.com/dejavuzhou/felix/ssh2ws/controllers.SshDelete (5 handlers)
[GIN-debug] GET    /api/sftp/:id             --> github.com/dejavuzhou/felix/ssh2ws/controllers.SftpLs (5 handlers)
[GIN-debug] GET    /api/sftp/:id/dl          --> github.com/dejavuzhou/felix/ssh2ws/controllers.SftpDl (5 handlers)
[GIN-debug] GET    /api/sftp/:id/cat         --> github.com/dejavuzhou/felix/ssh2ws/controllers.SftpCat (5 handlers)
[GIN-debug] GET    /api/sftp/:id/rm          --> github.com/dejavuzhou/felix/ssh2ws/controllers.SftpRm (5 handlers)
[GIN-debug] GET    /api/sftp/:id/rename      --> github.com/dejavuzhou/felix/ssh2ws/controllers.SftpRename (5 handlers)
[GIN-debug] GET    /api/sftp/:id/mkdir       --> github.com/dejavuzhou/felix/ssh2ws/controllers.SftpMkdir (5 handlers)
[GIN-debug] POST   /api/sftp/:id/up          --> github.com/dejavuzhou/felix/ssh2ws/controllers.SftpUp (5 handlers)
[GIN-debug] Listening and serving HTTP on :2222

```
![](/images/sshw2.png)

![](/images/sshw3.png)

![](/images/sshw4.png)

![](/images/sshw1.png)