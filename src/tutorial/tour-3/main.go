package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func forStatement() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func forContinued() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func whileStatement() {
	// C言語などにある while は、Goでは for だけを使います。
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func ifStatement() {
	fmt.Println(sqrt(2), sqrt(-4))
}

func pow(x, n, lim float64) float64 {
	// v は scope内でのみ有効
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func ifStatementWithScope() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

func ifElsePow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func ifElse() {
	fmt.Println(
		ifElsePow(3, 2, 10),
		ifElsePow(3, 3, 20),
	)
}

func mySqrt(x float64) float64 {
	z := 1.0
	prevVal := 0.0
	counter := 0
	for z < 10 {
		z -= (z*z - x) / (2 * z)
		fmt.Println("z = ", z)
		if z == prevVal {
			counter++
		}
		if counter >= 2 {
			break
		}
		prevVal = z
	}
	return z
}

func switchStatement() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
}

func switchEvaluationOrder() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

func switchWithNoCondition() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening.")
	}
}

func deferStatement() {
	defer fmt.Println("world")

	fmt.Println("hello")
}

func stackingDefers() {
	fmt.Println("counting...")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

func main() {
	// forStatement()
	// forContinued()
	// whileStatement()
	// ifStatement()
	// ifStatementWithScope()
	mySqrt(10000)
	switchStatement()
	switchEvaluationOrder()
	deferStatement()
	stackingDefers()
}
