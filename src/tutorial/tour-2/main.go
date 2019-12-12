package main

import (
	"fmt"
	"math"
	m "math"
	"math/cmplx"
	r "math/rand"
)

func imports() {
	fmt.Println("Now you have %g problems", m.Sqrt(7))
}

func rand() {
	fmt.Println("My favorite number is", r.Intn(10))
}

func pi() {
	fmt.Println(m.Pi)
}

/**
 * 関数の２つ以上の引数が同じ型である場合は、最後の型を残して省略可能
 * 以下は下記と一緒
 * (x, y int) => (x int, y int)
 */
func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

/**
 * Goでの戻り値となる変数に名前をつける( named return value )ことができます。
 * 戻り値に名前をつけると、関数の最初で定義した変数名として扱われます。
 */
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// 型変換
func typeConversion() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}

// 定数
func constants() {
	const Pi = 3.14
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}

// 数値の定数
const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println("===== tutorial 2 =====")
	rand()
	imports()
	pi()
	fmt.Println(add(42, 13))
	var a, b string = swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(17))

	// 型指定
	var c, python bool
	var java string
	var i int
	fmt.Println(i, c, python, java)

	i = 1
	// 初期値
	var j int = 2
	c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)

	// 暗黙的な型宣言
	k := 3
	fmt.Println(i, j, k, c, python, java)

	var (
		ToBe   bool       = false
		MaxInt uint64     = 1<<64 - 1
		z      complex128 = cmplx.Sqrt(-5 + 12i)
	)
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// 変数に初期値を与えずに宣言すると、ゼロ値( zero value )が与えられる
	var _i int
	var _f float64
	var _b bool
	var _s string
	fmt.Printf("%v %v %v %q\n", _i, _f, _b, _s)

	typeConversion()
	constants()

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
