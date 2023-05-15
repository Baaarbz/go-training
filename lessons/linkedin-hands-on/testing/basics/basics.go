package main

type numeric interface {
	~int | ~float64
}

// generic sum sums a slice of numbers
func sum[T numeric](numbers ...T) T {
	var s T
	for _, n := range numbers {
		s += n
	}
	return s
}
