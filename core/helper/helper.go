package helper

import (
	"bytes"
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	uuid2 "github.com/hashicorp/go-uuid"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io"
	"leexsh/file_sys/core/define"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

// 返回一个32位md5加密后的字符串
func Md5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

func GenerateToken(id int64, identity, name string, second int64) (string, error) {
	uc := define.UserClaim{
		Id:       id,
		Identity: identity,
		Name:     name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(second)).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	tokenString, err := token.SignedString([]byte(define.JwtKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func AnalyzeToke(token string) (*define.UserClaim, error) {
	uc := &define.UserClaim{}
	claims, err := jwt.ParseWithClaims(token, uc, func(token *jwt.Token) (interface{}, error) {
		return []byte(define.JwtKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return nil, errors.New("token is invalid")
	}
	return uc, err
}

func SendEmailCode() string {
	return "123456"
}

func GenerateEmailCode() string {
	str := "1234567890"
	code := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < define.EmailCodeLen; i++ {
		code += string(str[rand.Intn(len(str))])
	}
	return code
}

func GenerateUUID() string {
	str, err := uuid2.GenerateUUID()
	if err != nil {
		return ""
	}
	return str[0:15]
}

// upload file to COS
func UploadCos(req *http.Request) (string, error) {
	u, _ := url.Parse(define.COSADDR)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv(define.CloudId),
			SecretKey: os.Getenv(define.CloudKey),
		},
	})

	file, fileHeader, err := req.FormFile("file")
	key := "mystorage/" + GenerateUUID() + path.Ext(fileHeader.Filename)

	_, err = client.Object.Put(
		context.Background(), key, file, nil,
	)
	if err != nil {
		panic(err)
	}
	return define.COSADDR + "/" + key, nil
}

func CosInitPart(ext string) (string, string, error) {
	u, _ := url.Parse(define.COSADDR)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv(define.CloudId),
			SecretKey: os.Getenv(define.CloudKey),
		},
	})
	key := "mystorage/" + GenerateUUID() + "." + ext
	v, _, err := client.Object.InitiateMultipartUpload(context.Background(), key, nil)
	if err != nil {
		return "", "", err
	}
	return key, v.UploadID, nil
}

func CosPartUpload(r *http.Request) (string, error) {
	u, _ := url.Parse(define.COSADDR)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv(define.CloudId),
			SecretKey: os.Getenv(define.CloudKey),
		},
	})
	key := r.PostForm.Get("key")
	UploadID := r.PostForm.Get("uploadId")
	partNumber, err := strconv.Atoi(r.PostForm.Get("partNumber"))
	f, _, err := r.FormFile("file")
	if err != nil {
		return "", err
	}
	buf := bytes.NewBuffer(nil)
	io.Copy(buf, f)
	resp, err := client.Object.UploadPart(
		context.Background(), key, UploadID, partNumber, bytes.NewReader(buf.Bytes()), nil,
	)
	if err != nil {
		return "", nil
	}
	return strings.Trim(resp.Header.Get("ETag"), "\""), nil
}

// 分片上传的结束
func CosChunkUploadFinish(key, uploadId string, co []cos.Object) error {
	u, _ := url.Parse(define.COSADDR)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv(define.CloudId),
			SecretKey: os.Getenv(define.CloudKey),
		},
	})

	opt := &cos.CompleteMultipartUploadOptions{}
	opt.Parts = append(opt.Parts, co...)
	_, _, err := client.Object.CompleteMultipartUpload(
		context.Background(), key, uploadId, opt,
	)
	return err
}
