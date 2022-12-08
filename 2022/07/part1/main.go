package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// var pwd *Directory
	var pwd *Directory
	lines := getLines()
	for lineNo := 0; lineNo < len(lines); lineNo++ {
		// fmt.Printf("Processesing line %d: <%v>\n", lineNo, lines[lineNo])
		parts := strings.Split(lines[lineNo], " ")
		if parts[0] != "$" {
			panic("Recieved non-command: " + lines[lineNo])
		}
		// fmt.Printf("Recieved command <%v>\n", lines[lineNo])

		if parts[1] == "ls" {
			lineNo = Ls(pwd, lines, lineNo)
		} else if parts[1] == "cd" {
			pwd = Cd(pwd, parts[2])
		}
	}

	for i := 0; i < 10; i++ {
		pwd = Cd(pwd, "..")
	}

	// pwd.Print(0)

	_, result := Calc(pwd)
	fmt.Println(result)
}

func Ls(pwd *Directory, lines []string, lineNo int) int {
	var i int
	for i = 1; (lineNo+i < len(lines)) && (!strings.HasPrefix(lines[lineNo+i], "$")); i++ {
		parts := strings.Split(lines[lineNo+i], " ")
		if parts[0] == "dir" {
			pwd.Directories[parts[1]] = NewDirectory(pwd)
		} else {
			pwd.Files[parts[1]] = File(atoi(parts[0]))
		}
	}
	return lineNo + i - 1
}

func Cd(pwd *Directory, where string) *Directory {
	switch where {
	case "/":
		pwd = NewDirectory(nil)
		pwd.Parent = pwd
		return pwd
	case "..":
		return pwd.Parent
	default:
		return pwd.Directories[where]
	}
}

func getLines() []string {
	var lines []string
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	return lines
}

func atoi(a string) int {
	i, _ := strconv.Atoi(a)
	return i
}

func Calc(pwd *Directory) (int, int) {
	totalSize := 0
	answer := 0
	for _, dir := range pwd.Directories {
		sizeOfChild, soFar := Calc(dir)
		answer += soFar
		totalSize += sizeOfChild
	}
	for _, file := range pwd.Files {
		totalSize += int(file)
	}
	if totalSize <= 100_000 {
		// If the total size of this directory is at most 100,000, include it in the answer
		answer += totalSize
	}
	return totalSize, answer
}
