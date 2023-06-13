package main

import (
	"bytes"
	"os/exec"
	"runtime"
	"testing"
)

func TestCCR(t *testing.T) {
	t.Log("Testing CCR")
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("..\\..\\ccr.exe")
	} else {
		cmd = exec.Command("../../ccr")
	}

	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	err := cmd.Run()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("STDOUT: %s\nSTDERR: %s\n", outb.String(), errb.String())

	if runtime.GOOS == "windows" {
		cmd = exec.Command("..\\..\\a.exe")
	} else {
		cmd = exec.Command("../../a.out")
	}

	// reset buffers
	outb.Reset()
	errb.Reset()

	cmd.Stdout = &outb
	cmd.Stderr = &errb
	err = cmd.Run()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Logf("STDOUT: %s\nSTDERR: %s\n", outb.String(), errb.String())
}
