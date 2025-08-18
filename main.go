package main

import (
	"bufio"
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
	var phone string

	for i := 0; i < len(splitted); i++ {
		if strings.HasPrefix(splitted[i], "TEL;") {
			phoneStringSplitted := strings.Split(string(splitted[i]), ":")
			phone = phoneStringSplitted[1]
			if !strings.HasPrefix(phone, "+") && len(phone) == 13 {
				phoneSplitted := strings.Split(phone, "-")
				phoneStringRestored := phoneStringSplitted[0] + ":" + "+38" + phoneSplitted[0] + phoneSplitted[1] + phoneSplitted[2]
				splitted[i] = phoneStringRestored
			}

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
