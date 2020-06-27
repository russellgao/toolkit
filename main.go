package main

import (
	"fmt"
	"os"
	"russellgao/toolkit/cmd"
)

func init() {

}

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
