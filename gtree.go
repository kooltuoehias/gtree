package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/kooltuoehias/gtree/internal/treer"
)

func main() {
	flagLevel := flag.Int("L", 0, "displaying only a certain number of directories deep")
	flag.Parse()
	pwd, _ := os.Getwd()
	treer := treer.FolderFiles{
		Path: pwd,
	}

	fmt.Printf("%s\n", filepath.Base(pwd))
	treer.Traverse(*flagLevel + 1)
}
