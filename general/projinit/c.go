package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func initC(a Args) {
	t_s := getDTC()
	os.MkdirAll(strings.ToLower(a.ProjectName), 0755)
	contents := []byte(fmt.Sprintf("/* date = %s */\n", t_s) + "#include<stdio.h>\n\nint main() {\n    return 0;\n}\n")
	os.WriteFile(strings.ToLower(a.ProjectName)+"/main.c", contents, 0644)

	var buildcmd []byte
	buildcmd = []byte("@ECHO off\n\nccr")
	os.WriteFile(strings.ToLower(a.ProjectName)+"/build.bat", buildcmd, 0644)
	buildcmd = []byte("#!/bin/sh\n\nccr")
	os.WriteFile(strings.ToLower(a.ProjectName)+"/build.sh", buildcmd, 256|128|64)

	fmt.Println(a.ProjectName + " created")
}

func getDTC() string {
	// returns a string in format "17 Apr 2022 13:29"
	// monthtexts := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}

	t := time.Now()
	t_s := t.Format("02 Jan 2006 15:04")
	return t_s
}
