package cmd

import (
	"github.com/russellgao/toolkit/util"
	"github.com/spf13/cobra"
	"log"
	"regexp"
	"strings"
)

/*
 tkctl replace -h
文本替换，支持正则替换和非正则替换，类似与linux下的sed，但比sed更好用，而且可以跨平台使用

Usage:
  tkctl replace [flags]

Flags:
  -d, --dirs string      需要替换的目录, 默认为当前路径 (default ".")
  -h, --help             help for replace
  -m, --mode string      替换的模式，支持正则（regexp）和非正则（text）两种模式，默认非正则， (default "text")
  -p, --pattern string   需要替换的pattern  [required]
  -r, --repl string      目标字符串  [required]
*/

var (
	text_mode    string
	text_repl    string
	text_pattern string
	text_dirs    string
)

func NewTextReplaceCmd() *cobra.Command {
	textReplaceCmd := &cobra.Command{
		Use:   "replace",
		Short: "文本替换，支持正则替换和非正则替换，类似与linux下的sed，但比sed更好用，而且可以跨平台使用",
		Run: func(cmd *cobra.Command, args []string) {
			textReplace()
		},
	}
	textReplaceCmd.Flags().StringVarP(&text_mode, "mode", "m", "text", "替换的模式，支持正则（regexp）和非正则（text）两种模式，默认非正则，")
	textReplaceCmd.Flags().StringVarP(&text_repl, "repl", "r", "", "目标字符串  [required]")
	textReplaceCmd.Flags().StringVarP(&text_pattern, "pattern", "p", "", "需要替换的pattern  [required]")
	textReplaceCmd.Flags().StringVarP(&text_dirs, "dirs", "d", ".", "需要替换的目录, 默认为当前路径")
	return textReplaceCmd
}

// validate input
func validateFlags() {
	mode := map[string]bool{"text": true, "regexp": true}
	if !mode[text_mode] {
		log.Fatal("--text-mode 参数输入有误，只支持text 和 regexp 两种模式")
	}
	if text_pattern == "" {
		log.Fatal("--text-pattern 参数为空，请重新输入执行")
	}
	if text_repl == "" {
		log.Fatal("--text-repl 参数为空，请重新输入执行")
	}
}

// 进行正则替换
func textReplace() {
	validateFlags()
	// 遍历获取文件
	filenames, err := util.GetAllFile(text_dirs, []string{})
	if err != nil {
		log.Fatal(err)
	}
	for _, filename := range filenames {
		content, err := util.ReadFromFile(filename)
		if err != nil {
			log.Fatal(content)
		}
		if text_mode == "text" {
			if strings.Contains(content, text_pattern) {
				log.Println(filename, "中找到了匹配项，开始替换")
				_content := strings.ReplaceAll(content, text_pattern, text_repl)
				err := util.WriteToFile(filename, _content)
				if err != nil {
					log.Fatal(err)
				}
				log.Println(filename, "替换完成")
			}
		} else if text_mode == "regexp" {
			ok, err := regexp.Match(text_pattern, []byte(content))
			if err != nil {
				log.Fatal(err)
			}
			if ok {
				log.Println(filename, "中找到了匹配项，开始替换")
				rep, _ := regexp.Compile(text_pattern)
				_content := rep.ReplaceAllLiteralString(content, text_repl)
				err := util.WriteToFile(filename, _content)
				if err != nil {
					log.Fatal(err)
				}
				log.Println(filename, "替换完成")
			}
		}
	}
}
