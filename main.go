package main

// Basic UTF8 mime decoder using Go libs
// read stdin, output decoded or raw input if error to output
// 2023/03/30 : V0.2

import (
	"bufio"
	"fmt"
	"mime"
	"os"
)

func main() {

	dec := new(mime.WordDecoder)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		if header, err := dec.DecodeHeader(text); err != nil {
			// This could be better handled !
			fmt.Println(text)
		} else {
			fmt.Println(header)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

}
