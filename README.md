## 基于go-zero 和 xorm的云盘项目
>windows gozero:https://blog.csdn.net/mokyzone/article/details/124325021

- 使用到的命令
```go
goctl api new core
// get 
curl -i -X GET http://localhost:8888/from/you
// 使用api生成文件
goctl api go -api core.api -dir . -style go_zero
// go zero run
go run core.go  -f etc/core-api.yaml
```

### 模块
#### 用户
- 密码登录 Post

    localhost:8888/user/login
```go
req: name password
    
response: token
```
   


- 用户详细信息 Get

  localhost:8888/user/details

```go
req: identity
response: user.name, user.email
```

- 注册验证码发送 Post
  localhost:8888/mail/code/send/register
```go
// 验证码发送的库
go get  github.com/jordan-wright/email
```
```go
req: email
resp: code, msg
```

- 用户注册 Post

```go
localhost:8888/user/register
req: name, password, email, code
resp: msg
```

- 刷新鉴权


- 中间件

  - Auth 使用jwt进行鉴权

#### 文件 需要用户鉴权

- 文件上传

/file/upload

  1. 腾讯云对象存储

     腾讯云官网：https://cloud.tencent.com/

      sdk文档地址：https://cloud.tencent.com/document/product/436/65644

- 文件和对象进行关联
/user/repository/save

- 用户的文件列表
  /user/file/list post

- 用户文件名称修改

/user/file/update post

- 用户文件夹创建

/user/folder/create post


- 用户文件创建

/user/file/delete post

- 用户文件移动

/user/file/move put 

- 创建分享链接

/user/share/create post

- 获取分享链接的详情

/share/details get

- 资源保存

/file/shared/save

- 文件上传前的准备

/file/upload/prepare

- 文件上传

file/chunk/upload


- 文件上传成功

/file/chunk/upload/finish

### postman参考cloud_disk.postman_collection