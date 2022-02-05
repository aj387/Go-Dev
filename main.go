package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"

	"rsc.io/quote"
)

func getGOMAXPROCS() int {
	return runtime.GOMAXPROCS(0)
}

func main() {
	fmt.Printf("GOMAXPROCS is %d\n", getGOMAXPROCS())

	runtime.GOMAXPROCS(1)
	go func() {
		fmt.Println("1")
	}()
	go func() {
		fmt.Println("2")
	}()
	go func() {
		fmt.Println("3")
	}()

	fmt.Printf("GOMAXPROCS is %d\n", getGOMAXPROCS())
	fmt.Println("My favorite number is", rand.Intn(10000000000000000))

	time.Sleep(time.Second)

	fmt.Println(quote.Go())

}
