package main

import (
	"fmt"
	"time"
	"flag"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"github.com/aaletov/wol-translator/api/wtapi"
)

func main() {
	var (
		addr string
		stringMAC string
	)

	flag.StringVar(&addr, "addr", "localhost:50051", "")
	flag.StringVar(&stringMAC, "MAC", "", "")
	flag.Parse()

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Did not connect: %v", err)
		return
	}
	defer conn.Close()
	c := wtapi.NewWolControllerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = c.WakeUp(ctx, &wtapi.WakeUpRequest{MAC: stringMAC})
	if err != nil {
		fmt.Println("gRPC failed: %v", err)
		return
	}

	fmt.Println("gRPC OK")
	return
}