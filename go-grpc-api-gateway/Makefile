proto:
		protoc pkg/**/pb/*.proto --go_out=:. --go_opt=paths=source_relative

service:
	go run cmd/main.go