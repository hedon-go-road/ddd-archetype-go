http:
	swag init
	go run ./cmd/http/main.go

genpb:
	protoc --go_out=. ./protos/*