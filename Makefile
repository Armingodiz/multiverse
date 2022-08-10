calculatorClient:
	protoc --go_out=core/services/calculatorService/. --go-grpc_out=core/services/calculatorService/. core/services/calculatorService/calculatorpb/calculator.proto
calculatorServer:
	protoc --go_out=calculator/. --go-grpc_out=calculator/. calculator/calculatorpb/calculator.proto
welcomerClient:
	protoc --go_out=core/services/welcomerService/. --go-grpc_out=core/services/welcomerService/. core/services/welcomerService/welcomepb/welcome.proto
welcomerServer:
	protoc --go_out=welcomer/. --go-grpc_out=welcomer/. welcomer/welcomepb/welcome.proto
multiverseBuild:
	docker-compose build
multiverseRun:
	docker-compose up -d
multiverseCore:
	docker-compose up core
multiverse: multiverseBuild multiverseRun multiverseCore