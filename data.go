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

const (
	GIT      = "git"
	GET      = "--get"
	CONFIG   = "config"
	UEMAIL   = "user.email"
	UNAME    = "user.name"
	EMPTY    = " "
	HOMEDIR  = "/.leecen"
	HMESS    = " - licenses for your project"
	USAGE    = "leecen [-h | --help] | COMMAND [ARG] [OPTIONS]"
	NOARGS   = "No arguments is provided. Try leecene -h for help"
	FIRSTARG = "Commands must be used. Try leecene -h for help"
	HEXIST   = "License or header does not exist. Try leecene -h for help"
	MISSING  = "Argument is missing"
)

const (
	NOTEXIST = "Given path does not exist"
	EXIST    = "File already exist, we'll not rewrite it"
	NOEMAIL  = "WARNING: Please configure your git. I need your user email"
	NONAME   = "WARNING: Please configure your git. I need your user name"
)

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

/*
var HEADERS = []string{
	"agpl-3.0-header",
	"apache-2.0-header",
	"gpl-2.0-header",
	"gpl-3.0-header",
	"lgpl-2.1-header",
}
*/
