<p align="center"><a href="#readme"><img src="https://gh.kaos.st/branca.svg"/></a></p>

<p align="center">
  <a href="https://kaos.sh/g/branca.v1?docs"><img src="https://gh.kaos.st/godoc.svg" alt="PkgGoDev" /></a>
  <a href="https://kaos.sh/r/branca"><img src="https://kaos.sh/r/branca.svg" alt="GoReportCard" /></a>
  <a href="https://kaos.sh/c/branca"><img src="https://kaos.sh/c/branca.svg" alt="Coverage Status" /></a>
  <a href="https://kaos.sh/b/branca"><img src="https://codebeat.co/badges/eca8a1ed-a16f-4005-a7bc-0d16f8d70ae4" alt="Codebeat badge" /></a>
  <a href="https://kaos.sh/w/branca/ci"><img src="https://kaos.sh/w/branca/ci.svg" alt="GitHub Actions CI Status" /></a>
  <a href="https://kaos.sh/w/branca/codeql"><img src="https://kaos.sh/w/branca/codeql.svg" alt="GitHub Actions CodeQL Status" /></a>
  <a href="#license"><img src="https://gh.kaos.st/mit.svg" /></a>
</p>

<p align="center"><a href="#installation">Installation</a> • <a href="#usage-example">Usage example</a> • <a href="#benchmarks">Benchmarks</a> • <a href="#build-status">Build Status</a> • <a href="#license">License</a></p>

<br/>

`branca.go` is [branca token specification](https://github.com/tuupola/branca-spec) implementation for Golang 1.16+.

Features and benefits:

* Pure Go implementation;
* No third-party dependencies at all;
* 100% code coverage;
* Fuzz tests.

### Installation

Make sure you have a working Go 1.18+ workspace (_[instructions](https://golang.org/doc/install)_), then:


```bash
go get -u github.com/essentialkaos/branca
```

### Usage example

```go
package main

import (
  "fmt"
  
  "github.com/essentialkaos/branca"
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

| Branch | Status |
|--------|----------|
| `master` | [![CI](https://kaos.sh/w/branca/ci.svg?branch=master)](https://kaos.sh/w/branca/ci?query=branch:master) |
| `develop` | [![CI](https://kaos.sh/w/branca/ci.svg?branch=develop)](https://kaos.sh/w/branca/ci?query=branch:develop) |

### License

[MIT](LICENSE)

<p align="center"><a href="https://essentialkaos.com"><img src="https://gh.kaos.st/ekgh.svg"/></a></p>