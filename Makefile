.PHONY: all

all: darwin-amd64 darwin-arm64 linux-amd64 linux-arm64 windows-amd64

check:
	go mod tidy
	go vet ./

darwin-amd64:
	GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o ./releases/github_faster_darwin_amd64 ./

darwin-arm64:
	GOOS=darwin GOARCH=arm64 go build -ldflags "-s -w" -o ./releases/github_faster_darwin_arm64 ./

linux-amd64:
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ./releases/github_faster_linux_amd64 ./

linux-arm64:
	GOOS=linux GOARCH=arm64 go build -ldflags "-s -w" -o ./releases/github_faster_linux_arm64 ./

windows-amd64:
	GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o ./releases/github_faster_windows_amd64.exe ./