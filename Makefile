SHELL=/usr/bin/env bash

# 编译账户密钥工具
account:
	rm -rf bcos-account
	go build -ldflags "-s -w" -o bcos-account cmd/bcos-account/main.go

# 编译区块浏览器程序
explorer:
	rm -rf bcos-explorer
	go build -ldflags "-s -w" -o bcos-explorer cmd/bcos-explorer/main.go

# 编译合约调用工具
invoker:
	rm -rf bcos-invoker
	go build -ldflags "-s -w" -o bcos-invoker cmd/bcos-invoker/main.go

# 编译并启动区块浏览器
run:explorer
	./bcos-explorer --node-url http://192.168.20.108:8545 --contract-addr 0xd3D82d2515DF022b60F7Fad5daeD06AB046b42df

# 编译并启动合约调用工具部署合约
deploy:invoker
	./bcos-invoker deploy --node-url "http://192.168.20.108:8545"

# 编译并启动合约调用工
transfer:invoker
	./bcos-invoker transfer --node-url "http://192.168.20.108:8545" --contract-addr 0x57E4b24712495Ef15bA0EED024189041b11F5E49

