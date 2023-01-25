package main

import (
	"fmt"
	"github.com/SyliusLabs/gh-kit/cmd"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Print(r)
		}
	}()

	err := cmd.Run()

	if nil != err {
		panic(err)
	}
}
