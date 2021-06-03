package rpcDemo

import (
	"log"
	"net"
	"net/rpc"
)

type PersonService struct {
	Name string `json:"name"`
}

func (p *PersonService) Hello(arg string, reply *string) error {
	*reply = "hello," + arg
	return nil
}

func Register() {
	err := rpc.RegisterName("PersonService", new(PersonService))
	if err != nil {
		log.Fatal("register rpc ", err)
	}
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen port error ", err)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("accpet error ", err)
	}

	rpc.ServeConn(conn)
}
