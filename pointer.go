package main

import "fmt"

func zeroval(ival int) {
    ival = 0
}

func zeroptr(iptr *int) {
    *iptr = 0
}

func pointer_main() {
    i := 1
    fmt.Println("initial:", i)
		
    zeroval(i)
    fmt.Println("zeroval:", i)
		// zeroval = 1

    zeroptr(&i)
    fmt.Println("zeroptr:", i)
		// zeroval = 0

    fmt.Println("pointer:", &i)
		// pointer: 0xc0000180d0
}