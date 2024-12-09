package main

import (
	"fmt"
)

type Block struct {
	Start int
	End   int
}

func (b Block) Size() int { return b.End - b.Start }

type File struct {
	Block
	ID int
}

func main() {
	input := getInput()
	files, emptyBlocks := parseInput(input)

	compact(files, emptyBlocks)
	checksum := calculateChecksum(files)
	fmt.Println(checksum)
}

func compact(files []File, emptyBlocks []Block) {
	for f := len(files) - 1; f >= 0; f-- {
		for e := range emptyBlocks {
			if files[f].Start < emptyBlocks[e].Start {
				break
			}
			fileSize := files[f].Size()
			if fileSize <= emptyBlocks[e].Size() {
				files[f].Block = Block{
					Start: emptyBlocks[e].Start,
					End:   emptyBlocks[e].Start + fileSize,
				}
				emptyBlocks[e].Start += fileSize
				break
			}
		}
	}
}

func calculateChecksum(files []File) int {
	checksum := 0
	for _, file := range files {
		// TODO: Optimize this using math
		for i := file.Start; i < file.End; i++ {
			checksum += file.ID * i
		}
	}
	return checksum
}

func parseInput(input []byte) ([]File, []Block) {
	files := []File{}
	emptyBlocks := []Block{}
	addr := 0
	for i, b := range input {
		size := byteToInt(b)
		block := Block{
			Start: addr,
			End:   addr + size,
		}
		if i%2 == 0 {
			file := File{ID: i / 2, Block: block}
			files = append(files, file)
		} else if size > 0 {
			emptyBlocks = append(emptyBlocks, block)
		}
		addr += size
	}
	return files, emptyBlocks
}

func getInput() []byte {
	var input []byte
	_, err := fmt.Scanln(&input)
	if err != nil {
		panic(fmt.Sprintf("failed to read the really big string: %v", err))
	}
	return input
}

func byteToInt(b byte) int {
	return int(b & 0x0f)
}
