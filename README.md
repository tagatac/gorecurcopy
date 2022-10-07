# gorecurcopy

[![Build Status](https://travis-ci.org/tagatac/gorecurcopy.svg?branch=master)](https://travis-ci.org/tagatac/gorecurcopy)
[![GoDoc](https://godoc.org/github.com/tagatac/gorecurcopy?status.svg)](https://godoc.org/github.com/tagatac/gorecurcopy)
[![Go Report Card](https://goreportcard.com/badge/github.com/tagatac/gorecurcopy)](https://goreportcard.com/report/github.com/tagatac/gorecurcopy)
[![Version](https://img.shields.io/github/tag/tagatac/gorecurcopy)](https://github.com/tagatac/gorecurcopy/releases)



`gorecurcopy` copies directories recursively without external dependencies. Compatible with OSX, Linux, and Windows.

This fork improves upon
[plus3it's version](https://github.com/tagatac/gorecurcopy) by supplying a
mockable interface for easier testing.

Example:

```go
import (
	"github.com/tagatac/gorecurcopy"
)

...
copier := gorecurcopy.NewCopier()
err := copier.CopyDirectory("directory", "new_directory")
```
