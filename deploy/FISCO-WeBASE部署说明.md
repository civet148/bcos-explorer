# 系统环境

## 1. 硬件和操作系统环境

操作系统: Ubuntu20.04

CPU：8核以上

内存：8G+

磁盘：100G+

本机IP地址：192.168.20.108 (可通过 ip a | grep global 命令查看本机实际IP)

参考网站：https://www.jianshu.com/p/b00a663083b1	

## 2. 时区

```sh
# 选择时区为Asia -> China -> BeiJing
lory@host:~$ tzselect

# 覆盖当前时区文件
lory@host:~$ sudo cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime 
```

# 安装指南

## 1. 安装JDK11

```sh
# 安装JDK11
lory@host:~$ sudo apt install -y openjdk-11-jre-headless 

# 查看JDK安装结果
lory@host:~$ java -version
openjdk version "11.0.13" 2021-10-19
OpenJDK Runtime Environment (build 11.0.13+8-Ubuntu-0ubuntu1.20.04)
OpenJDK 64-Bit Server VM (build 11.0.13+8-Ubuntu-0ubuntu1.20.04, mixed mode, sharing)

# 查看JDK安装路径
lory@host:~$ ls -l /usr/lib/jvm
total 4
lrwxrwxrwx 1 root root   21 Oct 29 09:11 java-1.11.0-openjdk-amd64 -> java-11-openjdk-amd64
drwxr-xr-x 7 root root 4096 Mar  2 12:26 java-11-openjdk-amd64

# 编辑~/.bashrc添加环境变量
lory@host:~$ sudo vi ~/.bashrc

# 添加如下内容保存并退出vi
export JAVA_HOME=/usr/lib/jvm/java-11-openjdk-amd64
export JRE_HOME=${JAVA_HOME}/jre
export CLASSPATH=.:${JAVA_HOME}/lib:${JRE_HOME}/lib
export PATH=${JAVA_HOME}/bin:${JRE_HOME}/bin:$PATH

# 环境变量生效
lory@host:~$ source ~/.bashrc
```



## 2. 安装MySQL数据库

```sh
# 安装MySQL
lory@host:~$ sudo apt install -y mysql-server

# 查看MySQL默认用户和密码
lory@host:~$ sudo cat /etc/mysql/debian.cnf 
# Automatically generated for Debian scripts. DO NOT TOUCH!
[client]
host     = localhost
user     = debian-sys-maint
password = V55LGfMeYGnzsjhP
socket   = /var/run/mysqld/mysqld.sock
[mysql_upgrade]
host     = localhost
user     = debian-sys-maint
password = V55LGfMeYGnzsjhP
socket   = /var/run/mysqld/mysqld.sock

# 登录MySQL终端(密码V55LGfMeYGnzsjhP）
lory@host:~$ mysql -udebian-sys-maint -p mysql

# 修改root密码和口令加密方式并开启远程登录(视实际情况而定，如果无必要可以只修改密码不开启远程登录)
# host='%'表示开启远程访问，如果不开启就不要这个SQL字句
mysql> update user set plugin='mysql_native_password', host='%',authentication_string='' where user='root';
Query OK, 1 row affected (0.00 sec)
Rows matched: 1  Changed: 1  Warnings: 0

mysql> alter user 'root'@'%' IDENTIFIED WITH mysql_native_password BY 'Stos@123';
mysql> flush privileges;
mysql> quit

# 以root身份登录MySQL
lory@host:~$ mysql -u root -pStos@123

#建立webase账户，密码为webase_pswd
mysql> CREATE USER  'webase'@'localhost'  IDENTIFIED BY  'webase_pswd';
#建立两个数据库
mysql> create DATABASE webasenodemanager;
mysql> create DATABASE webasesign;

#将两个数据库的所有权限赋予本地的webase帐号
mysql> GRANT ALL PRIVILEGES ON webasenodemanager.*  To 'webase'@'localhost';
mysql> GRANT ALL PRIVILEGES ON webasesign.*  To 'webase'@'localhost';
mysql> flush privileges;
mysql> quit
```



## 3. Python3+PyMySQL安装

```sh
# 查看python3是否安装
lory@host:~$ python3 --version
Python 3.8.10

# 如果没安装则安装
lory@host:~$ sudo apt-get update && sudo apt-get install -y python3

# 创建软链接
lory@host:~$ sudo ln -nsf /usr/bin/python3 /usr/bin/python

# 安装PyMySQL依赖包
lory@host:~$ sudo apt-get update && sudo apt-get install -y python3-pip && sudo pip3 install PyMySQL
...
...
Installing collected packages: PyMySQL
Successfully installed PyMySQL-1.0.2
```



## 4. 下载WeBASE安装包

```sh
#进入项目目录
lory@host:~$ mkdir -p ~/webase && cd ~/webase

#获取部署安装包
lory@host:~$ wget https://osp-1257653870.cos.ap-guangzhou.myqcloud.com/WeBASE/releases/download/v1.5.0/webase-deploy.zip

#安装解压工具
lory@host:~$ sudo apt install -y unzip

#解压安装包
lory@host:~$ unzip webase-deploy.zip
```



## 5. 修改WeBASE数据库配置

