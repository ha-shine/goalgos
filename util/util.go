package util

import (
	"bufio"
	"io"
	"os"
)

func ReadFile(file string, fn func(string)) {
	f, _ := os.Open(file)
	reader := bufio.NewReader(f)
	word, err := reader.ReadString(' ')
	for err != io.EOF {
		fn(word)
		word, err = reader.ReadString(' ')
	}
}
