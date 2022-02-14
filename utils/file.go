package utils

import (
	"bufio"
	"os"
	"strconv"
)

type File struct {
	f       *os.File
	scanner *bufio.Scanner
}

func (recv *File) Scan() bool {
	return recv.scanner.Scan()
}

func (recv *File) Int() int {
	i, _ := strconv.Atoi(recv.scanner.Text())
	return i
}

func (recv *File) Text() string {
	return recv.scanner.Text()
}

func (recv *File) Close() {
	recv.f.Close()
}

func NewFile(name string) *File {
	f, err := os.Open(name)
	if err != nil {
		Panic("Error opening input.txt file", err)
	}

	return &File{
		f:       f,
		scanner: bufio.NewScanner(f),
	}
}
