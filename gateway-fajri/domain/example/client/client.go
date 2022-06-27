package client

import (
	"context"
	"os"

	pb "gitlab.com/qasir/gateway-fajri/domain/example/proto"
	qGrpc "gitlab.com/qasir/web/project/qasircore.git/transport/grpc"
)

// Example client handler
func Example(ctx context.Context, req *pb.ExampleRequest) (*pb.ExampleResponse, error) {
	conn := qGrpc.Dial(os.Getenv("EXAMPLE_GRPC"))
	defer conn.Close()
	client := pb.NewExampleClient(conn)

	return client.Example(ctx, req)
}