```toml

# 更改WeBASE数据账户和密码(IP或端口改动也需要进行相应修改),其他保持不变
lory@host:~$ cd ~/webase/webase-deploy && vi common.properties

# Mysql database configuration of WeBASE-Node-Manager
mysql.ip=localhost
mysql.port=3306
mysql.user=webase
mysql.password=webase_pswd
mysql.database=webasenodemanager

# Mysql database configuration of WeBASE-Sign
sign.mysql.ip=localhost
sign.mysql.port=3306
sign.mysql.user=webase
sign.mysql.password=webase_pswd
sign.mysql.database=webasesign

# Node rpc service port（RPC端口冲突，可修改为其他端口)
node.rpcPort=8545

# Use existing chain or not (yes/no)
if.exist.fisco=no

### if build new chain, [if.exist.fisco=no]
# Configuration required when building a new chain 
# Fisco-bcos version
fisco.version=2.7.2

# Number of building nodes (设置启动2个节点）
node.counts=2

### if using exited chain, [if.exist.fisco=yes]
# The path of the existing chain, the path of the start_all.sh script
# Under the path, there should be a 'sdk' directory where the SDK certificates (ca.crt, node.crt and node. Key) are stored
fisco.dir=/data/app/nodes/127.0.0.1
# Absolute path of the connected node in WeBASE-Front
# Under the path, there is a conf directory where node certificates (ca.crt, node.crt and node. Key) are stored
node.dir=/data/app/nodes/127.0.0.1/node0
```



## 6. 安装WeBASE

```sh
#进入webase目录
lory@host:~$ cd ~/webase/webase-deploy
#部署并启动所有WeBASE服务
lory@host:~$ python3 deploy.py installAll

lory@host:~/webase/webase-deploy$ python3 deploy.py installAll
============================================================

              _    _     ______  ___  _____ _____ 
             | |  | |    | ___ \/ _ \/  ___|  ___|
             | |  | | ___| |_/ / /_\ \ `--.| |__  
             | |/\| |/ _ | ___ |  _  |`--. |  __| 
             \  /\  |  __| |_/ | | | /\__/ | |___ 
              \/  \/ \___\____/\_| |_\____/\____/  
    
============================================================
==============      checking envrionment      ==============
check git...
check finished sucessfully.
check openssl...
check finished sucessfully.
check curl...
check finished sucessfully.
check wget...
check finished sucessfully.
check dos2unix...
[sudo] password for impool: 
Reading package lists... Done
Building dependency tree       
Reading state information... Done

# 安装过程提示输入y/N的全部输入y回车

onelineOutput: nginx: the configuration file /etc/nginx/nginx.conf syntax is ok
Defualt nginx config path: /etc/nginx
==============      Starting WeBASE-Web       ==============
==============      WeBASE-Web Started        ==============
==============  Init Front for Mgr start...   ==============
=============== Init Front for Mgr fail. Please view log file (default path:./log/). ==============
============================================================
==============      deploy  has completed     ==============
============================================================
==============    webase-web version  v1.5.0        ========
==============    webase-node-mgr version  v1.5.0   ========
==============    webase-sign version  v1.5.0       ========
==============    webase-front version  v1.5.0      ========
============================================================
```



## 7. 登录WeBASE

浏览器打开http://192.168.20.108:5000 初始账户：**admin** 密码：**Abcd1234**

注意：**如果出现验证码加载失败请看第8章节解决**

![img.png](img.png)

## 8. 网页图形验证码无法加载问题

```sh
# 修改webase-sign配置文件在数据库url最后加useSSL=false
lory@host:~/webase/webase-deploy$ vi webase-sign/conf/application.yml 

spring:
  cache:
    type: simple
  datasource:
    # 数据库连接信息(URL后加上 &useSSL=false&allowPublicKeyRetrieval=true)
    url: jdbc:mysql://localhost:3306/webasesign?serverTimezone=GMT%2B8&useUnicode=true&characterEncoding=utf8&useSSL=false&allowPublicKeyRetrieval=true
    # 数据库用户名
    username: "webase"
    # 数据库密码
    password: "webase_pswd"
    driver-class-name: com.mysql.cj.jdbc.Driver
    hikari:
      connection-test-query: SELECT 1 FROM DUAL
      connection-timeout: 30000
      maximum-pool-size: 20
      max-lifetime: 1800000
      minimum-idle: 5

# 修改webase-node-mgr配置文件在数据库url最后加 (&useSSL=false&allowPublicKeyRetrieval=true)
lory@host:~/webase/webase-deploy$ vi webase-node-mgr/conf/application.yml

# database connection configuration
spring:
  datasource:
    driver-class-name: com.mysql.cj.jdbc.Driver
    url: jdbc:mysql://localhost:3306/webasenodemanager?serverTimezone=GMT%2B8&useUnicode=true&characterEncoding=utf-8&zeroDateTimeBehavior=convertToNull&useSSL=false&allowPublicKeyRetrieval=true
    username: "webase"
    password: "webase_pswd"
    initialSize: 10
    minIdle: 5
    maxActive: 30


###------------------------------------------------------------------------------------------

# 重启服务
lory@host:~/webase/webase-deploy$ python3 deploy.py stopAll && python3 deploy.py startAll 
```

# docker部署指南

[一键Docker部署 — WeBASE v1.5.4 文档 (webasedoc.readthedocs.io)](https://webasedoc.readthedocs.io/zh_CN/latest/docs/WeBASE-Install/docker_install.html)
