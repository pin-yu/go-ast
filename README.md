# Go Abstract Syntax Tree Example

Go abstract syntax tree example

## How to leverage this repo

If one would like to know what `ArrayType` is, she / he can take a look at `array_type.go`. I have written some exlanation in the comments.

Let me take the comments in `array_type.go` as example.
There are two kinds of ArrayType, Array and Slice, in the following snippet. For Array, there are two children node, BasicLit and Ident. For Slice, there are only one child node, Ident.

Note that `-` in the comments means the tree edge.

```go
package goast

// Array

// ArrayType ([10]int)
// - BasicLit(10)
// - Ident(int)
var arr [10]int

// Slice

// ArrayType ([]int)
// - Ident(int)
var slice []int
```
