package module

import (
	"log"
	"net"
)

func NewServer(network, address string) *Server {
	return &Server{
		network:network,
		address:address,
	}
}

type Server struct {
	network string
	address string
}

func (ser *Server) ListenAndHandle() {
	l, err := net.Listen(ser.network, ser.address)
	if err != nil {
		log.Panicln(err)
	}
	defer l.Close()

	log.Println("Server is start!")
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println(err)
		}

		log.Println("Accept a Client!")
		go handleR(conn)
		go handleW(conn)
	}
}

func handleR(conn net.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil{
			log.Println(err)
			return
		}

		log.Printf("Server receive: %s\n",buf[:n])
	}
}

func handleW(conn net.Conn) {
	_, err := conn.Write([]byte("Hello World!"))
	if err != nil {
		log.Println(err)
		return
	}
}