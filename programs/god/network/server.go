package network

import (
	"context"
	"fmt"
	"gonux/god/network/proto"
	"net"
	"os"
	"sync"
	"time"

	"google.golang.org/grpc"
)

var callChan chan *proto.ActionCallEvent

type actionRef struct {
	Uuid string
}

type godServer struct {
	proto.UnimplementedGodServer

	actions map[string]actionRef
}

func (s godServer) Register(_ context.Context, req *proto.ActionRegisterRequest) (*proto.Empty, error) {
	// fmt.Printf("[Server] Received Register(%v, %v)\n", req.Name, req.Uuid)
	s.actions[req.Name] = actionRef{
		Uuid: req.Uuid,
	}
	return &proto.Empty{}, nil
}

func handleListener(sub proto.God_SubscribeServer) {
	for event := range callChan {
		err := sub.Send(event)
		if err != nil {
			fmt.Printf("[Server] Error: %v\n", err)
			break
		}
	}
}

func (s godServer) Subscribe(_ *proto.Empty, sub proto.God_SubscribeServer) error {
	// fmt.Printf("[Server] Received Subscribe()\n")
	handleListener(sub)
	return nil
}

func (s godServer) Call(c context.Context, req *proto.ActionIdRequest) (*proto.Empty, error) {
	// fmt.Printf("[Server] Received Call(%v)\n", req.Name)
	callChan <- &proto.ActionCallEvent{
		Uuid: s.actions[req.Name].Uuid,
	}
	return &proto.Empty{}, nil
}

func (s godServer) WaitFor(_ context.Context, req *proto.ActionIdRequest) (*proto.Empty, error) {
	// fmt.Printf("[Server] Received WaitFor(%v)\n", req.Name)
	for {
		if _, ok := s.actions[req.Name]; ok {
			break
		}
		time.Sleep(1 * time.Nanosecond)
	}
	return &proto.Empty{}, nil
}

func newServer() proto.GodServer {
	callChan = make(chan *proto.ActionCallEvent)
	return godServer{
		actions: make(map[string]actionRef),
	}
}

func StartGodServer(network string, address string, wg *sync.WaitGroup) {
	wg.Add(1)
	if network == "unix" {
		os.Remove(address)
	}
	lis, err := net.Listen(network, address)
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	proto.RegisterGodServer(grpcServer, newServer())
	go grpcServer.Serve(lis)
	wg.Done()
	for {
		time.Sleep(1 * time.Nanosecond)
	}
}
