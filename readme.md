# beego http demo

- prepare：

create db named bhg_main

run sql/bhg_main on msyql


- config:

config/dev.conf
```
mysqluser = "root"
mysqlpass = "root"
mysqlurls = "127.0.0.1"
mysqldb   = "bhg_main"
mysqlport   = 3306
redisAddr = "127.0.0.1:6379"
```
在容器中运行，可以改成host.docker.internal(宿主机)


- env：
```
GO111MODULE="on"
GOPROXY=https://goproxy.cn,direct
```


- tool：
```
go get -u github.com/beego/bee/v2
```


- install：
```
go get github.com/cbirdcn/bee_http_game@v0.1.0
cd bee_http_game\@v0.1.0/
bee run
```


- result：

run成功提示：

http服务运行在http://127.0.0.1:8081

admin服务运行在http://127.0.0.1:8088


- errors:

如果依赖包run有错误，就：
```
go mod tidy
```
也可能是无法连接db或redis提示connection error


- test:

crud rest api
```
curl -X POST http://127.0.0.1:8081/v1/server -d '{"capacity":10000, "showStatus": 1,  "realStatus": 1}' --header "Content-Type: application/json"
curl http://127.0.0.1:8081/v1/server
curl http://127.0.0.1:8081/v1/server/10
curl -X PUT http://127.0.0.1:8081/v1/server/10 -d '{"capacity":5000}' --header "Content-Type: application/json"
curl -X DELETE http://127.0.0.1:8081/v1/server/10
```


- other:

生成代码
```
bee generate appcode -conn="root:root@tcp(127.0.0.1:3306)/bhg_main" -level=3
```
>level2=model+controller

>level3=level2+router

自动生成代码POST请求无法填充自增id，需要自己补充


- swagger

swagger自动生成是在GOPATH环境下运行的，用go mod创建项目，会导致
```
bee run -gendoc=true -downdoc-true
```
运行时，提示

> Failed to generate the docs.

需要解决可以搬到gopath下，或者参考：

ref:https://www.jianshu.com/p/e91e6244ae5c
