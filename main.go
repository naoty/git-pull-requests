package main

import "os"

func main() {
	cli := &CLI{}
	code := cli.Run([]string{})
	os.Exit(code)
}
