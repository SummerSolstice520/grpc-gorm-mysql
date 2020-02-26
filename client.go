package main

import (
	"google.golang.org/grpc"

	"fmt"
	"context"

	pb "grpc-gorm-mysql/proto"
)

// server address and port
const address = "127.0.0.1:6664"

func main() {
	// create Server connector
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil{
		fmt.Println(err)
	}
	defer conn.Close()

	// create new client
	c := pb.NewOperationClient(conn)

	// request
	fmt.Printf("Successfully connected to server.\n")
	loop:
		for {
			fmt.Printf("Please select the following operations:\n 1-select\t 2-insert\t 3-delete\t 4-update\t 5-sql\t 0-quit\n")
			var op string
			fmt.Scanln(&op)
			var result *pb.Reply
			switch op {
			// select
			case "1":
				fmt.Printf("Please enter the table name:\n")
				var table string
				fmt.Scanln(&table)
				fmt.Printf("Please enter the columns name: \n  example1: name,price \n  example2: * (for all) \n")
				var columns string
				fmt.Scanln(&columns)
				fmt.Printf("Please enter the condition: (skip with enter)\n")
				var con string
				fmt.Scanln(&con)
				result, err = c.Select(context.Background(), &pb.SelectRequest{Columns:columns, Table:table, Condition:con})
			// insert
			case "2":
				fmt.Printf("Please enter the data: \n example: 2 yan 2.0 3 5 \n")
				var id int32
				var name string
				var price float32
				var typeId int32
				var createTime int64
				fmt.Scanf("%d %s %f %d %d", &id, &name, &price, &typeId, &createTime)
				result, err = c.Insert(context.Background(), &pb.InsDelUpdRequest{Id:id, Name:name, Price:price, TypeId:typeId, CreateTime:createTime})
			// delete
			case "3":
				fmt.Printf("Please enter the data: \n example: 2 yan 2.0 3 5 \n")
				var id int32
				var name string
				var price float32
				var typeId int32
				var createTime int64
				fmt.Scanf("%d %s %f %d %d\n", &id, &name, &price, &typeId, &createTime)
				result, err = c.Delete(context.Background(), &pb.InsDelUpdRequest{Id:id, Name:name, Price:price, TypeId:typeId, CreateTime:createTime})
			// update
			case "4":
				fmt.Printf("Please enter the data: \n example: 3 wang 20  50 \n")
				var id int32
				var name string
				var price float32
				var typeId int32
				var createTime int64
				fmt.Scanf("%d %s %f %d %d\n", &id, &name, &price, &typeId, &createTime)
				result, err = c.Update(context.Background(), &pb.InsDelUpdRequest{Id:id, Name:name, Price:price, TypeId:typeId, CreateTime:createTime})
			// sql
			case "5":
				fmt.Printf("Please enter the sql: \n")
				var sql string
				fmt.Scanln(&sql)
				result, err = c.ExecSql(context.Background(), &pb.SqlRequest{Sql:sql})
			// quit
			case "0":
				break loop
			}

			if err != nil{
				fmt.Println("Failed to get reply.")
				return
			}
			fmt.Println(result)
		}
	fmt.Println("Successfully disconnected.")
}