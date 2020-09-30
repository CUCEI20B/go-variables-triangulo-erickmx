package main

import "fmt"

func main()  {
	var base int32
	var height int32

	fmt.Scan(&base)
	fmt.Scan(&height)
	fmt.Print((base*height)/2)
}
