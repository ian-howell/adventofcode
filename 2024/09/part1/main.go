package main

import "fmt"

type Block struct {
	Start int
	End   int
}

func (b Block) Size() int { return b.End - b.Start }

type File struct {
	ID     int
	Blocks []Block
}

func main() {
	input := getInput()
	files, emptyBlocks := parseInput(input)

	compact(files, emptyBlocks)
	checksum := calculateChecksum(files)
	fmt.Println(checksum)
}

func compact(files []File, emptyBlocks []Block) {
	emptyBlockIndex := 0
	fileIndex := len(files) - 1
	for files[fileIndex].Blocks[0].Start > emptyBlocks[emptyBlockIndex].Start {
		// This is the *current* file size
		fileSize := files[fileIndex].Blocks[0].Size()
		emptyBlockSize := emptyBlocks[emptyBlockIndex].Size()
		if fileSize == emptyBlockSize {
			// Just move the whole file
			files[fileIndex].Blocks[0] = emptyBlocks[emptyBlockIndex]
			fileIndex--
			emptyBlockIndex++
		} else if fileSize < emptyBlockSize {
			// Put the file in the empty space
			files[fileIndex].Blocks[0].Start = emptyBlocks[emptyBlockIndex].Start
			files[fileIndex].Blocks[0].End = emptyBlocks[emptyBlockIndex].Start + fileSize
			fileIndex--
			// Adjust the empty space to account for the new file
			emptyBlocks[emptyBlockIndex].Start += fileSize
		} else {
			files[fileIndex].Blocks[0].End -= emptyBlockSize
			files[fileIndex].Blocks = append(
				files[fileIndex].Blocks,
				emptyBlocks[emptyBlockIndex],
			)
			emptyBlockIndex++
		}
	}
}

func calculateChecksum(files []File) int {
	checksum := 0
	for _, file := range files {
		// TODO: Optimize this using math
		for _, block := range file.Blocks {
			for i := block.Start; i < block.End; i++ {
				checksum += file.ID * i
			}
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
			file := File{ID: i / 2, Blocks: []Block{block}}
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
