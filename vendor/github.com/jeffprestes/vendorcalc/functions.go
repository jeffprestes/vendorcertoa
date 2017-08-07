package vendorcalc

//Delta calcula o delta para funcao de bascara
func Delta(a int, b int, c int) int {
	return b ^ 2 - (4 * a * c)
}
