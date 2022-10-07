# gorecurcopy

[![Build Status](https://travis-ci.org/tagatac/gorecurcopy.svg?branch=master)](https://travis-ci.org/tagatac/gorecurcopy)
[![GoDoc](https://godoc.org/github.com/tagatac/gorecurcopy?status.svg)](https://godoc.org/github.com/tagatac/gorecurcopy)
[![Go Report Card](https://goreportcard.com/badge/github.com/tagatac/gorecurcopy)](https://goreportcard.com/report/github.com/tagatac/gorecurcopy)
[![Version](https://img.shields.io/github/tag/tagatac/gorecurcopy)](https://github.com/tagatac/gorecurcopy/releases)



`gorecurcopy` copies directories recursively without external dependencies. Compatible with OSX, Linux, and Windows.

This fork improves upon
[plus3it's version](https://github.com/tagatac/gorecurcopy) by supporting the
[afero Fs interface](https://pkg.go.dev/github.com/spf13/afero?utm_source=godoc#Fs).

Example:

```go
import (
	"github.com/spf13/afero"
	"github.com/tagatac/gorecurcopy"
)

...
copier := gorecurcopy.NewCopier(afero.NewOsFs())
err := copier.CopyDirectory("directory", "new_directory")
```
