package main

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
		header, err := dec.DecodeHeader(text)
		if err != nil {
			// panic(err)
			fmt.Println(text)
		}
		fmt.Println(header)

	}

	if err := scanner.Err(); err != nil {
		// Handle error.
		fmt.Println(err)
	}

}
