# grpc-gorm-mysql
## server.go
- create a grpc server
- implement the service in .proto

## client.go
- create a grpc client
- cyclic send data with reply
- command line interaction

## mysql/mysql.go
- table struct
- connMysql function connects Mysql with dsn
- InsDelUpd function operates on row-level data
- Select function returns query results
- ExecSql function executes SQL statement

## proto/dboperate.proto
- define services: insert, delete, update, select, execute sql
- define messages: insDelUpdRequest, selectRequest, sqlRequest, reply
