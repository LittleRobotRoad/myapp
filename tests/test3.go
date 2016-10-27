package main

import (
	"fmt"
	"time"
)

func main() {
	//t:=time.Now().Format("Mon Jan _2 15:04:05 2006")
	t :=time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(t)
}
