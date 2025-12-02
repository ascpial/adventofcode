package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

var example = `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`

//go:embed input.txt
var input string

func main() {
	result := 0
	for curRange := range strings.SplitSeq(strings.TrimSpace(input), ",") {
		ends := strings.Split(curRange, "-")
		if len(ends) != 2 {
			panic("Ends not of length 2")
		}
		start, err1 := strconv.Atoi(ends[0])
		if err1 != nil {
			panic(err1)
		}
		end, err2 := strconv.Atoi(ends[1])
		if err2 != nil {
			panic(err2)
		}
		for i := range end - start + 1 {
			cur := start + i
			base10 := strconv.Itoa(cur)
			if len(base10)%2 == 0 && base10[:len(base10)/2] == base10[len(base10)/2:] {
				result += cur
			}
		}
	}
	fmt.Printf("%d\n", result)
}
