package main

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}

func GetFibonacci(n int) (interface{}, error) {
	return Fibonacci(n), nil
}
