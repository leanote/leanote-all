# leanote-all

该仓库包含leanote依赖及leanote最新源码, 特为需要源码安装的朋友准备

## 安装步骤

1. 安装golang
2. 下载该仓库
3. 安装mongodb
4. 导入初始数据
5. 使用revel运行leanote

## 安装golang

去golang.org上下载最新版golang (1.3.1+) 如果被墙, 可以在 http://golangtc.com/download 下载.

快速下载:

* linux 64位: http://www.golangtc.com/static/go/go1.3.1.linux-amd64.tar.gz
* linux 32位: http://www.golangtc.com/static/go/go1.3.1.linux-386.tar.gz

假设将文件下载到 /home/user1下, 解压文件

```
$> cd /home/user1
$> tar -xzvf go1.3.1.linux-amd64.tar.gz
```

在/home/user1下新建一个目录gopackage (这里面会放go的包和编译后的文件)

```
$> mkdir /home/user1/gopackage
```

配置环境变量, 编辑/etc/profile文件
```
export GOROOT=/home/user1/go
export GOPATH=/home/user1/gopackage
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```
使环境变量生效:
```
$> source /etc/profile
```
查看go是否安装成功
```
$> go version
# 若出现类似以下信息证明安装成功
go version go1.3.1 linux/amd64
```

## 下载该仓库

该仓库包含leanote所有依赖包和leanote最新源码

1. src复制到gopackage下 (/home/user1/gopackage/下)
2. 生成revel命令
```
$> go get github.com/revel/cmd/revel
``` 

## 安装mongodb

到 http://www.mongodb.org/downloads 去下载

传送门: 64位linux mongodb下载链接 http://www.mongodb.org/dr//fastdl.mongodb.org/linux/mongodb-linux-x86_64-2.6.4.tgz/download

下载到/home/user1下, 直接解压即可
```
$> cd /home/user1
$> tar -xzvf mongodb-linux-x86_64-2.6.4.tgz/
```

为了快速使用mongodb的命令, 可以配置环境变量,

编辑 ~/.profile或/etc/profile 将mongodb bin路径加入即可.
```
$> sudo vim /etc/profile
```
添加:
```
export PATH=$PATH:/home/user1/mongodb-linux-x86_64-2.6.4/bin
```
使环境变量生效:
```
$> source /etc/profile
```
### 简单使用mongodb

先在/home/user1下新建一个目录data存放mongodb数据
```
mkdir /home/user1/data
```

```
# 开启mongodb
mongod --dbpath /home/user1/data
```

这时mongod已经启动了

重新打开一个终端, 使用下mongodb
```
$> mongo
> show dbs
...数据库列表
```

mongodb安装到此为止, 下面为mongodb导入数据leanote初始数据

## 导入初始数据
首先确保mongodb已开启

leanote初始数据在/home/user1/gopackage/src/github.com/leanote/leanote/mongodb_backup/leanote_install_data中

打开终端, 输入以下命令导入数据:
```
$> mongorestore -h localhost -d leanote --directoryperdb /home/user1/gopackage/src/github.com/leanote/leanote/mongodb_backup/leanote_install_data/
```

现在在mongodb中已经新建了leanote数据库, 可用命令查看下leanote有多少张表
```
$> mongo
> show dbs #　查看数据库
leanote	0.203125GB
local	0.078125GB
> use leanote # 切换到leanote
switched to db leanote
> show collections # 查看表
files
has_share_notes
note_content_histories
note_contents
notebooks
notes
share_notebooks
share_notes
system.indexes
system.users
tags
tokens
user_blogs
users
```

初始数据users表中已有2个用户:
```
user1 username: admin, password: abc123 (管理员, 只有该用户可以管理后台)
user2 username: demo@leanote.com, password: demo@leanote.com (仅共体验使用)
```

## 配置leanote
文件: leanote/conf/app.conf

请修改要修改app.secret, 请随意修改一个值, 若不修改, 会有安全问题!

其它的配置请保持不变, 若需要配置数据库信息, 请查看下文"问题3"

## 运行leanote
```
$> revel run github.com/leanote/leanote
```
恭喜你, 打开浏览器输入: localhost体验leanote吧


## 问题汇总

### 问题1: 
```
Go to /@tests to run the tests.
panic: auth fails

goroutine 1 [running]:
github.com/leanote/leanote/app/db.Init()
	/home/life/gopackage1/src/github.com/leanote/leanote/app/db/Mgo.go:64 +0x356
```
解答:

数据库配置有问题, 请修改leanote/conf/app.conf文件, 是否用户名和密码配置有误?

### 问题2: 修改默认80端口?
修改leanote/conf/app.conf, 比如改成9000
```
http.port=9000

site.url=http://localhost:9000
```

### 问题3: 为数据库添加用户

像mysql一样有root用户, mongodb初始是没有用户的, 这样很不安全, 所以要为leanote数据库新建一个用户来连接leanote数据库(注意, 并不是为leanote的表users里新建用户, 而是新建一个连接leanote数据库的用户, 类似mysql的root用户).
```
# 首先切换到leanote数据库下
> use leanote;
# 添加一个用户root, 密码是abc123
> db.addUser("root", "abc123");
{
	"_id" : ObjectId("53688d1950cc1813efb9564c"),
	"user" : "root",
	"readOnly" : false,
	"pwd" : "e014bfea4a9c3c27ab34e50bd1ef0955"
}
# 测试下是否正确
> db.auth("root", "abc123");
1 # 返回1表示正确
```

用户添加好后重新运行下mongodb, 并开启权限验证. 在mongod的终端按ctrl+c即可退出mongodb.

启动mongodb:
```
$> mongod --dbpath /home/user1/data --auth
```

**还要修改配置文件** : 修改 leanote/conf/app.conf:

````
db.host=localhost
db.port=27017
db.dbname=leanote # required
db.username=root # if not exists, please leave blank
db.password=abc123 # if not exists, please leave blank
```
