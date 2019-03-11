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
	//	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func makeErrString(errConst string) string {
	errString := BOLD + RED + errConst + RESET
	return errString
}

// ErrFatal is a basic error handler
func errFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getYear() int {
	return time.Now().Year()
}

func getCreds() (string, string) {
	var mail, name *exec.Cmd

	name = exec.Command(GIT, CONFIG, GET, UNAME)
	username, err := name.Output()
	errFatal(err)

	mail = exec.Command(GIT, CONFIG, GET, UEMAIL)
	uemail, err := mail.Output()
	errFatal(err)

	return strings.Trim(string(username), "\n"), strings.Trim(string(uemail), "\n")
}

// iToa converts int to string
func iToa(i int) string {
	str := strconv.Itoa(i)

	return str
}

func getContext() (string, string, string) {
	mail, name := getCreds()
	year := iToa(getYear())

	return mail, name, year
}

func getFilepath(filename string) string {
	home := os.Getenv("HOME")

	return home + HOMEDIR + filename
}
