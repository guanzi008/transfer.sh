package main

import "github.com/guanzi008/transfer.sh/cmd"

func main() {
	app := cmd.New()
	app.RunAndExitOnError()
}
