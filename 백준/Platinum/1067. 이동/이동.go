package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

type base complex128
const PI = math.Pi

func fft(a []base, inv bool) {
	n := len(a)
	j := 0
	roots := make([]base, n/2)
	for i := 1; i < n; i++ {
		bit := n >> 1
		for j >= bit {
			j -= bit
			bit >>= 1
		}
		j += bit
		if i < j {
			a[i], a[j] = a[j], a[i]
		}
	}
	ang := 2 * PI / float64(n)
	if inv {
		ang = -ang
	}
	for i := 0; i < n/2; i++ {
		roots[i] = base(complex(math.Cos(ang*float64(i)), math.Sin(ang*float64(i))))
	}
	for i := 2; i <= n; i <<= 1 {
		step := n / i
		for j := 0; j < n; j += i {
			for k := 0; k < i/2; k++ {
				u := a[j+k]
				v := a[j+k+i/2] * roots[step*k]
				a[j+k] = u + v
				a[j+k+i/2] = u - v
			}
		}
	}
	if inv {
		for i := 0; i < n; i++ {
			a[i] /= base(complex(float64(n), 0))
		}
	}
}

func multiply(v, w []int64) []int64 {
	fv := make([]base, len(v))
	for i, val := range v {
		fv[i] = base(complex(float64(val), 0))
	}
	fw := make([]base, len(w))
	for i, val := range w {
		fw[i] = base(complex(float64(val), 0))
	}

	n := 1
	for n < len(v)+len(w) {
		n <<= 1
	}
	fv = append(fv, make([]base, n-len(fv))...)
	fw = append(fw, make([]base, n-len(fw))...)

	fft(fv, false)
	fft(fw, false)
	for i := 0; i < n; i++ {
		fv[i] *= fw[i]
	}
	fft(fv, true)

	ret := make([]int64, n)
	for i := 0; i < n; i++ {
		ret[i] = int64(math.Round(real(fv[i])))
	}
	return ret
}

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int64, n*2)
	b := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
		a[i+n] = a[i]
	}
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &b[n-i])
	}

	result := multiply(a, b)
	ans := int64(0)
	for _, val := range result {
		if val > ans {
			ans = val
		}
	}
	fmt.Fprintln(writer, ans)
}
