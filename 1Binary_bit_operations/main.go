package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str1 := "1.000101"
	str2 := "1.00001"

	s1 := strings.Split(str1, ".")
	fmt.Println(s1)
	for _, v := range s1 {
		atoi, err := strconv.Atoi(v)
		if err != nil {
			return
		}
		fmt.Println(atoi)

	}

	s2 := strings.Split(str2, ".")
	fmt.Println(s2)
}
