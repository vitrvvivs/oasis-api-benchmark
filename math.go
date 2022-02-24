package main

func sum(arr []int64) (r int64) {
	for _, el := range arr {
		r += el
	}
	return
}

func min(arr []int64) (r int64) {
	r = 0x7FFFFFFFFFFFFFFF
	for _, el := range arr {
		if el < r {
			r = el
		}
	}
	return
}

func max(arr []int64) (r int64) {
	r = -0x8000000000000000
	for _, el := range arr {
		if el > r {
			r = el
		}
	}
	return
}
