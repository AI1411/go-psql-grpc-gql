package server

import (
	"context"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"

	pb "github.com/AI1411/go-pg-ci-example/grpc"
)

const BufSize = 1024 * 1024

var lis *bufconn.Listener

func bufDialer(ctx context.Context, addr string) (net.Conn, error) {
	return lis.Dial()
}

func TestListTest(t *testing.T) {
	initializeForServerTest()
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewTestServiceClient(conn)
	want := &pb.ListTestResponse{
		Tests: []*pb.GetTestResponse{
			{
				Id:   1,
				Name: "test1",
			},
			{
				Id:   2,
				Name: "test2",
			},
		},
	}

	got, err := client.ListTest(ctx, &pb.ListTestRequest{})

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, want.Tests, got.Tests)
}

func TestGetTest(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewTestServiceClient(conn)
	want := uint32(1)
	got, err := client.GetTest(ctx, &pb.GetTestRequest{Id: 1})

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, want, got.Id)
}

func TestCreateTest(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewTestServiceClient(conn)
	want := "test"
	got, err := client.CreateTest(ctx, &pb.CreateTestRequest{Name: "test"})

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, want, got.Name)
}
