#SHELL=/usr/bin/env bash

# solidity合约.sol文件路径
CONTRACT_DIR=.
# 编译solidity合约生成的包名
GO_PKG=token
# 编译solidity合约生成.go .bin .abi文件输出目录
GO_DIR=${CONTRACT_DIR}/${GO_PKG}

all: contract

contract:
	mkdir -p ${GO_DIR} \
    && /usr/local/bin/abigen --sol ${CONTRACT_DIR}/token.sol --pkg ${GO_PKG} --out ${GO_DIR}/token.go \
    && solc --bin --abi --overwrite -o ${GO_DIR} ${CONTRACT_DIR}/token.sol