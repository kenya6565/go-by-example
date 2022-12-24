package main

import "fmt"

func zeroval(ival int) {
    ival = 0
}

func zeroptr(iptr *int) {
    *iptr = 0
}

func sliceptr(sptr []string) []string {
	// slc := *sptr[0]
	slice := sptr[:]
	
	fmt.Println(slice)

	slice[0] = "hello"
  fmt.Println(slice)
	fmt.Println(sptr)
	return sptr
	
}

func pointer_main() {
    i := 1
    fmt.Println("initial:", i)
		
    zeroval(i)
    fmt.Println("zeroval:", i)
		// zeroval = 1

    zeroptr(&i)
    fmt.Println("zeroptr:", i)
		// zeroptr = 0

    fmt.Println("pointer:", &i)
		// pointer: 0xc0000180d0

		s := []string{"pointer1", "pointer2", "pointer3"}
		var val []string
		val = sliceptr(s)

		fmt.Println(val)

}