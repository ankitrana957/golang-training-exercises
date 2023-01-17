package fibonacciseries

func FibonacciSeries() func(int) []int {
	a := []int{}
	return func(n int) []int {
		var v int = 0
		if len(a) > n {
			return a[:n]
		}
		if len(a) < n {
			v = len(a)
		}
		for i := v; i < n; i++ {
			length := len(a)
			switch length {
			case 0:
				a = append(a, 0)
			case 1:
				a = append(a, 1)
			default:
				fibonacciVal := a[length-1] + a[length-2]
				a = append(a, fibonacciVal)
			}
		}
		return a
	}

}
