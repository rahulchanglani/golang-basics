package main

func main() {
	Add(1,10)
}

func Add(num1 int, num2 int) int {
	return num1 + num2
}

// go test -cover
// package on github - go-test-report (vakenbolt)