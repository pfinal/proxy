package goproxy

import (
	"bufio"
	"log"
	"net"
)

type Server struct {
	listener net.Listener
	addr     string
}

func NewServer(Addr string) *Server {
	return &Server{
		addr: Addr,}

}

func (s *Server) Start() {
	var err error
	s.listener, err = net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("listening", s.addr)

	for {
		conn, err := s.listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go s.newConn(conn).serve()
	}
}

func (s *Server) newConn(rwc net.Conn) *conn {
	return &conn{
		server: s,
		rwc:    rwc,
		brc:    bufio.NewReader(rwc),
	}
}
