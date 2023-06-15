package main

import (
	"os"
	"strings"
)

func initC(a Args) {
	os.MkdirAll(strings.ToLower(a.ProjectName), 0755)
}
