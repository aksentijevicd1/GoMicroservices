swagger_v2:
	swag init -o ./docs/swagger -g ./main.go --parseDependency true

swagger_v3:
	go run openapi/convert/main.go

swagger: swagger_v2 swagger_v3
