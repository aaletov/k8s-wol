package server

import (
	"fmt"
	"net"
	"context"
	"github.com/aaletov/wol-translator/api/wtapi"
	"github.com/aaletov/wol-translator/utils"
	"google.golang.org/protobuf/types/known/emptypb"
)

// const (
// 	MagicPacketLen = 102
// )

// type Listener struct {
// 	ListenAddr *net.UDPAddr
// 	HostMAC []byte 
// 	TargetMAC []byte
// 	UdpTargetAddr *net.UDPAddr
// }

type Server struct {
	UdpListenAddr *net.UDPAddr
	UdpTargetAddr *net.UDPAddr
	wtapi.UnimplementedWolControllerServer
} 

func (s *Server) WakeUp(ctx context.Context, req *wtapi.WakeUpRequest) (*emptypb.Empty, error) {
	fmt.Printf("Received: %v", req)

	MAC, err := utils.StringToMac(req.MAC) 
	if err != nil {
		return &emptypb.Empty{}, err
	}

	s.sendMagicPacket(MAC)

	return &emptypb.Empty{}, nil
}

func (s *Server) sendMagicPacket(targetMAC []byte) error {
	packet := genMagicPacket(targetMAC)
	conn, err := net.ListenUDP("udp4", s.UdpListenAddr)

	if err != nil {
		return fmt.Errorf("Cannot listen on provided udpLocalAddr: %v\n", err)
	}

	defer conn.Close()

	_, err = conn.WriteTo(packet, s.UdpTargetAddr)

	if err != nil {
		return err
	}

	return nil
}

// func (l *Listener) Listen() error {
// 	fmt.Println("Start main loop")
// 	conn, err := net.ListenUDP("udp4", l.ListenAddr)

// 	if err != nil {
// 		return fmt.Errorf("Cannot listen on provided listenAddr: %v\n", err)
// 	}

// 	defer conn.Close()

// 	for {
// 		var n int

// 		recvBuf := make([]byte, 258)

// 		n, _, err = conn.ReadFromUDP(recvBuf)

// 		if err != nil {
// 			return fmt.Errorf("Cannot read from connection: %v\n", err)
// 		}

// 		fmt.Printf("Read %v bytes: %v\n", n, recvBuf)

// 		if isMagicPacket(recvBuf, n, l.HostMAC) {
// 			fmt.Println("Found magic packet")
// 			conn.Close()
// 			err = sendMagicPacket(l.TargetMAC, l.ListenAddr, l.UdpTargetAddr)
// 			if err != nil {
// 				return fmt.Errorf("Cannot send magic packet: %v", err)
// 			}
// 			conn, err = net.ListenUDP("udp4", l.ListenAddr)
// 		}
// 	}
// }

// func isMagicPacket(buf []byte, n int, myMac []byte) bool {
// 	return checkSyncSeq(buf[0:6]) && (n == MagicPacketLen) && checkMagicPacketMAC(buf[6:102], myMac)
// }

// func checkMagicPacketMAC(macs []byte, myMac []byte) bool {
// 	isMagicPacket := true 

// 	for i := 0; i < 96; i += 6 {
// 		isMagicPacket = isMagicPacket && bytes.Equal(macs[i : i + 6], myMac)
// 	}

// 	return isMagicPacket
// }

// func checkSyncSeq(slice []byte) bool {
// 	syncSequence := []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}
// 	return bytes.Equal(slice, syncSequence)
//}

// func sendMagicPacket(targetMAC []byte, udpLocalAddr *net.UDPAddr, udpTargetAddr *net.UDPAddr) error {
// 	packet := genMagicPacket(targetMAC)
// 	conn, err := net.ListenUDP("udp4", udpLocalAddr)

// 	if err != nil {
// 		return fmt.Errorf("Cannot listen on provided udpLocalAddr: %v\n", err)
// 	}

// 	defer conn.Close()

// 	_, err = conn.WriteTo(packet, udpTargetAddr)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

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