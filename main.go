package main

import "os"

// Name is the name of this application.
const Name = "git-pull-requests"

func main() {
	cli := &CLI{}
	code := cli.Run([]string{})
	os.Exit(code)
}
