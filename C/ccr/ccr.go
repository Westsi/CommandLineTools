// c compile recursively

package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

var EXCLUSIONS []string
var PATH string
var CFILES []string

func main() {

	EXCLUSIONS = append(EXCLUSIONS, ".git")
	EXCLUSIONS = append(EXCLUSIONS, ".vscode")

	_path, _ := os.Getwd()

	PATH = _path

	filepath.Walk(PATH+string(os.PathSeparator), getfiles)

	// fmt.Println(CFILES)
	var args string
	if runtime.GOOS == "windows" {
		args = "-o a.exe "
	} else {
		args = "-o a.out "
	}

	for _, a := range os.Args[1:] {
		args = args + a + " "
	}

	for _, cf := range CFILES {
		args = args + cf + " "
	}

	// fmt.Println(args)

	cmd := exec.Command("gcc", args)
	fmt.Println(cmd)

	// gcc says no input files if does not write to file
	var fname string
	if runtime.GOOS == "windows" {
		fname = "adfhgh87obbdscvj.bat"
		err := os.WriteFile(fname, []byte(fmt.Sprint(cmd)), 0644)
		if err != nil {
			fmt.Println(err)
		}
		cmd = exec.Command(".\\" + fname)
	} else {
		fname = "adfhgh87obbdscvj.sh"
		err := os.WriteFile(fname, []byte("#!/bin/sh"+"\n"+fmt.Sprint(cmd)), 256|128|64)
		if err != nil {
			fmt.Println(err)
		}
		cmd = exec.Command("./" + fname)
	}

	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("STDOUT:", outb.String(), "STDERR:", errb.String())

	os.Remove(fname)

}
func getfiles(p string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	pl := len(PATH)
	psliced := p[pl+1:]
	basefolder := strings.Split(psliced, string(os.PathSeparator))[0]
	if !contains(EXCLUSIONS, basefolder) {
		if isCSrc(psliced) {
			CFILES = append(CFILES, psliced)
		}
	}
	return nil
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func isCSrc(s string) bool {
	if len(s) < 3 {
		return false
	}
	sliced := s[len(s)-2:]
	return sliced == ".c"
}
