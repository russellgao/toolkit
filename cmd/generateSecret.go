package cmd

import (
	"github.com/russellgao/toolkit/internal"
	"github.com/spf13/cobra"
)

/*
 tkctl secret -h
生成随机密码，支持1～100位长度，可以指定是否包含特殊字符

Usage:
  tkctl secret [flags]

Flags:
  -c, --char string   是否包含特殊字符，默认为包含所有的特殊字符。如果不包含请传值为false，此时随机密码由数字和大小写字母组成；如果需要自定义特殊请通过此参数直接传入即可，例:只包含 #*& 的特殊字符，此参数应该为 --special_char #*&
  -h, --help          help for secret
  -n, --length int    随机密码的长度 (default 16)

*/
func newGenerateSecterCode() *cobra.Command {
	generateSecretCodeCmd := &cobra.Command{
		Use:   "secret",
		Short: "生成随机密码，支持1～100位长度，可以指定是否包含特殊字符",
		Run: func(cmd *cobra.Command, args []string) {
			internal.GenerateSecretCode(specialCharacter, length)
		},
	}
	special_char_help := "是否包含特殊字符，默认为包含所有的特殊字符。如果不包含请传值为false，此时随机密码由数字和大小写字母组成；如果需要自定义特殊请通过此参数直接传入即可，例:" +
		"只包含 #*& 的特殊字符，此参数应该为 --special_char #*&"
	generateSecretCodeCmd.Flags().StringVarP(&specialCharacter, "char", "c", "", special_char_help)
	generateSecretCodeCmd.Flags().IntVarP(&length, "length", "n", 16, "随机密码的长度")
	return generateSecretCodeCmd
}
