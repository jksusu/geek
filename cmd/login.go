package cmd

import (
	"fmt"
	"geek/app/login"

	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "登录账号",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Print("请输入账号密码：geek login username=  password=")
			return
		}
		loginCli()
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}

func loginCli() {
	login.GeekLogin()
}
