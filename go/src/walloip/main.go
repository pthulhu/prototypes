package main

import (
	"fmt"
	"net"
	"os"
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

		wallApp := "wall"
		echoApp := "echo"

		wallCmd := exec.Command(wallApp)
		echoCmd := exec.Command(echoApp, messageContent)

		wallCmd.Stdin, _ = echoCmd.StdoutPipe()
		wallCmd.Stdout = os.Stdout

		echoCmd.Start()
		wallCmd.Run()
		echoCmd.Wait()
	}
}
