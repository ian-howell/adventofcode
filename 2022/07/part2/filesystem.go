package main

import (
	"fmt"
	"strings"
)

type File int

type Directory struct {
	Files       map[string]File
	Directories map[string]*Directory
	Parent      *Directory
}

func NewDirectory(parent *Directory) *Directory {
	return &Directory{
		Parent:      parent,
		Files:       map[string]File{},
		Directories: map[string]*Directory{},
	}
}

func (d Directory) Print(level int) {
	for name, size := range d.Files {
		Indent(level, "%v: %d\n", name, size)
	}

	for name, dir := range d.Directories {
		Indent(level, "DIR %v\n", name)
		dir.Print(level + 1)
	}
}

func Indent(level int, msg string, args ...any) {
	fmt.Print(strings.Repeat("  ", level))
	fmt.Printf(msg, args...)
}
