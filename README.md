# 1. WeBASE管理平台

## 1.1 安装WeBASE
- WeBASE管理平台是一套包含BCOS节点、区块浏览器以及账号密钥和节点管理的综合平台
- 安装过程见deploy目录下的 《FISCO-WeBASE部署说明.md》

## 1.2 

# 2. 安装BCOS合约编译工具

## 2.1 下载FISCO-BCOS官方go sdk

```shell
$ cd $GOPATH/src && git clone https://github.com/FISCO-BCOS/go-sdk
```


## 2.1 下载solc编译器

- solc是solidity合约编译器

```shell
# 下载对应系统的solc版本v0.4.25并拷贝到/usr/loca/bin目录下更名为solc
$ cd $GOPATH/src/go-sdk/tools && chmod +x *.sh && ./download_solc.sh 0.4.25 && sudo cp solc-0.4.25 /usr/local/bin/solc

Downloading solc 0.4.25 solc-linux.tar.gz from https://github.com/FISCO-BCOS/solidity/releases/download/v0.4.25/solc-linux.tar.gz
==============================================================
[INFO] os            : linux
[INFO] solc version  : 0.4.25
[INFO] solc location : ./solc-0.4.25
==============================================================
[INFO] ./solc-0.4.25 --version
solc, the solidity compiler commandline interface
Version: 0.4.25+commit.46d177ad.mod.Linux.g++
```

## 2.2 编译abigen工具

- abigen工具用于将solidity合约代码生成go语言代码
   
```shell
# 通过代码编译abigen工具并拷贝到/usr/local/bin目录
$ cd $GOPATH/src/go-sdk/cmd/abigen
$ go mod tidy &&  go build -ldflags "-s -w" -o abigen && sudo cp abigen /usr/local/bin

# 查看abigen版本信息
$ abigen -v
abigen version 1.10.12-stable
```

# 3. 编译solidity合约

- 示例合约token.sol位于bcos-explorer/contract目录下

```shell
# 进入合约目录并make编译生成go代码到bcos-explorer/contract/token目录中
$ cd contract && make 

# .abi文件是合约接口描述文件，.bin是合约二进制接口，.go文件则是我们部署和调用合约方法的代码文件
$ ls -l ./token
total 88
drwxrwxrwx 1 lory lory  4096 Jun  2 08:22 ./
drwxrwxrwx 1 lory lory  4096 Jun  2 13:39 ../
-rwxrwxrwx 1 lory lory   913 Jun  2 13:39 Logger.abi*
-rwxrwxrwx 1 lory lory   164 Jun  2 13:39 Logger.bin*
-rwxrwxrwx 1 lory lory  1609 Jun  2 13:39 Token.abi*
-rwxrwxrwx 1 lory lory  2358 Jun  2 13:39 Token.bin*
-rwxrwxrwx 1 lory lory 76420 Jun  2 13:39 token.go*

```

