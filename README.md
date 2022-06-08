# 1. WeBASE管理平台

- WeBASE管理平台是一套包含BCOS节点、区块浏览器以及账号密钥和节点管理的综合平台
- 安装过程见deploy目录下的 《FISCO-WeBASE部署说明.md》
- 后续使用到的用户私钥可以通过WeBASE平台私钥管理菜单创建

# 2. 安装BCOS合约编译工具

## 2.1 下载FISCO-BCOS官方go sdk

```shell
$ cd $GOPATH/src && git clone https://github.com/FISCO-BCOS/go-sdk
```


## 2.2 下载solc编译器

- solc是solidity合约编译器

```shell
# 下载对应操作系统版本solc v0.4.25并拷贝到/usr/loca/bin目录下更名为solc
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

## 2.3 编译abigen工具

- abigen工具用于将solidity合约代码生成go语言代码
   
```shell
# 通过代码编译abigen工具并拷贝到/usr/local/bin目录
$ cd $GOPATH/src/go-sdk/cmd/abigen
$ go mod tidy &&  go build -ldflags "-s -w" -o abigen && sudo cp abigen /usr/local/bin

# 查看abigen版本信息
$ /usr/local/bin/abigen -v
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

# 4. 管理账户私钥

## 4.1 创建账户密钥
```shell
# 进入cmd/bcos-account目录编译bcos-account程序
$ cd cmd/bcos-account && go build -ldflags "-s -w" -o bcos-account
$ ./bcos-account new

address [0x5f9c56B86E94e9302D7074b8E5f1397a798452E5]
public key [032ca918b9e91ab3476d656ba66aea74a1318f5417f74e63de86ffa0d70e30cd24]
private key [44af095ec12d675dc338898a15c8cfd8e6d7ea8de608942a590a1b9fe27334dc]
phrase [road always practice armed know stem crucial wave crouch enact candy earn]

```

## 4.2 助记词恢复密钥

```shell
# 进入cmd/bcos-account目录编译bcos-account程序
$ cd cmd/bcos-account && go build -ldflags "-s -w" -o bcos-account
$ ./bcos-account recover "road always practice armed know stem crucial wave crouch enact candy earn"

address [0x5f9c56B86E94e9302D7074b8E5f1397a798452E5]
public key [032ca918b9e91ab3476d656ba66aea74a1318f5417f74e63de86ffa0d70e30cd24]
private key [44af095ec12d675dc338898a15c8cfd8e6d7ea8de608942a590a1b9fe27334dc]
phrase [road always practice armed know stem crucial wave crouch enact candy earn]

```

# 5. 部署合约

- 部署合约需要用到私钥(参考第4.1章节自行创建账户密钥)

```shell
# 进入cmd/bcos-invoker目录编译bcos-invoker程序
$ cd cmd/bcos-invoker && go build -ldflags "-s -w" -o bcos-invoker

# 编译成功后执行如下命令行(假设BCOS节点URL是 http://192.168.20.108:8545) 执行成功返回如下信息：合约地址、交易哈希以及owner账户余额信息
$ ./bcos-invoker deploy --node-url http://192.168.20.108:8545

contract address: 0xd3D82d2515DF022b60F7Fad5daeD06AB046b42df
tx hash: 0x67c2ecf858691294c74f0247bedd75c872e9e59cd113ccc34c644555b01df26c
owner 0x5B0c43004e0a68Eb197c629CE78Da62d65Aa6C03 balance 10000000 tokens

```

# 6. 转账(前提：合约部署成功)

```shell
# 增加--contract-addr 参数指定第4章节部署合约后得到的合约地址进行转账操作
$ ./bcos-invoker transfer --node-url "http://192.168.20.108:8545" --contract-addr 0xd3D82d2515DF022b60F7Fad5daeD06AB046b42df

tx hash [0x59f88e2f8a14e51b6461f936ae06484403b9e1589b92ac7bce17945fb698ba6a]
owner 0x5B0c43004e0a68Eb197c629CE78Da62d65Aa6C03 balance 9999900 tokens
receiver 0x90Cfd4D61C9D4C63f2e4648229775ABa19ced8dF balance 100 tokens
```

# 7. 解析区块数据

指定URL和合约地址（查询最后一个区块高度的交易数据并解析方法名称、参数名称和参数值）, 也可以通过添加--height选项指定固定高度去解析

```shell
$ go build -ldflags "-s -w" -o bcos-explorer cmd/bcos-explorer/main.go
$ ./bcos-explorer --node-url http://192.168.20.108:8545 --contract-addr 0xd3D82d2515DF022b60F7Fad5daeD06AB046b42df

method name [transferFrom]
input type [address] name [from]
input type [address] name [to]
input type [uint64] name [amount]
input name [amount] value [100]
input name [from] value [0x5B0c43004e0a68Eb197c629CE78Da62d65Aa6C03]
input name [to] value [0x90Cfd4D61C9D4C63f2e4648229775ABa19ced8dF]
```

