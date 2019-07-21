package main

import "capture/module"

func main() {
	client := module.NewClient("tcp", "114.55.170.124:49156")
	client.DialAndPlay()
}
