run:
	go run cmd/api/main.go
build:
	go build -v -o build/api cmd/api/main.go
gen:
	go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@latest -generate types,echo -package api_server schema/openapi.v3.json > generated/api-server/api.gen.go
	go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@latest -generate types,client -package tmdb schema/tmdb.oas.v3.json > generated/tmdb/tmdb.gen.go
test:
	go test -v ./...