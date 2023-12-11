package goast

func AssignStmt(input int) (int, int) {
	// AssignStmt
	// - Lhs
	// - Token
	// - Rhs
	a, b := 10, input

	var c int
	var d int
	c, d = a, b

	return c, d
}
