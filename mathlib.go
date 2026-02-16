package mathlib

import "math"

// Power returns base raised to the exponent.
func Power(base, exponent int) int {
	return int(math.Pow(float64(base), float64(exponent)))
}

// GCD returns the greatest common divisor of a and b using Euclid's algorithm.
func GCD(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// LCM returns the least common multiple of a and b.
func LCM(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	return abs(a*b) / GCD(a, b)
}

// IsPrime returns true if n is a prime number.
func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n < 4 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

// Fibonacci returns the nth Fibonacci number (0-indexed).
func Fibonacci(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
