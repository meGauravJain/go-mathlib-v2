# go-mathlib-v2

Extended math utility functions for Go. This is a **private** test repository.

## Functions

- `Power(base, exponent int) int` - Exponentiation
- `GCD(a, b int) int` - Greatest common divisor
- `LCM(a, b int) int` - Least common multiple
- `IsPrime(n int) bool` - Primality test
- `Fibonacci(n int) int` - Nth Fibonacci number

## Usage

```go
import mathlib "github.com/meGauravJain/go-mathlib-v2"

mathlib.Power(2, 10)    // 1024
mathlib.GCD(12, 8)      // 4
mathlib.IsPrime(17)     // true
mathlib.Fibonacci(10)   // 55
```
