# api-server-demo
Using go development language to develop server API interface
## 描述
* 用go开发语言开发服务器API接口
  
## 目录
- [x] 使用go modules 初始化项目
- [x] 启动一个最简单的 RESTful API 服务器
- [x] 配置文件读取 
- [x] 初始化Mysql数据库并建立连接
- [x] 自定义业务错误信息
- [x] 读取和返回HTTP请求
- [x] 用户业务逻辑处理（业务处理）
    - [x] 注册
    - [x] 登录
- [x] HTTP调用添加自定义处理逻辑
- [x] 路由中间件
    - [x] 路由请求校验 
    - [x] 分类使用方式
    - [x] 自定义中间件
    - [ ] 签名验证
        - [ ] 开启JWT认证
    - [x] 记录和管理API日志
    - [x] 异常捕获
    - [ ] 链路追踪
- [x] 参数验证(validator.v9)
    - [x] 模型绑定和验证
    - [x] 自定义验证器
- [x] 请求头
    - [x] 自定义请求头
    - [ ] Cookies应用
    - [ ] Session应用
- [x] API身份验证
- [ ] API性能分析
- [ ] 支持系统检测接口(如cpu,disk,health,memery)
- [ ] 给API命令增加版本号功能
- [ ] 生成Swagger在线文档
- [ ] 给API增加启动脚本
- [ ] 语言文件支持(zh and en)
- [ ] 支持优雅地重启或停止
- [x] 支持多种运行模式
- [x] 支持打包并压缩脚本

## 操作
* 安装
``` 
go get -u github.com/dsw0214/api-server-demo
```

* 初始化表格
    * 导入测试sql到数据库
     
* 修改配置
    * conf/config.yaml
    
* 直接运行
```
go run .
```

* 打包运行
```
go build .
```

* 打包并压缩脚本
```
sh build_compress.sh 
```

```
upx is /usr/local/bin/upx
***Use upx Compress beging***
                       Ultimate Packer for eXecutables
                          Copyright (C) 1996 - 2018
UPX 3.95        Markus Oberhumer, Laszlo Molnar & John Reiser   Aug 26th 2018

        File size         Ratio      Format      Name
   --------------------   ------   -----------   -----------
  15739628 ->   5836816   37.08%   macho/amd64   api-server-demo               

Packed 1 file.

```

## Api接口范例
* 注册

```
curl http://127.0.0.1:8888/register -X POST -d "Username=test01&Password=123456"
```
```  
{"code":20000,"message":"OK","data":{"username":"test01","message":"Register Success"}}
```
 
* 登录

```
curl http://127.0.0.1:8888/login -X POST -d "Username=test01&Password=123456"
```
```
{"code":20000,"message":"OK","data":{"hello":"Welcome","userName":"test01"}}
```

## 扩展包
* [Gin] is a web framework written in Go (Golang)
* [Viper] is Go configuration with fangs!
* [Gorm] is fantastic ORM library for Golang
* [log] is log package
* [Mysql] is a MySQL driver for Go's
* [Bcrypt] is a Encryption package
* [pprof] is HTTP server runtime profiling data in the format expected by the pprof visualization tool.

[Gin]: https://github.com/gin-gonic/gin "Gin is a web framework written in Go (Golang)"
[Viper]: https://github.com/spf13/viper "Go configuration with fangs!"
[Gorm]: https://github.com/jinzhu/gorm "The fantastic ORM library for Golang"
[log]: https://github.com/lexkong/log "Log库"
[Mysql]: https://github.com/go-sql-driver/mysql "mysql"
[Bcrypt]: https://golang.org/x/crypto/bcrypt "bcrypt"
[pprof]: https://github.com/gin-contrib/pprof "pprof"

## 国内代理
* goproxy
    * export GOPROXY="https://goproxy.cn"
* aliyun
    * export GOPROXY=https://mirrors.aliyun.com/goproxy/