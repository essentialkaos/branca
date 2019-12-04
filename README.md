<p align="center"><a href="#readme"><img src="https://gh.kaos.st/branca.svg"/></a></p>

<p align="center"><a href="#installation">Installation</a> • <a href="#usage-example">Usage example</a> • <a href="#benchmarks">Benchmarks</a> • <a href="#build-status">Build Status</a> • <a href="#license">License</a></p>

<p align="center">
  <a href="https://godoc.org/pkg.re/essentialkaos/branca.v1"><img src="https://godoc.org/pkg.re/essentialkaos/branca.v1?status.svg"></a>
  <a href="https://goreportcard.com/report/github.com/essentialkaos/branca"><img src="https://goreportcard.com/badge/github.com/essentialkaos/branca"></a>
  <a href="https://travis-ci.com/essentialkaos/branca"><img src="https://travis-ci.com/essentialkaos/branca.svg"></a>
  <a href='https://coveralls.io/github/essentialkaos/branca?branch=develop'><img src='https://coveralls.io/repos/github/essentialkaos/branca/badge.svg?branch=develop' alt='Coverage Status' /></a>
  <a href="https://codebeat.co/projects/github-com-essentialkaos-branca-master"><img alt="codebeat badge" src="https://codebeat.co/badges/eca8a1ed-a16f-4005-a7bc-0d16f8d70ae4" /></a>
  <img src="https://gh.kaos.st/mit.svg">
</p>

`branca.go` is [branca token specification](https://github.com/tuupola/branca-spec) implementation for Golang 1.10+.

Features and benefits:

* Pure Go implementation;
* No third-party dependencies at all;
* 100% code coverage;
* Fuzz tests.

### Installation

Before the initial install allows git to use redirects for [pkg.re](https://github.com/essentialkaos/pkgre) service (_reason why you should do this described [here](https://github.com/essentialkaos/pkgre#git-support)_):

```
git config --global http.https://pkg.re.followRedirects true
```

Make sure you have a working Go 1.10+ workspace (_[instructions](https://golang.org/doc/install)_), then:

````
go get pkg.re/essentialkaos/branca.v1
````

For update to latest stable release, do:

```
go get -u pkg.re/essentialkaos/branca.v1
```

### Usage example

```go
package main

import (
  "fmt"
  
  "pkg.re/essentialkaos/branca.v1"
)

func main() {
  key := "abcd1234abcd1234abcd1234abcd1234"
  brc, err := branca.NewBranca([]byte(key))

  if err != nil {
    fmt.Printf("Error: %v\n", err)
    return
  }

  payload := "MySuperSecretData"
  token, err := brc.EncodeToString([]byte(payload))

   if err != nil {
    fmt.Printf("Error: %v\n", err)
    return
  }

  fmt.Printf("Token: %s\n", token)
}

```

### Benchmarks

You can run benchmarks by yourself using `make bench` command.

```
BrancaSuite.BenchmarkBase62Decoding            1000000     1046 ns/op      384 B/op      6 allocs/op
BrancaSuite.BenchmarkBase62Encoding            1000000     1913 ns/op      512 B/op      6 allocs/op
BrancaSuite.BenchmarkBrancaDecoding            5000000      373 ns/op       48 B/op      2 allocs/op
BrancaSuite.BenchmarkBrancaDecodingFromString  1000000     1463 ns/op      432 B/op      8 allocs/op
BrancaSuite.BenchmarkBrancaEncoding            1000000     1677 ns/op      208 B/op      4 allocs/op
BrancaSuite.BenchmarkBrancaEncodingToString     500000     3977 ns/op      720 B/op     10 allocs/op
```

### Build Status

| Branch     | Status |
|------------|--------|
| `master` (_Stable_) | [![Build Status](https://travis-ci.com/essentialkaos/branca.svg?branch=master)](https://travis-ci.com/essentialkaos/branca) |
| `develop` (_Unstable_) | [![Build Status](https://travis-ci.com/essentialkaos/branca.svg?branch=develop)](https://travis-ci.com/essentialkaos/branca) |

### License

[MIT](LICENSE)

<p align="center"><a href="https://essentialkaos.com"><img src="https://gh.kaos.st/ekgh.svg"/></a></p>