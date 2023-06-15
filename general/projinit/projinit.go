package main

import (
	"fmt"
	"strings"
)

type Args struct {
	Language    string
	ProjectName string
}

var Version = "0.0.1"

func main() {
	var a Args
	fmt.Println("ProjInit v" + Version)
	fmt.Print("Project Name (no spaces): ")
	fmt.Scan(&a.ProjectName)
	fmt.Print("Language (c, cpp, cs, go, py): ")
	fmt.Scan(&a.Language)
	a.Language = strings.ToLower(a.Language)
	fmt.Printf("Initializing your project \"%s\" with %s...\n", a.ProjectName, a.Language)
	switch a.Language {
	case "c":
		initC(a)
	case "cpp":
		initCPP(a)
	case "cs":
		initCS(a)
	case "go":
		initGo(a)
	case "py":
		initPy(a)
	default:
		fmt.Println("Unknown language. If you would like it to be added, create a feature request issue on https://github.com/Westsi/CommandLineTools/issues")
	}

}
