package main

import (
	"fmt"
	"os"
	"strings"
)


func check(e error) {
	if e != nil {
		panic(e)
	}
}


func main() {
	data, err := os.ReadFile("./contacts.vcf")
	check(err)


	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		fmt.Printf("-%q\n", line)
	}
}
