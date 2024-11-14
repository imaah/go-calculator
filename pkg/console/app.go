package console

import (
	"bufio"
	"fmt"
	"github.com/imaah/go-calculator/pkg/parser"
	"log"
	"os"
	"strings"
)

func Start() {
	running := true

	for running {
		fmt.Print("> ")
		input := ReadLine()

		if input == "" {
			continue
		}

		if input == "quit" || input == "exit" {
			running = false
			continue
		}

		opt, err := parser.ParseV2(input)

		if err != nil {
			errFormat := fmt.Errorf("Parsing error: %w\n", err)
			log.Println(errFormat)
		} else {
			fmt.Println(opt)
			fmt.Println("=", opt.Eval())
		}
	}

	fmt.Println("Bye!")
}

func ReadLine() string {
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')

	if err != nil {
		log.Fatalln(err)
	}

	return strings.Trim(line, " \n\r")
}
