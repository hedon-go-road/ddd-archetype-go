start: dockerup db http

dockerup:
	podman compose up -d
dockerdown:
	podman compose down

http:
	swag init
	cd cmd && go run http/main.go

db:
	mysql -h 127.0.0.1 -P 33061 -u root -proot -e "source scripts/db/init_db.sql"

genpb:
	protoc --go_out=. ./protos/*