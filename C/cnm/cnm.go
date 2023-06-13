// c new module - makes a new directory with a premade header and source file

package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func createFile(path string, contents string) {
	os.WriteFile(path, []byte(contents), 0644)
}

func main() {
	t_s := getDTC()
	dir := os.Args[1]

	// make directory
	os.Mkdir(dir, 0755)

	// make header file
	createFile(fmt.Sprintf("%s/%s.h", dir, dir), fmt.Sprintf("/* date = %s */\n\n#ifndef %s_H_\n#define %s_H_\n\n#endif //%s_H", t_s, strings.ToUpper(dir), strings.ToUpper(dir), strings.ToUpper(dir)))

	// make source file
	createFile(fmt.Sprintf("%s/%s.c", dir, dir), fmt.Sprintf("/* date = %s */\n\n#include \"%s.h\"", t_s, dir))
}

func getDTC() string {
	// returns a string in format "17 Apr 2022 13:29"
	// monthtexts := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}

	t := time.Now()
	t_s := t.Format("02 Jan 2006 15:04")
	return t_s
}
