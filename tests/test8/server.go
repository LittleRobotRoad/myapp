package main

import (
	"errors"
	"net/rpc"
	"fmt"
	"os"
	"net"
)

type Args struct {
	A, B int
}
type Math int
type Quotient struct {
	Que, Rem int
}

func main() {
	math := new(Math)
	rpc.Register(math)
	//rpc.HandleHTTP()  //http
	//err := http.ListenAndServe(":1234", nil)  /http
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Println("Fatal error:", err.Error())
		os.Exit(2)
	}
	for {

		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("conn error:", err)
			continue
		}
		rpc.ServeConn(conn)  //tcp
		//jsonrpc.ServeConn(conn)  //json
	}
}

func (m *Math) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}
func (m *Math) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Que = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}