package main

import (
	"fmt"
	"image"
	"io"
	"math"
	"net"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/hashicorp/yamux"
)

var c, python, java bool

const Pi = 3.14
const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}
func needFloat(x float64) float64 {
	return x * 0.1
}

func test_for() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func test_for2() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func test_for3() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func test_for4() {
	sum := 1
	for {
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

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 100; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(i, z)
	}
	return z
}

func test_case() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Print("%s.\n", os)
	}
}

func test_case2() {
	fmt.Println("When's Satuaday?")
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

func test_case3() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func test_defer() {
	defer fmt.Println("world")
	fmt.Println("hello")
}

func test_defer2() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

func test_pointer() {
	i, j := 42, 2701

	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}

func test_struct() {
	type Vertex struct {
		X int
		Y int
		Z string
	}

	v := Vertex{1, 2, "abcd"}
	v.X = 456
	fmt.Println(v.X)
	fmt.Println(Vertex{1, 2, "abcd"})

	p := &v
	p.X = 1e9
	fmt.Println(v)

	var (
		v1 = Vertex{1, 2, "v1"}
		v2 = Vertex{X: 1}
		v3 = Vertex{}
		pp = &Vertex{1, 2, "vp"}
	)
	fmt.Println(v1, v2, v3, pp)
}

func test_array() {
	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 4, 5, 67, 99}
	fmt.Println(primes)
}
func test_slice() {
	primes := [6]int{2, 3, 4, 5, 67, 99}

	var s []int = primes[1:4]
	fmt.Println(s)
}

func test_slice_literals() {
	q := []int{2, 3, 4, 5, 7, 11, 13}
	q[2] = 50
	fmt.Println(q)

	p := q[:6]
	fmt.Println(p)
}
func test_slice_lin_cap() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	s = s[:0]
	printSlice(s)

	s = s[:4]
	printSlice(s)

	s = s[2:]
	printSlice(s)
}
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func test_nil_slices() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}

func test_make_slices() {
	a := make([]int, 5)
	printSlice(a)

	b := make([]int, 0, 5)
	printSlice(b)

	c := b[:2]
	printSlice(c)

	d := c[2:5]
	printSlice(d)
}

func test_slices_of_slice() {
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"
	board[0][0] = "O"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func test_slice_append() {
	var s []int
	printSlice(s)

	s = append(s, 0)
	printSlice(s)

	s = append(s, 1)
	printSlice(s)

	s = append(s, 2, 3, 4)
	printSlice(s)
}

func test_range() {
	var pow = []int{1, 2, 4, 5, 6}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	pow2 := make([]int, 10)
	for i := range pow2 {
		pow2[i] = 1 << uint(i)
	}
	for _, value := range pow2 {
		fmt.Printf("%d\n", value)
	}
}

func test_map() {
	type Vertex struct {
		Lat, Long float64
	}
	var m map[string]Vertex
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.65988, -74.655,
	}
	fmt.Println(m["Bell Labs"])
}

func test_mutaing_maps() {
	m := make(map[string]int)
	m["answer"] = 42
	fmt.Println("The value:", m["answer"])

	m["answer"] = 48
	fmt.Println("The value:", m["answer"])

	delete(m, "answer")
	fmt.Println("The value:", m["answer"])

	v, ok := m["answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func test_function() {
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

func test_function_closures() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

func fabi_gen() func() int {
	first := 1
	second := 0
	return func() int {
		sum := first + second
		first = second
		second = sum
		return first
	}
}
func test_fabi() {
	fabi := fabi_gen()
	for i := 0; i < 10; i++ {
		fmt.Println(fabi())
	}
}

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func test_method() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

type Abser interface {
	Abs() float64
}

func test_interface() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f
	fmt.Println(a.Abs())

	a = &v
	fmt.Println(a.Abs())

	a = v
	fmt.Println(a.Abs())
}

type I interface {
	M()
}
type T struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)
}
func test_interface_are_satisfied_implicitly() {
	var i I = T{"hello"}
	i.M()
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func test_interface_values() {
	var i I

	i = &T{"hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

type Vetexx struct {
	S string
}
type II interface {
	M()
}

func (t *Vetexx) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}
func test_interface_values_with_nil() {
	var i II
	var t *Vetexx
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()
}

func describe2(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
func test_empty_interface() {
	var i interface{}
	describe2(i)

	i = 42
	describe2(i)

	i = "hello"
	describe2(i)
}

func test_type_assertions() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64)
	fmt.Println(f)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}
func test_type_switches() {
	do(21)
	do("dasdf")
	do(true)
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%v:%v:%v:%v", ip[0], ip[1], ip[2], ip[3])
}

func test_stringer() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod beeblebrox", 9001}
	fmt.Println(a, z)

	var stringer fmt.Stringer
	stringer = a
	fmt.Println(stringer.String())

	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}
