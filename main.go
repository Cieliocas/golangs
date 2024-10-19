package main

import "fmt"
import "time"

func contador(count int) {
	for i := range count {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func main() {
	contador(5)
}