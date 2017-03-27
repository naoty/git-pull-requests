package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

// CLI represents the CLI for this application.
type CLI struct {
}

// Run runs a command with given args.
func (cli *CLI) Run(args []string) int {
	token, err := readToken()

	if err != nil {
		return 1
	}

	repo, err := getRepo()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	searcher := NewSearcher(repo, token)
	err = searcher.Run()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	return 0
}

func getRepo() (string, error) {
	result, err := exec.Command("git", "config", "remote.origin.url").Output()

	if err != nil {
		return "", err
	}

	dir, filename := filepath.Split(string(result))
	ext := filepath.Ext(filename)
	repo := filename[0 : len(filename)-len(ext)]

	_, user := filepath.Split(dir[0 : len(dir)-1])
	repository := strings.Join([]string{user, repo}, "/")

	return repository, nil
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
