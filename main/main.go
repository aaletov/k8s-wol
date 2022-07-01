package main

import (
	"fmt"
	"os"
	"encoding/hex"
	"net"
	"errors"
	"bytes"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("%v\n", "Invalid count of arguments")
		return
	}

	targetMAC := os.Args[1]

	udpAddr, err := net.ResolveUDPAddr("udp4", "192.168.0.111:9")
	conn, err := net.ListenUDP("udp4", udpAddr)

	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	defer conn.Close()

	myMac, _ := stringToMac("1C-BF-C0-C0-82-23")

	for {
		var n int

		fmt.Println("Start main loop")

		recvBuf := make([]byte, 258)

		n, _, err = conn.ReadFromUDP(recvBuf)

		if err != nil {
			fmt.Printf("Cannot read from connection: %v\n", err)
			return
		}

		fmt.Printf("Read %v bytes: %v\n", n, recvBuf)

		if bytes.Equal(recvBuf[0 : 6], []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}) && (n == 102) {
			fmt.Println("Found sync sequence")
			isMagicPacket := true

			for i := 6; i < 102; i += 6 {
				isMagicPacket = isMagicPacket && bytes.Equal(recvBuf[i : i + 6], myMac)
			}

			if isMagicPacket {
				sendMagicPacket(targetMAC)
			}
		}
	}


}

func sendMagicPacket(targetMAC string) error {
	byteMac, err := stringToMac(targetMAC)

	if err != nil {
		return err
	}

	packet := genMagicPacket(byteMac)

	local,err := net.ResolveUDPAddr("udp4", ":9")
	destinationAddress, err := net.ResolveUDPAddr("udp4", "192.168.0.255:9")
	conn, err := net.DialUDP("udp4", local, destinationAddress)
	defer conn.Close()

	if err != nil {
		return err
	}

	_, err = conn.Write(packet)

	if err != nil {
		return err
	}

	return nil
}

func stringToMac(stringMac string) ([]byte, error) {
	var (
		byteMac = make([]byte, 6)
		byteBuf []byte
		err error
	)

	if len(stringMac) != 17 {
		return nil, errors.New("MAC length is incorrect")
	}

	for i := 0; i < 6; i += 1 {
		byteBuf, err = hex.DecodeString(string(stringMac[3 * i : 3 * i + 2]))

		if err != nil {
			return nil, err
		}

		if len(byteBuf) != 1 {
			return nil, errors.New("MAC is incorrect")
		}

		byteMac[i] = byteBuf[0]
	}

	return byteMac, nil
}

func genMagicPacket(targetMAC []byte) []byte {
	packet := make([]byte, 102)

	for i := 0; i < 6; i++ {
		packet[i] = 0xFF
	}

	for i := 6; i < 102; i += 6 {
		for j := 0; j < 6; j++ {
			packet[i + j] = targetMAC[j]
		}
	}

	return packet
}