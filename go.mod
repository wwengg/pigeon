module github.com/wwengg/pigeon

go 1.15

require (
	github.com/golang/protobuf v1.5.2
	github.com/smallnest/rpcx v1.6.9
	github.com/spf13/cobra v1.2.1 // indirect
	github.com/wwengg/arsenal v0.0.2-0.20210822075139-8466df097abe
	go.uber.org/zap v1.19.0
)

//replace github.com/wwengg/arsenal => ../arsenal
