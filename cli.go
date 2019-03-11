// Copyright 2019 hIMEI
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	//"flag"
	"fmt"
	"os"
	"strings"
)

var COMMANDS = []string{
	"header",
	"license",
}

type osArgs struct {
	// file to record
	OutFile string

	// license or header
	Command string

	// name of license or header
	Item string

	// Year
	Year string

	// Autor's name
	Name string

	// Autor's email
	Email string
}

func getIndex(lines []string, line string) int {
	var index int

	for i := range lines {
		if lines[i] == line {
			index = i
		}
	}

	return index
}

func nextArg(arg string) string {
	index := getIndex(os.Args, arg)

	return os.Args[index+1]
}

func isCommand(arg string) bool {
	check := false

	for i := range COMMANDS {
		if COMMANDS[i] == arg {
			check = true
			break
		}
	}

	return check
}

func isOption(arg string) bool {
	check := false

	if strings.HasPrefix(arg, "-") == true {
		check = true
	}

	return check
}

func isArg(arg string) bool {
	check := false

	if isCommand(arg) == false && isOption(arg) == false {
		check = true
	}

	return check
}

func cliParser() osArgs {
	var args = osArgs{
		"LICENSE",
		" ",
		" ",
		" ",
		" ",
		" ",
	}

	// no args
	if len(os.Args) < 2 {
		help()
		os.Exit(1)
	}

	// help is not first arg
	for i := range os.Args {
		if os.Args[i] == "-h" && len(os.Args) > 2 {
			help()
			os.Exit(1)
		}
	}

	// first arg is not command
	for i := range os.Args {
		if isCommand(os.Args[i]) == true && i >= 2 {
			help()
			os.Exit(1)
		}
	}

	switch os.Args[1] {
	//help
	case "-h":
		help()
		os.Exit(1)

	case "headers":
		args.Command = "header"

		if os.Args[2] == "-l" {
			fmt.Println(HEADERS)
			os.Exit(1)
		}

		if os.Args[2] == "-l" && len(os.Args) > 3 {
			help()
			os.Exit(1)
		}

		if isArg(os.Args[2]) == true {
			for i := range HEADERS {
				if os.Args[2] == HEADERS[i] {
					args.Item = os.Args[2]
				}
			}

			// output file provided
			for i := range os.Args {
				if os.Args[i] == "-f" && isArg(nextArg(os.Args[i])) == true {
					args.OutFile = os.Args[i]
				}
			}
		}

	case "license":
		args.Command = "license"

		if os.Args[2] == "-l" {
			fmt.Println(LICENSES)
			os.Exit(1)
		}

		if os.Args[2] == "-l" && len(os.Args) > 3 {
			help()
			os.Exit(1)
		}

		if isArg(os.Args[2]) == true {
			for i := range LICENSES {
				if os.Args[2] == LICENSES[i] {
					args.Item = os.Args[2]
				}
			}

			// output file provided
			for i := range os.Args {
				if os.Args[i] == "-f" && isArg(nextArg(os.Args[i])) == true {
					args.OutFile = os.Args[i]
				}
			}
		}

	}

	return args
}
