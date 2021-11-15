package console

import (
	"bufio"
	"emorisse.fr/go-calculator/pkg/parser"
	"fmt"
	"log"
	"os"
	"strings"
)

func Start() {
	var running = true

	for running {
		fmt.Print("> ")
		var input = ReadLine()

		if input == "" {
			continue
		}

		if input == "quit" || input == "exit" {
			running = false
			continue
		}

		opt, err := parser.Parse(input)

		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("=", opt.Eval().GetString())
	}

	fmt.Println("Bye!")
}

func ReadLine() string {
	var reader = bufio.NewReader(os.Stdin)
	var line, err = reader.ReadString('\n')

	if err != nil {
		log.Fatalln(err)
	}

	return strings.Trim(line, " \n\r")
}
