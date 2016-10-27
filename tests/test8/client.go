package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Args struct {
	A, B int
}
type Quotient struct {
	Que, Rem int
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	//client, err := jsonrpc.Dial("tcp", "localhost:1234")  //json
	//client, err := rpc.DialHTTP("tcp", "localhost:1234") //http
	if err != nil {
		log.Fatal("dialing:", err)
	}
	args := Args{17, 8}
	var reply int
	err = client.Call("Math.Multiply", args, &reply)
	if err != nil {
		log.Fatal("Math error:", err)
	}

	fmt.Printf("Math: %d+%d=%d\n", args.A, args.B, reply)

	var quo Quotient
	err = client.Call("Math.Divide", args, &quo)
	if err != nil {
		log.Fatal("Math error:", err)
	}

	fmt.Printf("Math: %d/%d=%d\n  remaider:%d", args.A, args.B, quo.Que, quo.Rem)
}