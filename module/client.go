package module

import (
	"io"
	"log"
	"net"
	"sync"
	"time"
)

func NewClient(network, address string) *Client {
	return &Client{
		network:network,
		address:address,
	}
}

type Client struct {
	network string
	address string

	wg sync.WaitGroup
}

func (cli *Client) DialAndPlay() {
	conn, err := net.Dial(cli.network, cli.address)
	if err != nil {
		log.Panicln(err)
	}
	log.Println("Dial is success!")

	cli.wg.Add(2)

	go cli.playR(conn)
	go cli.playW(conn)

	cli.wg.Wait()
}

func (cli *Client)playR(conn net.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil && err != io.EOF {
			log.Println(err)
			break
		}

		log.Printf("Client receive: %s\n",buf[:n])
	}

	cli.wg.Done()
}

func(cli *Client) playW(conn net.Conn) {
	select {
	case <-time.After(time.Hour):
		cli.wg.Done()
	}
}