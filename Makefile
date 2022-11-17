.PHONY: all

all: darwin-amd64 darwin-arm64 linux-amd64 linux-arm64

check:
	go vet ./

darwin-amd64:
	GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w -X main.buildTime=`date +%Y%m%d.%H:%M:%S` -X main.buildCommit=`git rev-parse --short=12 HEAD` -X main.buildBranch=`git branch --show-current`" -o ./releases/go_github_faster_darwin_amd64 ./

darwin-arm64:
	GOOS=darwin GOARCH=arm64 go build -ldflags "-s -w -X main.buildTime=`date +%Y%m%d.%H:%M:%S` -X main.buildCommit=`git rev-parse --short=12 HEAD` -X main.buildBranch=`git branch --show-current`" -o ./releases/go_github_faster_darwin_arm64 ./

linux-amd64:
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w -X main.buildTime=`date +%Y%m%d.%H:%M:%S` -X main.buildCommit=`git rev-parse --short=12 HEAD` -X main.buildBranch=`git branch --show-current`" -o ./releases/go_github_faster_linux_amd64 ./

linux-arm64:
	GOOS=linux GOARCH=arm64 go build -ldflags "-s -w -X main.buildTime=`date +%Y%m%d.%H:%M:%S` -X main.buildCommit=`git rev-parse --short=12 HEAD` -X main.buildBranch=`git branch --show-current`" -o ./releases/go_github_faster_linux_arm64 ./
