package main

func counter() func() int {
	var count int = 0
	return func() int {
		count += 1
		return count
	}
}

func main() {
	inr := counter()

	println(inr())
}
