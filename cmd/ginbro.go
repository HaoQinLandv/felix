package cmd

import (
	"fmt"
	"github.com/dejavuzhou/felix/ginbro"
	"github.com/spf13/cobra"
)

var appListen, appDir, authTable, authColumn, dbUser, dbPassword, dbAddr, dbName, dbCharset, dbType string

// restCmd represents the rest command
var restCmd = &cobra.Command{
	Use:     "ginbro",
	Short:   "根据数据库配置生成RESTfulAPIs APP",
	Long:    `generate a RESTful APIs app with gin and gorm for gophers`,
	Example: `felix rest -u root -p password -a "127.0.0.1:3306" -d dbname -c utf8 --authTable=users --authColumn=pw_column -o=FelixRestOut"`,

	Run: func(cmd *cobra.Command, args []string) {
		err := ginbro.Run(dbUser, dbPassword, dbAddr, dbName, dbCharset, dbType, appDir, appListen, authTable, authColumn)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(restCmd)

	restCmd.Flags().StringVarP(&appListen, "appListen", "l", "127.0.0.1:5555", "APP的端口,可以配置文件中修改")
	restCmd.Flags().StringVarP(&appDir, "appDir", "o", "", "输出项目路径,相对于$GOPATH/src")

	restCmd.Flags().StringVar(&authTable, "authTable", "users", "用户登陆表")
	restCmd.Flags().StringVar(&authColumn, "authColumn", "password", "用户登陆表的bcrypt密码列")

	restCmd.Flags().StringVarP(&dbUser, "dbUser", "u", "root", "数据库用户名")
	restCmd.Flags().StringVarP(&dbPassword, "dbPassword", "p", "password", "数据库密码")
	restCmd.Flags().StringVarP(&dbAddr, "dbAddr", "a", "127.0.0.1:3306", "数据库地址")
	restCmd.Flags().StringVarP(&dbName, "dbName", "n", "", "数据库名称")
	restCmd.Flags().StringVarP(&dbCharset, "dbCharset", "c", "utf8", "数据库字符集")
	restCmd.Flags().StringVarP(&dbType, "dbType", "t", "mysql", "数据库类型mysql/postgres/mssql/sqlite")

	restCmd.MarkFlagRequired("appDir")
	restCmd.MarkFlagRequired("dbAddr")

}
