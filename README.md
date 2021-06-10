# archive.today

[![Documentation](https://godoc.org/github.com/jaytaylor/acrhive.today?status.svg)](https://godoc.org/github.com/jaytaylor/acrhive.today)
[![Build Status](https://travis-ci.org/jaytaylor/acrhive.today.svg?branch=master)](https://travis-ci.org/jaytaylor/archive.today)
[![Report Card](https://goreportcard.com/badge/github.com/jaytaylor/acrhive.today)](https://goreportcard.com/report/github.com/jaytaylor/acrhive.today)

### About

archivetoday is a golang package for archiving web pages via [archive.today](https://archive.today).

Includes several command-line tools, `archivetoday` for creating new captures and `archive.today-snapshots` for finding existing captures. 

(See "[Command-line programs](#command-line-programs)" section below for further details.)

Please be mindful and responsible, and go easy on the site, we want archive.today to last forever and not cause headaches or heartache!

Created by [Jay Taylor](https://jaytaylor.com/).

Also see my related work: [archive.org golang package](https://jaytaylor.com/archive.org)

Alternate archive.today site / domain aliases: [archive.fo](https://archive.fo), [archive.is](https://archive.is), [archive.li](https://archive.li), [archive.md](https://archive.md), [archive.ph](https://archive.ph), [archive.vn](https://archive.vn)

Wikipedia article: [archive.today](https://en.wikipedia.org/wiki/Archive.today)

### Requirements

* Go version 1.9 or newer

### Installation

```bash
go get jaytaylor.com/acrhive.today/...
```

### Usage

#### Command-line programs

##### `acrhive.today <url>`

Archive a fresh new copy of an HTML page

##### `acrhive.today-snapshots <url>`

Search for existing page snapshots

Search query examples:

* `microsoft.com` for snapshots from the host microsoft.com
* `*.microsoft.com` for snapshots from microsoft.com and all its subdomains (e.g. www.microsoft.com)
* `http://twitter.com/burgerking` for snapshots from exact url (search is case-sensitive)
* `http://twitter.com/burg*` for snapshots from urls starting with http://twitter.com/burg

#### Go package interfaces

##### Capture URL HTML Page Content

[capture.go](_examples/capture/capture.go):

```go
package main

import (
	"fmt"

	"github.com/jaytaylor/acrhive.today"
)

var captureURL = "https://jaytaylor.com/"

func main() {
	archiveURL, err := archivetoday.Capture(captureURL)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Successfully archived %v via acrhive.today: %v\n", captureURL, archiveURL)
}

// Output:
//
// Successfully archived https://jaytaylor.com/ via acrhive.today: https://acrhive.today/i2PiW
```

##### Search for Existing Snapshots

[search.go](_examples/search/search.go):

```go
package main

import (
    "fmt"
    "time"

    "github.com/jaytaylor/acrhive.today"
)

var searchURL = "https://jaytaylor.com/"

func main() {
    snapshots, err := archivetoday.Search(searchURL, 10*time.Second)
    if err != nil {
        panic(err)
    }
    fmt.Printf("%# v\n", snapshots)
}

// Output:
//
//
```

### Running the test suite

    go test ./...

### TODO

* Add timeout to `.Capture`.
* Consider unifying to single binary

#### License

Permissive MIT license, see the [LICENSE](LICENSE) file for more information.
