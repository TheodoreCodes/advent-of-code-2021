package utils

import (
	"bufio"
	"os"
	"strconv"
)

type File struct {
	name    string
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

func (recv *File) BinaryInt() int {
	i, _ := strconv.ParseInt(recv.scanner.Text(), 2, 64)
	return int(i)
}

func (recv *File) Text() string {
	return recv.scanner.Text()
}

func (recv *File) Close() {
	recv.f.Close()
}

func NewFile(name string) *File {
	file := &File{
		name: name,
	}

	var err error
	file.f, err = os.Open(file.name)
	if err != nil {
		Panic("Error opening input file", err)
	}

	file.scanner = bufio.NewScanner(file.f)

	return file
}
