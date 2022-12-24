package main

import (
	"aoc22/day07/file"
	"aoc22/util"
	"fmt"
	"sort"
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

func AccumulateDirSizes(f *file.File, accDirs *[]DirSummary) int {
	var totalSize int

	for _, subdir := range f.Childs {
		if subdir.Childs == nil {
			totalSize += subdir.Size
		} else {
			totalSize += AccumulateDirSizes(subdir, accDirs)
		}
	}

	if f.Childs != nil {
		*accDirs = append(*accDirs, DirSummary{File: f, TotalSize: totalSize})
	}

	return totalSize
}

func CalculateDirSizes(root *file.File) []DirSummary {
	var matchingDirs []DirSummary
	AccumulateDirSizes(root, &matchingDirs)

	return matchingDirs
}

func solveA() int {
	var scanner *util.FileScanner = util.ReadFileScanner()
	defer scanner.Close()

	root := ParseFileTree(scanner)
	const maxSize int = 100_000
	dirs := CalculateDirSizes(root)

	fmt.Printf("Found %d dirs\n", len(dirs))

	var sum int
	var i int
	for _, sumDir := range dirs {
		if sumDir.TotalSize <= maxSize {
			sum += sumDir.TotalSize
			i += 1
		}
	}
	fmt.Printf("Found %d dirs with size less or equal to %d\n", i, maxSize)

	return sum
}

func solveB() int {
	var scanner *util.FileScanner = util.ReadFileScanner()
	defer scanner.Close()

	const diskSize int = 70_000_000
	const requestedSize int = 30_000_000

	root := ParseFileTree(scanner)
	dirs := CalculateDirSizes(root)

	sort.Slice(dirs, func(i, j int) bool {
		return (dirs)[i].TotalSize < (dirs)[j].TotalSize
	})

	unusedSize := diskSize - dirs[len(dirs)-1].TotalSize

	for _, dir := range dirs {
		if unusedSize+dir.TotalSize >= requestedSize {
			return dir.TotalSize
		}
	}

	return -1
}

func main() {
	fmt.Println("=== DAY 07a solution ===")
	fmt.Println(solveA())
	fmt.Println("=== DAY 07b solution ===")
	fmt.Println(solveB())
}
