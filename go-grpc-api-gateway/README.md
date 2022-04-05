go get -d google.golang.org/grpc/cmd/protoc-gen-go-grpc
go install google.golang.org/protobuf/cmd/protoc-gen-go

$ protoc \*.proto --go-grpc_out=. --go-grpc_opt=paths=source_relative
