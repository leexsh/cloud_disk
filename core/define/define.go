package define

import "github.com/dgrijalva/jwt-go"

// jwt加密，如果是自定义结构体需要实现jwt.StandardClaims结构体，里面可以定义过期时间颁发者等等
type UserClaim struct {
	jwt.StandardClaims
	Id       int64
	Identity string
	Name     string
}

var JwtKey = "my_cloud_disk"

// 验证码长度
var EmailCodeLen = 6

// 验证码过期时间
var CodeExprie = 300

// 腾讯云
var CloudKey = "TECENTCLOUDSECRETKEY"
var CloudId = "TECENTCLOUDSECRETID"
var COSADDR = "https://my-storage-1306331535.cos.ap-nanjing.myqcloud.com"

// 默认的分页
var Pagesize int = 20

var DateTime = "2006-01-02 15:04:5"

var TokenExpire int64 = 3600 * 12
var RefreshTokenExpire int64 = 3600 * 24
