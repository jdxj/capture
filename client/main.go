package main

import "capture/module"

func main() {
	client := module.NewClient("tcp", "127.0.0.1:49152")
	client.DialAndPlay()
}
