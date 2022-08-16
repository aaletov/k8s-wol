package main

import (
	"fmt"
	"net"
	"flag"
	"google.golang.org/grpc"
	"github.com/aaletov/k8s-wol/api/wtapi"
	"github.com/aaletov/k8s-wol/utils"
	"github.com/aaletov/k8s-wol/server"
)

var (
	logPath = flag.String("logpath", "", "Log path for Node Volume Manager service")
)

func main() {
	flag.Parse()

	var (
		hostAddr string
		targetAddr string
		port int
	)

	logger, err := utils.InitLogger(*logPath)
	logger.Info("Initialized logger")

	err = server.RegisterNode(logger)
	if err != nil {
		logger.Errorf("Failed to register node: %v", err)
		return
	}
	logger.Info("Registered node")

	flag.StringVar(&hostAddr, "host_addr", "", "")
	flag.StringVar(&targetAddr, "target_addr", "", "")
	flag.IntVar(&port, "port", 50051, "")
	flag.Parse()

	udpHostAddr, _ := net.ResolveUDPAddr("udp4", hostAddr)
	udpTargetAddr, _ := net.ResolveUDPAddr("udp4", targetAddr)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Println("Failed to listen: %v", err)
		return
	}

	s := grpc.NewServer()
	wtapi.RegisterWolControllerServer(s, &server.Server{udpHostAddr, udpTargetAddr, wtapi.UnimplementedWolControllerServer{}})

	fmt.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		fmt.Println("Failed to serve: %v", err)
	}
}
