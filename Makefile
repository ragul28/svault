build:
	GOOS=linux GOARCH=amd64 go build

run:
	go build && ./svault

init:
	GO111MODULE=on go mod init github.com/ragul28/svault
	GO111MODULE=on go get -u

mod:
	GO111MODULE=on go mod tidy
	GO111MODULE=on go mod verify
	GO111MODULE=on go mod vendor
