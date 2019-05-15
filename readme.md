# Article Reader

This is a golang exercise for our team to practice the concept of interface and package.

## Build

```sh
git clone git@github.com:vicky-sunshine/ariticle-reader.git
cd article-reader
go build .
./artread
```

## Usage

```sh
Usage:
  artread [command]

Available Commands:
  help        Help about any command
  hkns        Read hacker news
  rdgl        Read reddit /r/golang

Flags:
  -h, --help         help for artread
  -n, --number int   Specify number of top articles (default 10)

Use "artread [command] --help" for more information about a command.
```

### Hacker News

```text
$ ./artread hkns
19895310  How I Run a Company with ADHD by askins4trouble 6 hours ago
19893682  Facebook sues analytics firm Rankwave over data misuse by JumpCrisscross 11 hours ago
19895885  The Neko Virtual Machine by azhenley 4 hours ago
19895672  Coin found off northern Australia may be from pre-1400 Africa by curtis 5 hours ago
19886835  How Hot Chicken Happened by jger15 yesterday
19895218  3D Game Shaders for Beginners by lettier 7 hours ago
19894798  F() vs. F(void) in C vs. C++ by headalgorithm 8 hours ago
19893518  Senate Testimony on Privacy Rights and Data Collection in a Digital Economy by aaronbrethorst 12 hours ago
19894673  Awesome Pascal – A curated list of Delphi, FreePascal, and Pascal shiny things by peter_d_sherman 9 hours ago
19893283  Openpilot – open-source self-driving agent by boramalper 12 hours ago
```

### Reddit Golang

```text
$ ./artread rdgl
bnu47l  High Performance DICOM Medical Image Parser in Golang by /u/suyashkumar 9 hours ago
bnvik4  Modulo in Golang by /u/stewi1014 7 hours ago
bnyrzp  Experienced developer but new to Go, what did I do wrong / right in my first package? by /u/itmayziii 2 hours ago
bnmcwk  My first useful open source project by /u/Enapiuz 23 hours ago
bnzcxy  A simple HTTP access logger for golang by /u/timesmaster 1 hours ago
bnqkid  Rove - MySQL database migration tool inspired by Liquibase by /u/motojo 14 hours ago
bnwle4  How to pass keyboard input to a running program? by /u/naturalizedcitizen 5 hours ago
bnvn5m  How do I capture variables at time of an anonymous function's execution? by /u/Spaceface16518 7 hours ago
bnnahw  GitHub - samonzeweb/profilinggo: A quick tour (or reminder) of Go performance tools by /u/samonzeweb 20 hours ago
bncg3z  Go-style concurrency in C by /u/AtomicOrbital yesterday
```

## Reference

- Hacker News API: https://github.com/HackerNews/API
- Reddit Rss: https://www.reddit.com/wiki/rss
- Gofeed: https://github.com/mmcdole/gofeed
  - Struct Definition: https://github.com/mmcdole/gofeed/blob/master/feed.go
- Atom feed format: https://en.wikipedia.org/wiki/Atom_(Web_standard)