func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}
func test_errors() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("%v error", float64(e))
}
func nSqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return 0, nil
}
func test_exercise_errors() {
	fmt.Println(nSqrt(2))
	fmt.Println(nSqrt(-2))
}

func test_reader() {
	r := strings.NewReader("hello, reader!")
	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

type MyReader struct{}

func (e *MyReader) Read(a []byte) (int, error) {
	var size = cap(a)
	for i := 0; i < size; i++ {
		a[i] = 'A'
	}
	//fmt.Println(size)

	return size, nil
}
func test_reader_exercise() {
	var r MyReader
	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v, err = %v b=%v\n", n, err, b)
	}
}

type rot13Reader struct {
	r io.Reader
}

func rot13(b byte) byte {
	switch {
	case 'A' <= b && b <= 'M':
		b += 13
	case 'M' < b && b <= 'Z':
		b -= 13
	case 'a' <= b && b <= 'm':
		b += 13
	case 'm' < b && b <= 'z':
		b -= 13
	}
	return b
}
func (rr *rot13Reader) Read(b []byte) (int, error) {
	n, e := rr.r.Read(b)
	for i := 0; i < n; i++ {
		b[i] = rot13(b[i])
	}
	return n, e
}

func test_exercise_rot_reader() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

func test_image() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
	fmt.Println(m.At(100, 0).RGBA())
}

type Image struct{}

func (e *Image) Bounds() image.Rectangle {
	return image.Rect(100, 0, 0, 0)
}
func test_exercise_images() {
	m := Image{}
	fmt.Println(m.Bounds())
	//pic.ShowImage(m)
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func test_goroutines() {
	go say("world")
	say("hello")
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func test_channels() {
	s := []int{7, 5, 3, 7, 8, 9, 1, 2}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}

func test_buffered_channels() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func test_range_and_close() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

func fibonacci2(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
func test_select() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	go fibonacci2(c, quit)
	time.Sleep(10000 * time.Millisecond)
}

func test_default_selection() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	c.v[key]++
	c.mux.Unlock()
}
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

func test_mutex_counter() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}
	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}

func client() {
	// Get a TCP connection
	conn, err := net.Dial("tcp", "127.0.0.1:9090")
	if err != nil {
		panic(err)
	}

	// Setup client side of yamux
	session, err := yamux.Client(conn, nil)
	if err != nil {
		panic(err)
	}

	//Open a new stream
	stream, err := session.Open()
	if err != nil {
		panic(err)
	}
	//Stream implements net.Conn
	stream.Write([]byte("ping"))
}

func server() {
	listener, err := net.Listen("tcp", "127.0.0.1:9090")
	if err != nil {
		fmt.Printf("listen failed, err: %v\n", err)
		return
	}
	//Accept a TCP connect
	conn, err := listener.Accept()
	if err != nil {
		panic(err)
	}

	//Setup server side of yamux
	session, err := yamux.Server(conn, nil)
	if err != nil {
		panic(err)
	}

	//Accept a stream
	stream, err := session.Accept()
	if err != nil {
		panic(err)
	}
	buf := make([]byte, 4)
	stream.Read(buf)
	fmt.Printf("recv: %s", buf)
}
func test_yamux() {
	go server()
	time.Sleep(time.Second)
	go client()

	time.Sleep(time.Minute)
}
func main() {
	test_yamux()
	return

	fmt.Println("hello %v", 123)
	test_mutex_counter()

	test_default_selection()
	test_select()
	return
	test_range_and_close()
	test_buffered_channels()
	test_channels()
	test_goroutines()
	test_exercise_images()
	test_image()
	test_exercise_rot_reader()
	//test_reader_exercise()
	return
	test_reader()
	test_exercise_errors()
	test_errors()
	test_stringer()
	test_type_switches()
	//test_type_assertions()

	test_empty_interface()
	test_interface_values_with_nil()
	test_interface_values()
	test_interface_are_satisfied_implicitly()
	test_interface()
	test_method()
	return
	test_fabi()
	test_function_closures()
	test_function()
	test_mutaing_maps()
	test_map()
	test_range()
	test_slice_append()
	test_slices_of_slice()
	test_make_slices()
	test_nil_slices()
	test_slice_lin_cap()
	return
	test_slice_literals()
	test_slice()
	test_array()
	test_struct()
	test_pointer()
	test_defer()
	test_defer2()
	test_case3()
	test_case()
	test_case2()
	//Sqrt(200000000);
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
	fmt.Println(sqrt(5), sqrt(-4))
	test_for()
	test_for2()
	test_for3()
	//test_for4()
	//fmt.Println(needInt(Big))
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Big))

	fmt.Printf("hello\n")
	var i int
	fmt.Println(i, c, python, java)

	const World = "xxb"
	fmt.Println("hello", World)
	fmt.Println("happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
