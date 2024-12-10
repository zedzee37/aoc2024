package main

import (
	"fmt"
	"os"
	"strconv"
)

type File struct {
	blockCount uint
	id         uint
}

type Fragment struct {
	size  uint
	files []File
}

func (b Fragment) getCapacity() uint {
	var sum uint = 0
	for _, file := range b.files {
		sum += file.blockCount
	}
	return b.size - sum
}

func (b Fragment) isEmpty() bool {
	return len(b.files) == 0
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func toFragments(input string) ([]Fragment, error) {
	fragments := make([]Fragment, len(input))

	isFile := true
	for i, ch := range input {
		num, err := strconv.Atoi(string(ch))
		size := uint(num)
		if err != nil {
			return fragments, err
		}

		var files []File
		if isFile {
			id := uint(i / 2)
			files = append(files, File{
				blockCount: size,
				id:         id,
			})
			isFile = false
		} else {
			isFile = true
		}

		fragments[i] = Fragment{
			size:  size,
			files: files,
		}
	}

	return fragments, nil
}

func findAdequateFragment(fragments []Fragment, file File, max int) int {
	for i := 0; i < max; i++ {
		frag := fragments[i]

		if frag.getCapacity() >= file.blockCount {
			return i
		}
	}
	return -1
}

func arrangeFiles(fragments *[]Fragment) {
	for i := len(*fragments) - 1; i > -1; i-- {
		frag := (*fragments)[i]

		if frag.isEmpty() {
			continue
		}

		for j, file := range frag.files {
			newLoc := findAdequateFragment(*fragments, file, i)
			if newLoc == -1 {
				continue
			}

			frag.files = frag.files[j+1:]
			newFrag := (*fragments)[newLoc]
			newFrag.files = append(newFrag.files, file)
			(*fragments)[newLoc] = newFrag
		}

		(*fragments)[i] = frag
	}
}

func getChecksum(fragments []Fragment) uint {
	var sum uint = 0
	var blockPos uint = 0
	for _, fragment := range fragments {
		var offset uint = 0
		for _, file := range fragment.files {
			id := file.id

			for i := 0; i < int(file.blockCount); i++ {
				sum += id * (blockPos + offset)
				offset += 1
			}
		}
		blockPos += fragment.size
	}
	return sum
}

func main() {
	data, err := os.ReadFile("input.txt")
	check(err)

	text := string(data)
	check(err)

	fragments, err := toFragments(text)
	check(err)

	arrangeFiles(&fragments)
	checksum := getChecksum(fragments)

	fmt.Println(checksum)
}
