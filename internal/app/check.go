package app

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Check(s [][]string) {
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < len(s); i++ {
		subCheck(s[i], *reader)
	}
}

func subCheck(p []string, reader bufio.Reader) {
	fmt.Print(p[0] + ": ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	if text == p[1] {
		fmt.Println("success")
	} else {
		fmt.Println("wrong. right - ", p[1],
			" len: current ", strconv.Itoa(len(text)),
			" right ", strconv.Itoa(len(p[1])))
	}
}
