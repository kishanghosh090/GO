package main

func changeNumber(num *int) {
	*num = 34
	println(num)
}
func main() {
	num := 23
	changeNumber(&num)
	// pointrs
	println(num)
}
