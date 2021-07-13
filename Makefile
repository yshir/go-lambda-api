AWS_PROFILE := go-lambda-infra

.PHONY: deps
deps:
	go mod download

.PHONY: build
build: deps
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o bin/lambda cmd/lambda/main.go

.PHONY: test
test:
	export DYNAMO_TABLE_USERS=local_users;\
	go test -race -v -count=1 ./

.PHONY: generate-server
generate-server:
	rm -rf ./gen/models ./gen/restapi/operations
	swagger generate server --exclude-main -f ./swagger.yml -t gen

.PHONY: zip
zip: build
	zip -j bin/lambda.zip bin/lambda

.PHONY: deploy
deploy: zip
	aws lambda update-function-code --profile $(AWS_PROFILE) --region ap-northeast-1 --function-name go-lambda-api --zip-file fileb://bin/lambda.zip
