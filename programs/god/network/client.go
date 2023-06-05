package network

import (
	"context"
	"fmt"
	"gonux/god/contracts"
	"gonux/god/network/proto"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GodClient struct {
	conn      *grpc.ClientConn
	transport string
	addr      string
	stub      proto.GodClient
	actions   map[string]func()
}

func MakeGodClient(serverAddress string) contracts.God {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(serverAddress, opts...)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	// defer conn.Close()

	var transport string
	if strings.HasPrefix(serverAddress, "unix:") {
		transport = "unix"
	} else {
		transport = "tcp"
	}
	return GodClient{
		conn:      conn,
		transport: transport,
		addr:      serverAddress,
		stub:      proto.NewGodClient(conn),
		actions:   make(map[string]func()),
	}
}

func (g GodClient) Register(name string, impl func()) {
	id := fmt.Sprintf("%d", len(g.actions)) // TODO: Generate proper UUID
	g.actions[id] = impl
	// fmt.Println("[Client] Calling Register")
	g.stub.Register(context.Background(), &proto.ActionRegisterRequest{
		Name: name,
		Uuid: id,
	})
}

func (g GodClient) Call(name string) {
	// fmt.Println("[Client] Calling Call")
	g.stub.Call(context.Background(), &proto.ActionIdRequest{
		Name: name,
	})
}

func (g GodClient) WaitFor(name string) {
	// fmt.Println("[Client] Calling WaitFor")
	g.stub.WaitFor(context.Background(), &proto.ActionIdRequest{
		Name: name,
	})
}

func (g GodClient) Subscribe() {
	// fmt.Println("[Client] Calling Subscribe")
	sub, err := g.stub.Subscribe(context.Background(), &proto.Empty{})
	if err != nil {
		fmt.Printf("[Client] Error: %v\n", err)
		return
	}
	for {
		event, err := sub.Recv()
		if err != nil {
			fmt.Printf("[Client] Error: %v\n", err)
			break
		}
		g.actions[event.Uuid]()
	}
}
