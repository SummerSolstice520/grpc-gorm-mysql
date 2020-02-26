package main

import (
	"google.golang.org/grpc"

	"fmt"
	"context"
	"net"

	pb "grpc-gorm-mysql/proto"
	"grpc-gorm-mysql/mysql"
)

// server port
const port = ":6664"

// myServer implements the service in dboperate.proto
type myServer struct {}

func (m *myServer) Insert(ctx context.Context, in *pb.InsDelUpdRequest) (*pb.Reply, error) {
	mysql.Insert(in.GetTable(), in.GetId(), in.GetName(), in.GetPrice(), in.GetTypeId(), in.GetCreateTime())
	return &pb.Reply{Result: "Insert completed."}, nil
}

func (m *myServer) Delete(ctx context.Context, in *pb.InsDelUpdRequest) (*pb.Reply, error) {
	mysql.Delete(in.GetTable(), in.GetId(), in.GetName(), in.GetPrice(), in.GetTypeId(), in.GetCreateTime())
	return &pb.Reply{Result: "Delete completed."}, nil
}

func (m *myServer) Update(ctx context.Context, in *pb.InsDelUpdRequest) (*pb.Reply, error) {
	mysql.Update(in.GetTable(), in.GetId(), in.GetName(), in.GetPrice(), in.GetTypeId(), in.GetCreateTime())
	return &pb.Reply{Result: "Update completed."}, nil
}

func (m *myServer) Select(ctx context.Context, in *pb.SelectRequest) (*pb.Reply, error) {
	result := mysql.Select(in.GetTable(), in.GetColumns(), in.GetCondition())
	return &pb.Reply{Result: result}, nil
}

func (m *myServer) ExecSql(ctx context.Context, in *pb.SqlRequest) (*pb.Reply, error) {
	mysql.ExecSql(in.GetSql())
	return &pb.Reply{Result: "Execution completed."}, nil
}

func main(){
	// create server listening port
	list, err := net.Listen("tcp", port)
	if err != nil{
		fmt.Println(err)
	}

	// create new server
	server := grpc.NewServer()
	// register service
	pb.RegisterOperationServer(server, &myServer{})

	// grpc service starts
	fmt.Println("grpc service starts...")
	server.Serve(list)
}