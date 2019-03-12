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
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	//	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"text/template"
	"time"
)

func makeErrString(errConst string) string {
	errString := BOLD + RED + errConst + RESET
	return errString
}

// errFatal is a basic error handler
func errFatal(err error) {
	if err != nil {
		//		log.Fatal(err)
		panic(err)
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

func getText(command, item string) string {
	var path string

	home := os.Getenv("HOME")
	fmt.Println(home)

	if command == "license" {
		path = home + HOMEDIR + LCNS + item
		fmt.Println(path)
	}

	if command == "header" {
		path = home + HOMEDIR + HDRS + item
		fmt.Println(path)
	}

	text, err := ioutil.ReadFile(path)
	errFatal(err)

	return string(text)
}

// ToFile writes rendered text to given file.
func toFile(text string, args *osArgs) {
	var buffer bytes.Buffer
	path, err := os.Getwd()
	errFatal(err)

	path = path + "/" + args.OutFile

	if _, err := os.Stat(path); os.IsExist(err) {
		errString := makeErrString(EXIST)
		newerr := errors.New(errString)
		errFatal(newerr)
	}

	tmpl := template.New(" ")
	tmpl, err = tmpl.Parse(text)
	errFatal(err)

	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)

	//errFatal(err)
	defer file.Close()

	err1 := tmpl.Execute(&buffer, args)
	errFatal(err1)

	str := buffer.String()

	file.WriteString(str + "\n")
	errFatal(err)
}
