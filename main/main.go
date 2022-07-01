package main

import (
	"fmt"
	"os"
	"encoding/hex"
	"net"
	"errors"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Errorf("%v\n", "Invalid count of arguments")
		return
	}

	targetMAC := os.Args[1]
	byteMac, err := stringToMac(targetMAC)

	if err != nil {
		fmt.Errorf("%v\n", "TargetMAC is invalid: %v", err)
		return
	}

	packet := genMagicPacket(byteMac)

	local,err := net.ResolveUDPAddr("udp4", ":9")
	destinationAddress, err := net.ResolveUDPAddr("udp4", "192.168.0.255:9")
	connection, err := net.DialUDP("udp4", local, destinationAddress)
	defer connection.Close()

	if err != nil {
		fmt.Errorf("%v\n", "Unable to create socket:%v", err)
		return
	}

	_, err = connection.Write(packet)

	if err != nil {
		fmt.Errorf("%v\n", "Unable to write to the socker: %v", err)
		return
	}
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