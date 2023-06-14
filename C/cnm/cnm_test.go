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
		// expectedCContents := "\n    \n    #include \"" + TESTDIRECTORY + ".h\""
		// expectedHContents := "\n\n#ifndef " + strings.ToUpper(TESTDIRECTORY) + "_H_\n#define " + strings.ToUpper(TESTDIRECTORY) + "_H_\n\n#endif //" + strings.ToUpper(TESTDIRECTORY) + "_H_"
		if isH || isC {
			// not working - attempts to make sure that file contents are correct
			// if isC {
			// 	spl := strings.Split(string(b), "*/")
			// 	if spl[1] != expectedCContents {
			// 		t.Errorf("C File contents incorrect.")
			// 		t.Log(spl[1])
			// 		t.Log(expectedCContents)
			// 		t.FailNow()
			// 	}
			// } else if isH {
			// 	if strings.Split(string(b), "*/")[1] != expectedHContents {
			// 		t.Errorf("H File contents incorrect.")
			// 		t.Log(string(b))
			// 		t.Log(expectedHContents)
			// 		t.FailNow()
			// 	}
			// } else {
			// 	t.Errorf("Something has gone really wrong.")
			// 	t.FailNow()
			// }
		} else {
			t.Logf("isC: %v, isH: %t", isC, isH)
			t.Logf("Comparing to H: %s, C: %s", TESTDIRECTORY+string(os.PathSeparator)+TESTDIRECTORY+".h", TESTDIRECTORY+string(os.PathSeparator)+TESTDIRECTORY+".c")
			t.Logf("\n%s\n%s\n%s", fp, TESTDIRECTORY+string(os.PathSeparator)+TESTDIRECTORY+".h", TESTDIRECTORY+string(os.PathSeparator)+TESTDIRECTORY+".c")
			// t.Errorf("Filename incorrect, expected %s, got %s", TESTDIRECTORY+string(os.PathSeparator)+TESTDIRECTORY+"[.c, .h]", fp)
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
