package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/pkg/profile"
	"os"
	"strconv"
)

func main() {
	defer profile.Start().Stop()
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	var input, _, _ = reader.ReadLine()
	var height, _ = strconv.Atoi(string(input))
	var level int = 0
	for i := 0; i < height; i++ {
		var buffer bytes.Buffer
		for j := 0; j < height-level; j++ {
			buffer.WriteString(" ")
		}
		for j := 0; j < level; j++ {
			buffer.WriteString("#")
		}
		level += 1
		fmt.Println(buffer.String())
	}
}
