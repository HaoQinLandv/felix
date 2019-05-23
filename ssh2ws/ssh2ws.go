package ssh2ws

import (
	"github.com/dejavuzhou/felix/models"
	"github.com/dejavuzhou/felix/ssh2ws/internal"
	"github.com/dejavuzhou/felix/staticbin"
	"github.com/gin-gonic/gin"
	"time"
)

func RunSsh2ws(bindAddress, user, password, secret string, expire time.Duration) error {

	//config jwt variables
	models.AppSecret = secret
	models.ExpireTime = expire
	models.AppIss = "felix.mojotv.cn"
	models.DefaultUser = user
	models.DefaultPassword = password

	//gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.MaxMultipartMemory = 32 << 20

	//sever static file in http's root path
	binStaticMiddleware, err := staticbin.NewGinStaticBinMiddleware("/")
	if err != nil {
		return err
	}
	r.Use(binStaticMiddleware)

	api := r.Group("api")
	r.POST("api/login", internal.Login)
	r.GET("dlg", internal.GinbroDownload)
	api.Use(internal.JwtMiddleware)
	{
		api.GET("ws/:id", internal.WsSsh)

		api.GET("ssh", internal.SshAll)
		api.POST("ssh", internal.SshCreate)
		api.GET("ssh/:id", internal.SshOne)
		api.PATCH("ssh/:id", internal.SshUpdate)
		api.DELETE("ssh/:id", internal.SshDelete)

		api.GET("sftp/:id", internal.SftpLs)
		api.GET("sftp/:id/dl", internal.SftpDl)
		api.GET("sftp/:id/cat", internal.SftpCat)
		api.GET("sftp/:id/rm", internal.SftpRm)
		api.GET("sftp/:id/rename", internal.SftpRename)
		api.GET("sftp/:id/mkdir", internal.SftpMkdir)
		api.POST("sftp/:id/up", internal.SftpUp)

		api.POST("ginbro/gen", internal.GinbroGen)
		api.POST("ginbro/db", internal.GinbroDb)
		api.GET("ginbro/dl", internal.GinbroDownload)
	}

	if err := r.Run(bindAddress); err != nil {
		return err
	}
	return nil
}
