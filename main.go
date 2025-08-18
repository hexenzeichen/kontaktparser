package main

import (
	"fmt"
	"os"
	"bufio"
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
		if strings.HasPrefix(splitted[i], "TEL;") {
		    fmt.Println(splitted[i])
		}
		
	}


	file, err := os.Create("./newcontacts.vcf")
	check(err)
	defer file.Close()


	writer := bufio.NewWriter(file)
	for _, line := range splitted {
		_, err = writer.WriteString(line + "\n")
		check(err)
	}
	writer.Flush()

	fmt.Println("Done.")
}
