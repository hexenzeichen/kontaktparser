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


	splitted := strings.Split(string(data), "\n")

	for i := 0; i < len(splitted); i++ {
		splitted[i] = "subs"
	}

	
	for _, line := range splitted {
		fmt.Printf("-%q\n", line)
	}

}
