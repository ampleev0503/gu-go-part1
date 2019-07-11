package statistic

//Average func avg of slice
func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

//Sum func sum of slice
func Sum(sl []float64) float64 {
	res := float64(0)
	for _, val := range sl {
		res += val
	}
	return res
}
