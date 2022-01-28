package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/Gaboose/translit"
)

func main() {
	in := flag.String("in", "", "tuple file path")
	flag.Parse()

	if len(*in) == 0 {
		println("error: --in is required")
		return
	}

	var tl translit.Translator

	if err := tl.ReadFile(*in); err != nil {
		println("error:", err.Error())
	}

	r := bufio.NewReader(os.Stdin)

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			println("error:", err.Error())
			return
		}

		input := strings.TrimRight(line, "\r\n")

		tpl := tl.Translate(input)
		fmt.Println(tpl)
	}
}
