package module

import (
	"log"
	"net"
	"sync"
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

		wg := &sync.WaitGroup{}
		wg.Add(2)

		go handleClose(wg, conn)
		go handleR(wg, conn)
		go handleW(wg, conn)
	}
}

func handleClose(wg *sync.WaitGroup, conn net.Conn) {
	wg.Wait()
	conn.Close()
}

func handleR(wg *sync.WaitGroup, conn net.Conn) {
	defer wg.Done()

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

func handleW(wg *sync.WaitGroup, conn net.Conn) {
	defer wg.Done()

	_, err := conn.Write([]byte("Hello World!"))
	if err != nil {
		log.Println(err)
		return
	}
}