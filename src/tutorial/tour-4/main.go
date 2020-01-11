package main

import (
	"fmt"
	"math"
	"strings"
)

func pointers() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

type Vertex struct {
	X int
	Y int
}

func structFields() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}

func structPointer() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}

func structLiterals() {
	var (
		v1 = Vertex{1, 2}  // has type Vertex
		v2 = Vertex{X: 1}  // Y:0 is implicit
		v3 = Vertex{}      // X:0 and Y:0
		p  = &Vertex{1, 2} // has type *Vertex
	)

	fmt.Println(v1, p, v2, v3)
}

func array() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

func slices() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
}

func slicesPointers() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	/**
	 * スライスはどんなデータも格納しておらず、単に元の配列の部分列を指し示しています。
	 * スライスの要素を変更すると、その元となる配列の対応する要素が変更されます。
	 * 同じ元となる配列を共有している他のスライスは、それらの変更が反映されます。
	 */
	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

func sliceLiterals() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

func sliceBounds() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}

func sliceLenCap() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func nilSlices() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}

func printMakingSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func maikingSlices() {
	a := make([]int, 5)
	printMakingSlice("a", a)

	b := make([]int, 0, 5)
	printMakingSlice("b", b)

	c := b[:2]
	printMakingSlice("c", c)

	d := c[2:5]
	printMakingSlice("d", d)
}

func slicesOfSlice() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func appendPrintSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func appendSlice() {
	var s []int
	appendPrintSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	appendPrintSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	appendPrintSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	appendPrintSlice(s)
}

func rangeGo() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

func rangeContinued() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

// func Pic(dx, dy int) [][]uint8 {
// 	return nil
// }

type MapVertex struct {
	Lat, Long float64
}

func mapGo() {
	var m map[string]MapVertex
	m = make(map[string]MapVertex)
	m["Bell Labs"] = MapVertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}

func mapLiterals() {

	var m = map[string]MapVertex{
		// MapVertex は省略も可能
		"Bell Labs": MapVertex{
			40.68433, -74.39967,
		},
		"Google": MapVertex{
			37.42202, -122.08408,
		},
	}
	fmt.Println(m)
}

func mutatingMaps() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func functionValues() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func functionClosures() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

type methodVertex struct {
	X, Y float64
}

func (v methodVertex) methodsAbs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func funcMethods() {
	v := methodVertex{3, 4}
	fmt.Println(v.methodsAbs())
}

type continueFloat float64

func (f continueFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func funcMethodContinue() {
	f := continueFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

type methodsPointerVertex struct {
	X, Y float64
}

func methodsPointerAbs(v methodsPointerVertex) float64 {
	return math.Sqrt(v.X * v.X * v.Y * v.Y)
}

func methodPointerScale(v *methodsPointerVertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func funcMethodsPointerExplained() {
	mpv := methodsPointerVertex{3, 4}
	methodPointerScale(&mpv, 10)
	fmt.Println(methodsPointerAbs(mpv))
}

type IndirectionVertex struct {
	X, Y float64
}

func (v *IndirectionVertex) IndirectionScale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func IndirectionScaleFunc(v *IndirectionVertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func indirection() {
	v := IndirectionVertex{3, 4}
	v.IndirectionScale(2)
	IndirectionScaleFunc(&v, 10)

	p := &IndirectionVertex{4, 3}
	p.IndirectionScale(3)
	IndirectionScaleFunc(p, 8)

	fmt.Println(v, p)
}

type InterfaceAbser interface {
	InterfaceAbs() float64
}
type InterfaceMyFloat float64

func (f InterfaceMyFloat) InterfaceAbs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *IndirectionVertex) InterfaceAbs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func Interface() {
	var a InterfaceAbser
	f := InterfaceMyFloat(-math.Sqrt2)
	v := IndirectionVertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	// a = v

	fmt.Println(a.InterfaceAbs())
}

func main() {
	// pointers()

	// fmt.Println(Vertex{1, 2})

	// structFields()
	// structPointer()
	// structLiterals()

	// array()
	// slices()
	// slicesPointers()
	// sliceLiterals()
	// sliceBounds()
	// sliceLenCap()
	// nilSlices()
	// maikingSlices()
	// slicesOfSlice()
	// appendSlice()
	// rangeGo()

	// pic.Show(Pic)

	// mapGo()
	// mapLiterals()
	// mutatingMaps()

	functionValues()
	functionClosures()
	funcMethods()
	funcMethodContinue()
	funcMethodsPointerExplained()

	indirection()
}
