# [garf](http://open.backenderia.com/garf/) [![Build Status](https://img.shields.io/travis/backenderia/garf.svg?style=flat-square)](https://travis-ci.org/backenderia/garf)

A (one-code-to-rule-them-all) **G**olang **A**PI (**R**estful) **F**ramework and factory.

## Use Cases

## Features / Uses

- Generate a simple RESTful API using command line
- Same code to run on any supported framework
- Rich collection of supported middlewares and helpers
- Easy to create bundles or import existing one
- Customizable bootstrap templates

## Supported web frameworks

- [echo](http://github.com/labstack/echo)
* (not yet) [gin-gonic](#)
* (not yet) [gorilla mux + negroni](#)

## Supported middleware collection

- [Garf collection](http://github.com/backenderia/garf-contrib)
- [Echo collection](#)
- [Gin-gonic collection](#)
- [Negroni collection](#)

## Installation

```shell
$ go install github.com/backenderia/garf
```

## Command Usage

```shell
$ garf help
```

## Documentation

- [Server & Context API](http://godoc.org/github.com/backenderia/garf/server)
- [Registry API](http://godoc.org/github.com/backenderia/garf/registry)

## Guides

- [Creating RESTful API](#)
- [Creating multiple bundles on your API](#)
- [Using other frameworks middlewares](#)

## Stable milestone

- [ ] Echo support
- [ ] Gin support
- [ ] Mux + Negroni support
- [ ] Complete test coverage
- [ ] Complete management using command line tool

## Contribute

- Reporting problems
- Sending suggestions
- Submiting middlewares to our [collection](http://github.com/backenderia/garf-contrib)
- Creating guides and examples

Feel free to submit using issues.
