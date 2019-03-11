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
	//"os"
)

/*
var (
	outputFlag = flag.String("f", "LICENSE", "save to file")
	helpFlag   = flag.Bool("h", false, "help message")
	listFlag   = flag.Bool("l", false, "list of licenses or headers available")

	headerCmd  = flag.NewFlagSet("header", flag.ExitOnError)
	headerFlag = headerCmd.String()
	licenseCmd = flag.NewFlagSet("license", flag.ExitOnError)
)*/

/*
// ToFile saves results to given file.
func toFile(filename string, parsed []*Host) {
	//	dir := path.Dir(filepath)

	//	if _, err := os.Stat(dir); os.IsNotExist(err) {
	//		errString := makeErrString(NOTEXIST)
	//		newerr := errors.New(NOTEXIST)
	//		ErrFatal(newerr)
	//	}

	if _, err := os.Stat(filename); os.IsExist(err) {
		errString := makeErrString(EXIST)
		newerr := errors.New(errString)
		ErrFatal(newerr)
	}

	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0666)
	ErrFatal(err)
	defer file.Close()

	for i := range parsed {
		if toJson == false {
			file.WriteString(parsed[i].String() + "\n\n\n")
			ErrFatal(err)
		} else {
			file.Write(parsed[i].hostToJson())
		}
	}
}*/

func main() {
	fmt.Println(getContext())
}
