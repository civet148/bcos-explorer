SHELL=/usr/bin/env bash

# 编译区块浏览器二进制程序
explorer:
	rm -rf bcos-explorer
	go build -ldflags "-s -w" -o bcos-explorer cmd/explorer/main.go

# 编译区块浏览器二进制程序
invoker:
	rm -rf bcos-invoker
	go build -ldflags "-s -w" -o bcos-invoker cmd/invoker/main.go

# 编译并启动区块浏览器
run:explorer
	./bcos-explorer --node-url "http://192.168.20.108:8545"

# 编译并启动合约调用工具部署合约
deploy:invoker
	./bcos-invoker deploy --node-url "http://192.168.20.108:8545"

