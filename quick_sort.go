package main

import "fmt"

func main() {
	n := 0
	fmt.Scanf("%d", &n)
	q := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &q[i])
	}
	quick_sort(q, 0, n-1)
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", q[i])
	}
}

func quick_sort(q []int, l, r int) {
	if l >= r {
		return
	}

	i := l - 1
	j := r + 1
	x := q[(l+r)>>1]

	for i < j {
		for {
			i++
			if q[i] >= x {
				break
			}
		}
		for {
			j--
			if q[j] <= x {
				break
			}
		}
		if i < j {
			q[i], q[j] = q[j], q[i]
		}
	}
	quick_sort(q, l, j)
	quick_sort(q, j+1, r)
}
