<p align="center"><a href="#readme"><img src="https://gh.kaos.st/branca.svg"/></a></p>

<p align="center"><a href="#installation">Installation</a> • <a href="#usage-example">Usage example</a> • <a href="#benchmarks">Benchmarks</a> • <a href="#build-status">Build Status</a> • <a href="#license">License</a></p>

<p align="center">
  <a href="https://godoc.org/pkg.re/essentialkaos/branca.v1"><img src="https://godoc.org/pkg.re/essentialkaos/branca.v1?status.svg"></a>
  <a href="https://goreportcard.com/report/github.com/essentialkaos/branca"><img src="https://goreportcard.com/badge/github.com/essentialkaos/branca"></a>
  <a href="https://travis-ci.org/essentialkaos/branca"><img src="https://travis-ci.org/essentialkaos/branca.svg"></a>
  <a href="https://codebeat.co/projects/github-com-essentialkaos-branca-master"><img alt="codebeat badge" src="https://codebeat.co/badges/eca8a1ed-a16f-4005-a7bc-0d16f8d70ae4" /></a>
  <a href="https://essentialkaos.com/ekol"><img src="https://gh.kaos.st/ekol.svg"></a>
</p>

`branca.go` is Pure-Go [branca token specification](https://github.com/tuupola/branca-spec) implementation.

### Installation

Before the initial install allows git to use redirects for [pkg.re](https://github.com/essentialkaos/pkgre) service (_reason why you should do this described [here](https://github.com/essentialkaos/pkgre#git-support)_):

```
git config --global http.https://pkg.re.followRedirects true
```

Make sure you have a working Go 1.9+ workspace (_[instructions](https://golang.org/doc/install)_), then:

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
PASS: branca_test.go:269: BrancaSuite.BenchmarkBase62Decoding             500000    3855 ns/op
PASS: branca_test.go:261: BrancaSuite.BenchmarkBase62Encoding             500000    5968 ns/op
PASS: branca_test.go:229: BrancaSuite.BenchmarkBrancaDecoding            5000000     371 ns/op
PASS: branca_test.go:252: BrancaSuite.BenchmarkBrancaDecodingFromString   500000    4294 ns/op
PASS: branca_test.go:220: BrancaSuite.BenchmarkBrancaEncoding            1000000    1695 ns/op
PASS: branca_test.go:243: BrancaSuite.BenchmarkBrancaEncodingToString     200000    7968 ns/op
```

### Build Status

| Branch     | Status |
|------------|--------|
| `master` (_Stable_) | [![Build Status](https://travis-ci.org/essentialkaos/branca.svg?branch=master)](https://travis-ci.org/essentialkaos/branca) |
| `develop` (_Unstable_) | [![Build Status](https://travis-ci.org/essentialkaos/branca.svg?branch=develop)](https://travis-ci.org/essentialkaos/branca) |

### License

[EKOL](https://essentialkaos.com/ekol)

<p align="center"><a href="https://essentialkaos.com"><img src="https://gh.kaos.st/ekgh.svg"/></a></p>