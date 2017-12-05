package main

import(
	"fmt"
	"time"
)

const (
	Male = 1
	Female = 2
)

func main() {
	now := time.Now()
	second := now.Unix()
	if second % Female == 0 {
		fmt.Println("%d female", second)
	} else {
		fmt.Println("%d man", second)
	}

}