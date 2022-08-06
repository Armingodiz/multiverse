protoc --go_out=welcomer/. --go-grpc_out=welcomer/. welcomer/welcomepb/welcome.proto
protoc --go_out=calculator/. --go-grpc_out=calculator/. calculator/calculatorpb/calculator.proto