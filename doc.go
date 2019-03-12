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

/*
Description

**leecen** - License generator for your project.
It creates file LICENSE in current directory with context of given license.
It may also create file with custom name and write license header to file.

Installation

	git clone https://github.com/hIMEI29A/leecen && cd leecen && ./install.sh

Usage

Here is a leecen's help message

	leecen - licenses for your project

	Usage: leecen [-h] | COMMAND [ARG] [OPTIONS]
	=====================================================

	Commands:
        header - make license header for given file
        license - make LICENSE file
	Options:
        -h - read this message
        -l - list licenses or headers available
        -e - set custom autor's email
        -n - set custom autor's name
        -y - set custom year
        -f - set custom license file's name

*/
package main
