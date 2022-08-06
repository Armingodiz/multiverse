protoc --go_out=welcomer/. --go-grpc_out=welcomer/. welcomer/welcomepb/welcome.proto
go run welcomer/server/server.go
go run welcomer/client/client.go
protoc --go_out=calculator/. --go-grpc_out=calculator/. calculator/calculatorpb/calculator.proto
go run calculator/server/server.go
go run calculator/client/client.go