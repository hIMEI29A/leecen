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
	"errors"
	"fmt"
	"os"
	"strings"
)

// Structure to store arguments parsed
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

// Gets os.Args' index of given cli arg
func getIndex(lines []string, line string) int {
	var index int

	for i := range lines {
		if lines[i] == line {
			index = i
		}
	}

	return index
}

// Gets cli argument followed given one.
// If given argument is last argument, gets the error
func nextArg(arg string) string {
	var nindex string
	index := getIndex(os.Args, arg)

	if index != len(os.Args)-1 {
		nindex = os.Args[index+1]
	} else {
		err := makeErrString(MISSING)
		newerr := errors.New(err)
		errFatal(newerr)
	}

	return nindex
}

// // Checks if given cli option is command
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

// // Checks if given cli option is option e.g. "-h" or "-l"
func isOption(arg string) bool {
	check := false

	if strings.HasPrefix(arg, "-") == true {
		check = true
	}

	return check
}

// Checks if given cli option is argument
func isArg(arg string) bool {
	check := false

	if isCommand(arg) == false && isOption(arg) == false {
		check = true
	}

	return check
}

// Checks if given arg is header or license
func validateArg(arg string) bool {
	check := false

	if isArg(arg) == true {
		for i := range LICENSES {
			if LICENSES[i] == arg {
				check = true
				break
			}
		}
	}

	return check
}

// Help
func help() {
	fmt.Print("\n" + BOLD + GRN + "leecen" + RESET + HMESS + "\n\n")
	fmt.Print(CYN + BOLD + "Usage: " + RESET + USAGE + "\n")
	fmt.Print(GRN + BOLD + LINE + RESET + "\n\n")
	fmt.Print("\t" + YEL + "-h" + RESET + " - read this message" + "\n")
	fmt.Print("Commands:" + "\n")
	fmt.Print("\t" + YEL + "header" + RESET + GHEADER + "\n")
	fmt.Print("\t" + YEL + "license" + RESET + GLICENSE + "\n")
	fmt.Print("Options:" + "\n")
}

// Lists headers or licenss available
func list(arg string) {
	if arg == "header" {
		for i := range LICENSES {
			if strings.HasSuffix(LICENSES[i], "-header") == true {
				fmt.Println(LICENSES[i] + " ")
			}
		}
	} else {
		for i := range LICENSES {
			if strings.HasSuffix(LICENSES[i], "-header") == false {
				fmt.Println(LICENSES[i] + " ")
			}
		}
	}
}

// Cli parser
func cliParser() *osArgs {
	var args = &osArgs{
		"LICENSE",
		" ",
		" ",
		" ",
		" ",
		" ",
	}

	// no args
	if len(os.Args) < 2 {
		err := makeErrString(NOARGS)
		newerr := errors.New(err)
		errFatal(newerr)
	}

	// help
	if os.Args[1] == "-h" {
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
			err := makeErrString(FIRSTARG)
			newerr := errors.New(err)
			errFatal(newerr)
		}
	}

	if isCommand(os.Args[1]) == true {
		args.Command = os.Args[1]

		if os.Args[2] != "-l" && validateArg(os.Args[2]) == false {
			err := makeErrString(HEXIST)
			newerr := errors.New(err)
			errFatal(newerr)
		} else {
			args.Item = os.Args[2]
		}

		// list licences or headers available
		if os.Args[2] == "-l" {
			list(os.Args[1])
			os.Exit(1)
		}

		// output file provided
		for i := range os.Args {
			if os.Args[i] == "-f" && isArg(nextArg(os.Args[i])) == true {
				args.OutFile = nextArg(os.Args[i])
			}
		}

		// email provided
		for i := range os.Args {
			if os.Args[i] == "-e" && isArg(nextArg(os.Args[i])) == true {
				args.Email = nextArg(os.Args[i])
			}
		}

		// name provided
		for i := range os.Args {
			if os.Args[i] == "-n" && isArg(nextArg(os.Args[i])) == true {
				args.Name = nextArg(os.Args[i])
			}
		}

		// year provided
		for i := range os.Args {
			if os.Args[i] == "-y" && isArg(nextArg(os.Args[i])) == true {
				args.Year = nextArg(os.Args[i])
			}
		}
	}

	return args
}
