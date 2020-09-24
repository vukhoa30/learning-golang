package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func splitInt(sum int) (x, y int) {
	x = sum / 10
	y = sum % 10
	return
}

type Vertex struct {
	X int
	Y int
}

// fibonacci is a function that returns
// a function that returns an int.
// This concept is called "function closure"
func fibonacci() func() int {
	prevOfPrev := 0
	prev := 1
	count := -1
	return func() int {
		count++
		if count < 2 {
			return count
		}
		temp := prevOfPrev
		prevOfPrev = prev
		prev = prev + temp
		return prev
	}
}

func main() {
	s1, s2 := swap("1", "2")    // => multiple outputs; => ":=" is "var"
	fmt.Println(s1, s2)         // 2 1
	fmt.Println(swap("1", "2")) // 2 1 => multiple outputs for multiple params

	var n int // => can be string, int, bool, float32, float64, byte
	n = 17
	fmt.Println(splitInt(n)) // 1 7 => naked return

	fmt.Printf("v is of type %T\n", 33) // int

	for i := 0; i < 1; i++ {
		fmt.Println("for called")
	}
	for n > 17 {
		n++
	} // => this is golang's "while"
	for {
		fmt.Println("infinite loop!")
		break
	}

	if n < n+1 {
		if temp := 5; temp < 6 {
			fmt.Println("if called!", temp)
		} else {
			fmt.Println("not gonna happen", temp)
		}
	}

	switch temp := 2; temp {
	case 1:
		fmt.Println("temp is 1")
	case 2:
	case 3:
		fmt.Println("temp is 2 or 3")
	default:
		fmt.Println("temp is sth else that I don't care")
	} // => break is automatically called

	switch {
	case true:
		fmt.Println("switch with no initial condition")
	case false:
		fmt.Println("tsk tsk")
	case 1 == 1:
		fmt.Println("tsk tsk 2")
	}

	for i := 1; i <= 2; i++ {
		defer fmt.Println("deferred:", i) // => stacking defers called in the end of function
	}

	// pointer
	v := 6
	ptr := &v
	fmt.Println(ptr)  // memory address
	fmt.Println(&v)   // same address
	fmt.Println(*ptr) // 6
	*ptr = 7
	fmt.Println(ptr)  // same address
	fmt.Println(*ptr) // 7
	fmt.Println(v)    // 7

	// struct
	v1 := Vertex{1, 2} // type Vertex
	v1.Y = 3
	vPtr := &v1 // type *Vertex
	vPtr.X = 4
	fmt.Println(v1.X)         // 4
	fmt.Println(v1)           // {4 3}
	fmt.Println(Vertex{X: 1}) // {1 0}

	// array
	var a [5]string
	a[0] = "meep"
	a[4] = "ayo"
	fmt.Println(a) // [meep    ayo]
	primes := [6]int{2, 3, 5, 7, 11, 13}
	subArr := primes[1:4] // the start "1" or end "4" is optional
	fmt.Println(subArr)   // [3 5 7] => slice
	subArr[0] = 9
	fmt.Println(primes[1]) // 9 => subArr uses the same same address
	customArr := []struct {
		i int
		s string
	}{
		{1, "ey"},
		{4, "froy"}, // this "," is important
	}
	fmt.Println(customArr)                        // [{1, "ey"} {4, "froy"}]
	fmt.Println(len(primes), cap(primes))         // 6 6
	fmt.Println(len(primes[:1]), cap(primes[:1])) // 1 6 => primes[:1] is called "slice" and it has a UNDERLYING array "primes"
	fmt.Println(len(subArr), cap(subArr))         // 3 5 => subArr is primes' slice, it starts from 1 so its CAPacity is 5
	newSlice := append(primes[:], 17)
	fmt.Println(newSlice, len(newSlice), cap(newSlice)) // [2 9 5 7 11 13 17] 7 12 => why 12?, shouldn't it be 7 as len?
	for i, v := range customArr {                       // i can be "_" if unsed
		fmt.Println(i, v.s)
	} // 0 ey\n1 froy => similar to nodejs Array.map()
	fmt.Println(make([]int, 5)) // [0 0 0 0 0]

	// map
	m := make(map[string]Vertex)
	m["District 9"] = Vertex{13, 15}
	fmt.Println(m["District 9"]) // {13 15}
	m = map[string]Vertex{
		"District 9": {13, 15}, // again, "," is important here
	} // => same as above
	elem, isExisting := m["District 9"]
	fmt.Println(elem, isExisting) // {13, 15} true
	delete(m, "District 9")
	fmt.Println(m["District 9"]) // {0 0}

	// functions can be used as variables like nodejs
	funcA := func(input int) int {
		return input * 2
	}
	funcB := func(fn func(int) int, input int) int {
		return fn(input) + 3
	}
	fmt.Println(funcB(funcA, 5)) // 13

	// function closure
	f := fibonacci()
	for i := 0; i < 5; i++ {
		fmt.Println(f())
	} // 0 1 1 2 3

	// To be continue: https://tour.golang.org/methods/1
}
