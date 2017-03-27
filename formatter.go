package main

import (
	"fmt"
	"strings"

	"github.com/google/go-github/github"
)

// Formatter is a formatter for issues.
type Formatter struct {
	Issues []github.Issue
}

// Format returns formatted strings.
func (formatter *Formatter) Format() string {
	buf := []string{}

	numberMaxLength := 0
	usernameMaxLength := 0

	for _, issue := range formatter.Issues {
		length := len(fmt.Sprint(issue.GetNumber()))
		if length > numberMaxLength {
			numberMaxLength = length
		}

		length = len(issue.User.GetLogin())
		if length > usernameMaxLength {
			usernameMaxLength = length
		}
	}

	for _, issue := range formatter.Issues {
		number := issue.GetNumber()
		numberSpace := strings.Repeat(" ", numberMaxLength-len(fmt.Sprint(number)))

		username := issue.User.GetLogin()
		usernameSpaces := strings.Repeat(" ", usernameMaxLength-len(username))

		formatted := fmt.Sprintf("#%d%v %v%v %v", number, numberSpace, username, usernameSpaces, issue.GetTitle())
		buf = append(buf, formatted)
	}

	return strings.Join(buf, "\n")
}
