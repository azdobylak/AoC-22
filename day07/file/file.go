package file

import (
	"fmt"
	"strconv"
)

type File struct {
	Parent *File
	Childs []*File

	Name string
	Size int
}

func ParseSize(size string) int {
	s, err := strconv.Atoi(size)
	if err != nil {
		errString := fmt.Sprintf("Unable to parse size: %s. Error: #%v", size, err)
		panic(errString)
	}
	return s
}

func CreateRootDir() *File {
	return &File{nil, nil, "/", 0}
}

func (this *File) CreateFile(name string, size int) {
	File := File{this, nil, name, size}
	this.Childs = append(this.Childs, &File)
}

func (this *File) CreateDir(name string) {
	File := File{this, nil, name, 0}
	this.Childs = append(this.Childs, &File)
}

func (this *File) ChangeDir(name string) *File {
	for _, subdir := range this.Childs {
		if subdir.Name == name {
			return subdir
		}
	}
	return nil
}
