package main

import (
	"fmt"
	"github.com/russellgao/toolkit/cmd"
	"os"
)

func init() {

}

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
