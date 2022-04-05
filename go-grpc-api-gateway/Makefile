proto:
		protoc pkg/**/pb/*.proto --go-grpc_out=. --go-grpc_opt=paths=source_relative

service:
	go run cmd/main.go