package goast

// BasicLit
// - Kind
// - Value
type BasicLit struct {
	ID  string `json:"id"`
	Val int
}

var BasicLitInstance = BasicLit{ID: "this is an id", Val: 87}
