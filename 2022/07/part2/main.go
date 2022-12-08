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

	totalSize := GetTotalSize(pwd)
	currentlyFree := 70_000_000 - totalSize
	needToFree := 30_000_000 - currentlyFree

	_, toDelete := Calc(pwd, needToFree)
	fmt.Println(toDelete)
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

func Calc(pwd *Directory, need int) (int, int) {
	totalSize := 0
	answer := 70_000_000
	for _, dir := range pwd.Directories {
		sizeOfChild, smallestSoFar := Calc(dir, need)
		if smallestSoFar >= need && smallestSoFar < answer {
			answer = smallestSoFar
		}
		totalSize += sizeOfChild
	}
	for _, file := range pwd.Files {
		totalSize += int(file)
	}
	if totalSize >= need && totalSize < answer {
		answer = totalSize
	}
	return totalSize, answer
}

func GetTotalSize(pwd *Directory) int {
	totalSize := 0
	for _, dir := range pwd.Directories {
		totalSize += GetTotalSize(dir)
	}
	for _, file := range pwd.Files {
		totalSize += int(file)
	}
	return totalSize
}
