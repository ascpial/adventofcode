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
		for cur := start; cur <= end; cur++ {
			base10 := strconv.Itoa(cur)
		subSizes:
			for k := 1; k <= len(base10)/2; k++ {
				// fmt.Printf("k: %d\n", k)
				if len(base10)%k == 0 && len(base10)/k >= 2 {
					repeat := true
					for l := 0; l < len(base10)/k-1 && repeat; l++ {
						// fmt.Printf("%s; %s\n", base10[l*k:(l+1)*k], base10[(l+1)*k:(l+2)*k])
						repeat = repeat && base10[l*k:(l+1)*k] == base10[(l+1)*k:(l+2)*k]
						// fmt.Printf("is repeating: %t\n", repeat)
					}
					if repeat {
						result = result + cur
						// fmt.Printf("Does repeat: %d\n", cur)
						break subSizes
					}
				}
			}
		}
	}
	fmt.Printf("%d\n", result)
}
