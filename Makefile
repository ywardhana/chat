dep:
	GO111MODULE=on go mod vendor

binary:
	GO111MODULE=on go build -o chat_app app/main.go

run: binary
	GO111MODULE=on ./chat_app

test:
	GO111MODULE=on ./bin_coverage/coverage coverage.out
