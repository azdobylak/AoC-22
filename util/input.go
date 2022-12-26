package util

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type FileScanner struct {
	io.Closer
	*bufio.Scanner
}

func ReadFileScanner() *FileScanner {
	if len(os.Args) < 2 {
		panic("Missing parameter, provide file name!")
	}

	pwd, _ := os.Getwd()
	path := filepath.Join(pwd, os.Args[1])
	fmt.Println("Reading input from:", path)

	return ReadFileScannerFromPath(path)
}

func ReadFileScannerFromLocalPath(filename string) *FileScanner {
	pwd, _ := os.Getwd()
	path := filepath.Join(pwd, filename)

	return ReadFileScannerFromPath(path)
}

func ReadFileScannerFromPath(path string) *FileScanner {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	return &FileScanner{file, scanner}
}
