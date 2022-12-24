package main

import (
	"aoc22/day07/file"
	"aoc22/util"
	"fmt"
	"strings"
)

func ParseFileTree(scanner *util.FileScanner) *file.File {
	var currentDir *file.File
	root := file.CreateRootDir()

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, " ")
		if words[0] == "$" {
			if words[1] == "cd" {
				dirname := words[2]
				if dirname == "/" {
					currentDir = root
				} else if dirname == ".." {
					currentDir = currentDir.Parent
				} else {
					currentDir = currentDir.ChangeDir(dirname)
				}
			}
		} else {
			if words[0] == "dir" {
				currentDir.CreateDir(words[1])
			} else {
				size := file.ParseSize(words[0])
				currentDir.CreateFile(words[0], size)
			}
		}
	}

	return root
}

type DirSummary struct {
	File      *file.File
	TotalSize int
}

func CollectDirsLessOrEqual(f *file.File, maxSize int, matchingDirs *[]DirSummary) int {
	var totalSize int

	for _, subdir := range f.Childs {
		if subdir.Childs == nil {
			totalSize += subdir.Size
		} else {
			totalSize += CollectDirsLessOrEqual(subdir, maxSize, matchingDirs)
		}
	}

	if totalSize <= maxSize {
		*matchingDirs = append(*matchingDirs, DirSummary{File: f, TotalSize: totalSize})
	}

	return totalSize
}

func solveA() int {
	var scanner *util.FileScanner = util.ReadFileScanner()
	defer scanner.Close()

	root := ParseFileTree(scanner)
	var dirs []DirSummary
	const maxSize int = 100_000
	CollectDirsLessOrEqual(root, maxSize, &dirs)
	fmt.Printf("Found %d dirs with size less or equal to %d\n", len(dirs), maxSize)

	var sum int
	for _, sumDir := range dirs {
		sum += sumDir.TotalSize
	}

	return sum
}

func solveB() int {
	var scanner *util.FileScanner = util.ReadFileScanner()
	defer scanner.Close()

	root := ParseFileTree(scanner)
	fmt.Println(root)

	return 0
}

func main() {
	fmt.Println("=== DAY 07a solution ===")
	fmt.Println(solveA())
	//fmt.Println("=== DAY 07b solution ===")
	//fmt.Println(solveB())
}
