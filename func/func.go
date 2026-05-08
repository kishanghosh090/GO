package main

func main() {
	// val := fact(3)
	// println(val)

	// lang1, lang2, _ := getLangs()
	// println(lang1, lang2, )

	fn := func(el int) int {
		return el + 100
	}
	println(process(fn))

	fun := processIt()

	println(fun(10), fun)
}

// return a fun
func processIt() func(a int) int {
	return func(a int) int {
		return a * a * a
	}
}

// callback fun or higher order func
func process(fn func(a int) int) int {
	return fn(23) * fn(23)
}

func getLangs() (string, string, int) {
	return "golang", "js", 1
}
func fact(a int) int {
	if a == 0 {
		return 1
	}
	return a * fact(a-1)
}
