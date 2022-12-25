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

func sliceNonPointer(slc []string)  {
	fmt.Println(slc)

	slc[0] = "hello"
	fmt.Println(slc)


	// return slc

}

func convertString(s string) {
	s = "change"

}

func convertPointerString(s *string) {
	*s = "change"

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
		str := "hoge"
		convertString(str)
		fmt.Println(str)
		convertPointerString(&str)
		fmt.Println(str)
    // sliceptr(s)

		var s2 = s
		s2[0] ="fuga"
    fmt.Println(s2)
		// var nonPointer []string
		sliceNonPointer(s)
		// fmt.Println(nonPointer)
		fmt.Println(s)
		// var val []string
		// val = sliceptr(s)

		// fmt.Println(val)

}