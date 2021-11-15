package runner

import (
	"emorisse.fr/go-calculator/pkg/console"
	"emorisse.fr/go-calculator/pkg/web"
	"errors"
	"fmt"
	"regexp"
)

type Arguments map[string]string

var ArgumentAliases = map[rune]string{
	'w': "web-server",
	'p': "port",
	'a': "bind-address",
	'h': "help",
}

var KeyRegex = regexp.MustCompile(`^--([a-zA-z0-9\\-]+)|-([a-zA-Z])$`)

func (a Arguments) GetOrDefault(key, defaultValue string) string {
	if value, contains := a[key]; contains {
		return value
	}
	return defaultValue
}

func Run(args []string) {
	var arguments = ReadArguments(args)

	fmt.Println(arguments)

	if usingWeb, contains := arguments["web-server"]; contains && usingWeb == "true" {
		var port = arguments.GetOrDefault("port", "8080")
		var addr = arguments.GetOrDefault("bind-address", "localhost")

		web.StartServer(addr, port)
	} else {
		console.Start()
	}
}

func ReadArguments(args []string) Arguments {
	var key = ""
	var arguments = make(Arguments)

	for _, arg := range args {
		fmt.Println(arg, KeyRegex.MatchString(arg))
		if KeyRegex.MatchString(arg) {
			if key != "" {
				arguments[key] = "true"
			}

			// ignoring error because we know that arg matches the regex.
			key, _ = ReadArgumentKey(arg)
			fmt.Println(arg, key)
		} else if key != "" {
			arguments[key] = arg
			key = ""
		}
	}

	if key != "" {
		arguments[key] = "true"
	}

	return arguments
}

func ReadArgumentKey(key string) (string, error) {
	var groups = KeyRegex.FindAllStringSubmatch(key, 1)

	if len(groups) > 0 {
		var parsedKey = groups[0][1]

		if parsedKey == "" {
			var shortcut = rune(groups[0][2][0])
			if fKey, contains := ArgumentAliases[shortcut]; contains {
				return fKey, nil
			}
		}

		return parsedKey, nil
	}

	return "", errors.New("DoesNotMatchKey")
}
