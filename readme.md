# Hybrid News Reader

This is a golang exercise for our team to practice the concept of interface and package.

## Build

```sh
git clone git@github.com:vicky-sunshine/hn-reader.git
cd hn-reader
go build .
./hn-reader -h
```

## Usage

```sh
Usage:
  hn-reader [command]

Available Commands:
  help        Help about any command
  hkns        Read hacker news
  rdgl        Read reddit /r/golang

Flags:
  -h, --help         help for hn-reader
  -n, --number int   Specify number of top articles (default 10)

Use "hn-reader [command] --help" for more information about a command.
```
