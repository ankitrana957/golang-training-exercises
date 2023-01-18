package fibonaaciseries

type newSlice []int

func (d1 *newSlice) SetSlice(val int) {
	*d1 = append(*d1, val)
}

func (d1 newSlice) GetIndexItem(index int) int {
	return d1[index]
}

func (d1 newSlice) GetSlice(end int) newSlice {
	return d1[:end]
}

func FibonacciSeries(k *newSlice, n int) newSlice {
	var v int = 0
	if len(*k) > n {
		return k.GetSlice(n)
	}
	if len(*k) < n {
		v = len(*k)
	}
	for i := v; i < n; i++ {
		length := len(*k)
		switch length {
		case 0:
			k.SetSlice(0)
		case 1:
			k.SetSlice(1)
		default:
			fibonacciVal := k.GetIndexItem(length-1) + k.GetIndexItem(length-2)
			k.SetSlice(fibonacciVal)
		}
	}
	return *k
}
