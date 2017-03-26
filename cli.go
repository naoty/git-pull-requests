package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

// CLI represents the CLI for this application.
type CLI struct {
}

// Run runs a command with given args.
func (cli *CLI) Run(args []string) int {
	token, err := readToken()

	if err == nil {
		fmt.Println(token)
	} else {
		return 1
	}

	return 0
}

func readToken() (string, error) {
	path := configFilePath()
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	config := Config{}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return "", err
	}

	return config.Token, nil
}

func configFilePath() string {
	dir := os.Getenv("HOME")
	path := filepath.Join(dir, ".config", Name)
	return path
}
