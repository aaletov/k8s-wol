package server

import (
	"fmt"
	"net"
	"context"
	"github.com/aaletov/k8s-wol/api/wtapi"
	"github.com/aaletov/k8s-wol/utils"
	"google.golang.org/protobuf/types/known/emptypb"
)

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
