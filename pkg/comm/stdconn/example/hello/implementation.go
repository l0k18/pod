package main

import (
	"fmt"
	"net/rpc"
	"os"
	
	qu "github.com/l0k18/pod/pkg/util/quit"
	
	"github.com/l0k18/pod/pkg/comm/stdconn"
)

type Hello struct {
	Quit qu.C
}

func NewHello() *Hello {
	return &Hello{qu.T()}
}

func (h *Hello) Say(name string, reply *string) (err error) {
	r := "hello " + name
	*reply = r
	return
}

func (h *Hello) Bye(_ int, reply *string) (err error) {
	r := "i hear and obey *dies*"
	*reply = r
	h.Quit.Q()
	return
}

func main() {
	printlnE("starting up example worker")
	hello := NewHello()
	stdConn := stdconn.New(os.Stdin, os.Stdout, hello.Quit)
	err := rpc.Register(hello)
	if err != nil {
		printlnE(err)
		return
	}
	go rpc.ServeConn(stdConn)
	hello.Quit.Wait()
	printlnE("i am dead! x_X")
}

func printlnE(a ...interface{}) {
	out := append([]interface{}{"[Hello]"}, a...)
	_, _ = fmt.Fprintln(os.Stderr, out...)
}
