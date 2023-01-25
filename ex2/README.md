# Exercise 2 - Shred tool in Go

## Description
Implementation of a `Shred(path)` function that will overwrite the
given file (e.g. “randomfile”) 3 times with random data and delete the
file afterwards. Note that the file may contain any type of data.

## Usage
Run `go build Shred`. This generates a binary `Shred` that shreds a
file with the path passed to it as a command line argument.
