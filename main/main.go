package main

import (
	"fmt"
	"os"
	"encoding/hex"
	"net"
)

func main() {
	targetMAC := os.Args[1]

	fmt.Println(targetMAC)

	byteMac, _ := stringToMac(targetMAC)
	fmt.Println(hex.EncodeToString(byteMac))

	packet := genMagicPacket(byteMac)
	fmt.Println(hex.EncodeToString(packet))

	conn, err := net.Dial("udp", "127.0.0.1:1234")

	if err != nil {
		fmt.Printf("Some error %v", err)
		return
	}

	_, err = fmt.Fprintf(conn, hex.EncodeToString(packet))

	if err != nil {
		fmt.Printf("Some error %v", err)
		return
	}
}

func stringToMac(stringMac string) ([]byte, error) {
	var (
		byteMac = make([]byte, 6)
		byteBuf []byte
		err error
	)

	for i := 0; i < 6; i += 1 {
		byteBuf, err = hex.DecodeString(string(stringMac[3 * i : 3 * i + 2]))

		if err != nil {
			return nil, err
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