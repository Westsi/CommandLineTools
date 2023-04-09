// c new module - makes a new directory with a premade header and source file

package main

import (
	"fmt"
	"os"
	"strings"
)

func createFile(path string, contents string) {
	os.WriteFile(path, []byte(contents), 0644)
}

func main() {
	dir := os.Args[1]

	// make directory
	os.Mkdir(dir, 0755)

	// make header file
	createFile(fmt.Sprintf("%s/%s.h", dir, dir), fmt.Sprintf("#ifndef %s_H_\n#define %s_H_\n\n#endif", strings.ToUpper(dir), strings.ToUpper(dir)))

	// make source file
	createFile(fmt.Sprintf("%s/%s.c", dir, dir), fmt.Sprintf("#include \"%s.h\"", dir))
}
