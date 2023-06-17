package main

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestCNM(t *testing.T) {
	TESTDIRECTORY := "test_cnm"

	t.Cleanup(func() {
		os.RemoveAll(TESTDIRECTORY)
	})

	t.Log("Testing CNM")

	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("..\\..\\cnm.exe", TESTDIRECTORY)
	} else {
		cmd = exec.Command("../../cnm", TESTDIRECTORY)
	}

	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	err := cmd.Run()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	// t.Logf("STDOUT: %s\nSTDERR: %s\n", outb.String(), errb.String())
	var filepaths []string

	err = filepath.Walk(TESTDIRECTORY, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			filepaths = append(filepaths, path)
		}
		return nil
	})
	if err != nil {
		t.Errorf("ERROR IN TEST SUITE filepath.Walk: %s", err.Error())
		t.FailNow()
	}

	for _, fp := range filepaths {
		b, e := os.ReadFile(fp)
		b = b
		if e != nil {
			t.Errorf("ERROR IN TEST SUITE os.ReadFile: %s", e.Error())
			t.FailNow()
		}
		// t.Log(string(b))
		fpSplit := strings.Split(fp, string(os.PathSeparator))
		isH := isEqual(fpSplit, []string{TESTDIRECTORY, TESTDIRECTORY + ".h"})
		isC := isEqual(fpSplit, []string{TESTDIRECTORY, TESTDIRECTORY + ".c"})
		if !(isH || isC) {
			t.Errorf("ERROR IN TEST SUITE %s", fp)
			t.FailNow()
		}

	}

}

func isEqual(fpSplit []string, expected []string) bool {
	for i, e := range expected {
		if e != fpSplit[i] {
			return false
		}
	}
	return true
}
