package kata

func coloredCells(n int) int64 {
	k := int64(n)
	return k*k + (k-1)*(k-1)
}
