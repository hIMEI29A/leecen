# leecen

**leecen** - License generator for your project. This is Golang port of [licen](https://github.com/lord63/licen).

[![Go Report Card](https://goreportcard.com/badge/github.com/hIMEI29A/leecen)](https://goreportcard.com/report/github.com/hIMEI29A/leecen) [![GoDoc](https://godoc.org/github.com/hIMEI29A/leecen?status.svg)](http://godoc.org/github.com/hIMEI29A/leecen)

## TOC

- [About](#about)
- [Requirements](#requirements)
- [Install](#install)
- [Usage](#usage)

## About

**Leecen** creates file LICENSE in current directory with context of given license. It may also create file with custom name and write license header to file.

## Requirements

`Debian/Ubuntu Linux` with `amd64` arch.

## Install

To install **leecen** do next

```sh
git clone https://github.com/hIMEI29A/leecen && cd leecen && ./install.sh
```
 
## Usage

To get help with **leecen** use option `-h`

```sh
leecen -h
```

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


