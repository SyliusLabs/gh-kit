package main

import "github.com/jakubtobiasz/gh-kit/cmd"

func main() {
	err := cmd.Run()

	if nil != err {
		panic(err)
	}
}
