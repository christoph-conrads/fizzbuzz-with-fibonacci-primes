package main


import (
	"fizzbuzz_prime/fibonacci"
	"testing"
)


func Example() {
	// f1, f2, ... =
	//    1,    1,    2,    3,    5,    8,   13,   21,   34,   55,   89,  144,
	//  233,  377,  610,  987, 1597, 2584, 4181, 6765,
	fizzbuzz(20)
	// Output:
	// 1
	// 1
	// BuzzFizz
	// BuzzFizz
	// BuzzFizz
	// 8
	// BuzzFizz
	// Buzz
	// 34
	// Fizz
	// BuzzFizz
	// Buzz
	// BuzzFizz
	// 377
	// Fizz
	// Buzz
	// BuzzFizz
	// 2584
	// 4181
	// FizzBuzz
}


func BenchmarkFibonacciPrimeTest(b *testing.B) {
	b.StopTimer()
	m := 50
	n0 := 1000

	for i := 0; i < b.N; i++ {
		gen := fibonacci.MakeGenerator()
		for n := 0; n < n0; n++ {
			gen.Execute()
		}

		b.StartTimer()
		for n := n0; n < 2000; n++ {
			f := gen.Execute()
			f.ProbablyPrime(m)
		}
		b.StopTimer()
	}
}
