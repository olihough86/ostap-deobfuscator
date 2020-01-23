package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
)

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
	var first []int
	var second []int
	var charcode []int
	var output string

	println("Tool to deobfuscate ostap JSE - Oliver Hough (@olihough86)")

	var re = regexp.MustCompile(`(?m)this\['([^']*)']=2100;`)
	tfile := os.Args[1]
	content, err := ioutil.ReadFile(tfile)
	if err != nil {
		log.Fatal(err)
	}
	str := string(content)
	a := re.FindStringSubmatch(str)

	fmt.Println("Found possible array identifier: " + a[1])

	re = regexp.MustCompile(`(?m)\[` + a[1] + `-\(` + a[1] + `\/3\)\]=([^;]+)`)
	b := re.FindAllStringSubmatch(str, -1)
	for i := range b {
		v, _ := strconv.Atoi(b[i][1])
		first = append(first, v)
	}
	re = regexp.MustCompile(`(?m)\[` + a[1] + `\]=([^;]+)`)
	c := re.FindAllStringSubmatch(str, -1)
	for i := range c {
		v, _ := strconv.Atoi(c[i][1])
		second = append(second, v)
	}

	if len(first) != len(second) {
		println("Amount of found values do not match, attempting to decode anyway")
	}

	for i := range first {
		v := first[i] - second[i]
		charcode = append(charcode, v)
	}

	for i := range charcode {
		output = output + string(charcode[i])
	}

	println(output)
}
