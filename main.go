package main

func main() {
	println("welcome to heartbreak")
	factorial(5)
}

func factorial(n int) {
	result := 1
	for i := 0; i < n; i++ {
		result *= (i + 1)
		println(result)
	}
}
