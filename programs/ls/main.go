package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
)

func b2i(b bool) uint8 {
	if b {
		return 1
	}
	return 0
}

func getWorkingDir() string {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	return path
}

func main() {
	workingDir := getWorkingDir()
	fmt.Printf("Files in %v:\n", workingDir)
	entries, _ := os.ReadDir(workingDir)

	formatEntry := func(entry os.DirEntry) string {
		fi, err := os.Lstat(workingDir + "/" + entry.Name())
		if err != nil {
			log.Fatal(err)
		}
		isLink := fi.Mode()&fs.ModeSymlink != 0
		return fmt.Sprintf("%v%v %v", entry.Name(), []string{"", "/"}[b2i(entry.IsDir())], []string{"", "(L)"}[b2i(isLink)])
	}

	for _, entry := range entries {
		fmt.Printf(" - %v\n", formatEntry(entry))
	}
}
