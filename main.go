package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Print("# ")
	reader := bufio.NewReader(os.Stdin)

	text, _ := reader.ReadString('\n')
	// newline on Windows is actually \r\n not \n
	text = strings.Replace(text, "\r\n", "", 1)

	fmt.Printf("Hello, %s!\n", text)
}
