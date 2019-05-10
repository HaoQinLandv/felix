package cmd

import (
	"github.com/dejavuzhou/felix/ginbro"
	"github.com/spf13/cobra"
	"path/filepath"
)

var appListen, dir, authTable, authColumn, dbUser, dbPassword, dbAddr, dbName, dbCharset, dbType, packageName string

// restCmd represents the rest command
var restCmd = &cobra.Command{
	Use:     "ginbro",
	Short:   "generate a RESTful code project from SQL database",
	Long:    `generate a RESTful APIs app with gin and gorm for gophers`,
	Example: `felix rest -u root -p password -a "127.0.0.1:3306" -d dbname -c utf8 --authTable=users --authColumn=pw_column -o=FelixRestOut"`,

	Run: func(cmd *cobra.Command, args []string) {
		appDir := filepath.Clean(filepath.Join(dir, packageName))
		err := ginbro.Run(dbUser, dbPassword, dbAddr, dbName, dbCharset, dbType, appDir, appListen, authTable, authColumn, packageName)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(restCmd)

	restCmd.Flags().StringVarP(&appListen, "appListen", "l", "127.0.0.1:5555", "app's listening addr")
	restCmd.Flags().StringVarP(&dir, "dir", "d", ".", "code project output directory")
	restCmd.Flags().StringVarP(&packageName, "pkg", "P", "", "eg1: github.com/dejavuzhou/ginSon, eg2: ginbroSon")

	restCmd.Flags().StringVar(&authTable, "authTable", "users", "login user table")
	restCmd.Flags().StringVar(&authColumn, "authColumn", "password", "bcrypt password column")

	restCmd.Flags().StringVarP(&dbUser, "dbUser", "u", "root", "database username")
	restCmd.Flags().StringVarP(&dbPassword, "dbPassword", "p", "password", "database user password")
	restCmd.Flags().StringVarP(&dbAddr, "dbAddr", "a", "127.0.0.1:3306", "datatbase connection addr")
	restCmd.Flags().StringVarP(&dbName, "dbName", "n", "", "database name")
	restCmd.Flags().StringVarP(&dbCharset, "dbCharset", "c", "utf8", "database charset")
	restCmd.Flags().StringVarP(&dbType, "dbType", "t", "mysql", "database type: mysql/postgres/mssql/sqlite")

	restCmd.MarkFlagRequired("package")
	restCmd.MarkFlagRequired("dbAddr")
}
