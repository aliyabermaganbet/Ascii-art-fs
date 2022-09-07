package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"fs/initial"
)

func main() {
	if len(os.Args) == 3 {
		th := os.Args[2]
		t := th + ".txt"
		if initial.Check_for_hash(t) {
			file, err := os.Open(t)
			if err != nil {
				log.Fatal(err)
			}
			r := bufio.NewScanner(file)
			var array []string
			for r.Scan() {
				array = append(array, r.Text())
			}
			arg := os.Args[1]
			d := initial.Check_for_letters(th, arg)
			if d != nil {
				fmt.Println(d.Error())
				return
			}
			initial.Create_the_map(array, arg)
		} else {
			fmt.Println("Usage: go run . [STRING] [BANNER]\nEX: go run . something standard")
		}
	} else if len(os.Args) == 2 {
		f := "standard"
		arg := os.Args[1]
		d := initial.Check_for_letters(f, arg)
		if d != nil {
			fmt.Println(d.Error())
			return
		}
		file, err := os.Open(f + ".txt")
		if err != nil {
			log.Fatal(err)
		}
		r := bufio.NewScanner(file)
		var array []string
		for r.Scan() {
			array = append(array, r.Text())
		}
		initial.Create_the_map(array, arg)
	}
}
