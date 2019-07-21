package main

import "capture/module"

func main() {
	server := module.NewServer("tcp", ":49156")
	server.ListenAndHandle()
}