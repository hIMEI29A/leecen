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

// Output colorizing
const (
	RED   = "\x1B[31m"
	GRN   = "\x1B[32m"
	YEL   = "\x1B[33m"
	BLU   = "\x1B[34m"
	CYN   = "\x1B[36m"
	WHT   = "\x1B[97m"
	RESET = "\x1B[0m"
	BOLD  = "\x1B[1m"
)

// useful constants
const (
	GIT      = "git"
	GET      = "--get"
	CONFIG   = "config"
	UEMAIL   = "user.email"
	UNAME    = "user.name"
	EMPTY    = " "
	HOMEDIR  = "/.leecen/templates"
	LCNS     = "/licenses/"
	HDRS     = "/headers/"
	HMESS    = " - licenses for your project"
	USAGE    = "leecen [-h] | COMMAND [ARG] [OPTIONS]"
	LINE     = "====================================================="
	GHEADER  = " - make license header for given file"
	GLICENSE = " - make LICENSE file"
	GHELP    = " - read this message"
	GLIST    = " - list licenses or headers available"
	GMAIL    = " - set custom autor's email"
	GNAME    = " - set custom autor's name"
	GYEAR    = " - set custom year"
	GFILE    = " - set custom license file's name"
)

// Error messages
const (
	NOTEXIST = "Given path does not exist"
	EXIST    = "File already exist, we'll not rewrite it"
	NOEMAIL  = "WARNING: Please configure your git. I need your user email"
	NONAME   = "WARNING: Please configure your git. I need your user name"
	NOARGS   = "No arguments is provided. Try leecene -h for help"
	FIRSTARG = "Commands must be used. Try leecene -h for help"
	HEXIST   = "License or header does not exist. Try leecene -h for help"
	MISSING  = "Argument is missing"
)

// all licenses and headers
var LICENSES = []string{
	"agpl-3.0",
	"bsd-2-clause",
	"epl-1.0",
	"isc",
	"mit",
	"unlicense",
	"apache-2.0",
	"bsd-3-clause",
	"gpl-2.0",
	"lgpl-2.1",
	"mpl-2.0",
	"artistic-2.0",
	"cc0",
	"gpl-3.0",
	"lgpl-3.0",
	"no-license",
	"agpl-3.0-header",
	"apache-2.0-header",
	"gpl-2.0-header",
	"gpl-3.0-header",
	"lgpl-2.1-header",
}

// Cli commands
var COMMANDS = []string{
	"header",
	"license",
}

// Cli options
var OPTIONS = []string{
	"-h",
	"-l",
	"-e",
	"-n",
	"-y",
	"-f",
}

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

// Checks if given cli option is option e.g. "-h" or "-l"
func isOption(arg string) bool {
	check := false

	for i := range os.Args {
		if arg == os.Args[i] {
			check = true
		}
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
	fmt.Print("Commands:" + "\n")
	fmt.Print("\t" + YEL + "header" + RESET + GHEADER + "\n")
	fmt.Print("\t" + YEL + "license" + RESET + GLICENSE + "\n")
	fmt.Print("Options:" + "\n")
	fmt.Print("\t" + YEL + "-h" + RESET + GHELP + "\n")
	fmt.Print("\t" + YEL + "-l" + RESET + GLIST + "\n")
	fmt.Print("\t" + YEL + "-e" + RESET + GMAIL + "\n")
	fmt.Print("\t" + YEL + "-n" + RESET + GNAME + "\n")
	fmt.Print("\t" + YEL + "-y" + RESET + GYEAR + "\n")
	fmt.Print("\t" + YEL + "-f" + RESET + GFILE + "\n")
	fmt.Print("\n")
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
	for i := range os.Args {
		if os.Args[i] == "-h" {
			help()
			os.Exit(1)
		}
	}

	// first arg is not command
	if isCommand(os.Args[1]) == false && os.Args[1] != "-h" {
		err := makeErrString(FIRSTARG)
		newerr := errors.New(err)
		errFatal(newerr)
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
