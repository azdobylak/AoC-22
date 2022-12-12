package util

import (
    "os"
    "io"
    "fmt"
    "bufio"
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

	file, err := os.Open(path)
    if err != nil {
        panic(err)
    }

    scanner := bufio.NewScanner(file)
    return &FileScanner{file, scanner}
}
