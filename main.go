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

//"flag"
//"fmt"
//"os"

func main() {
	args := cliParser()

	if args.Year == " " {
		_, _, year := getContext()
		args.Year = year
	}

	if args.Name == " " {
		_, name, _ := getContext()
		args.Name = name
	}

	if args.Email == " " {
		email, _, _ := getContext()
		args.Email = email
	}

	text := getText(args.Command, args.Item)

	toFile(text, args)
}
