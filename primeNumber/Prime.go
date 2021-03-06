package main

import (
  "fmt"
  "math"

)

const N = 100
func main(){
	var x, y, n int
	nsqrt := math.Sqrt(N)

	Isprime := [N]bool{}

	for x = 1; float64(x) <= nsqrt; x++ {
		for y = 1; float64(y) <= nsqrt; y++ {
			n = 4*(x*x) + y*y
			if n <= N && (n%12 == 1 || n%12 == 5) {
				Isprime[n] = !Isprime[n]

			}
			n = 3*(x*x) + y*y
			if n <= N && n%12 == 7 {
				Isprime[n] = !Isprime[n]
			}
			n = 3*(x*x) - y*y
			if x > y && n <= N && n%12 == 11 {
				Isprime[n] = !Isprime[n]
			}
		}
	}

	for n = 5; float64(n) <= nsqrt; n++ {
		if Isprime[n] {
			for y = n * n; y < N; y += n * n {
				Isprime[y] = false
			}
		}
	}

	Isprime[2] = true
	Isprime[3] = true

	primes := make([]int, 0, 1270606)
	for x = 0; x < len(Isprime)-1; x++ {
		if Isprime[x] {
			primes = append(primes, x)
		}
	}

	// primes is now a slice that contains all primes numbers up to N
	// so let's print them
	for _, x := range primes {
		fmt.Println(x)
	}
}