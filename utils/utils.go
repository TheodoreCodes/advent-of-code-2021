package utils

import (
	"fmt"
	"os"
)

func Panic(message string, err error) {
	fmt.Printf("%s: %+v", message, err)
	os.Exit(1)
}
