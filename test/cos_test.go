package test

import (
	"bytes"
	"context"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"leexsh/file_sys/core/define"
	"net/http"
	"net/url"
	"os"
	"testing"
)

func TestCos(t *testing.T) {
	u, _ := url.Parse("https://my-storage-1306331535.cos.ap-nanjing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv(define.CloudId),
			SecretKey: os.Getenv(define.CloudKey),
		},
	})

	key := "mystorage/go_book.pdf"

	_, _, err := client.Object.Upload(
		context.Background(), key, "./test_file/go_book.pdf", nil,
	)
	if err != nil {
		panic(err)
	}
}

// 分片上传的准备
func TestInitChunkUpload(t *testing.T) {
	u, _ := url.Parse("https://my-storage-1306331535.cos.ap-nanjing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv(define.CloudId),
			SecretKey: os.Getenv(define.CloudKey),
		},
	})
	key := "mystorage/go_book_1.pdf"
	v, _, err := client.Object.InitiateMultipartUpload(context.Background(), key, nil)
	if err != nil {
		t.Fatal(err)
	}
	UploadID := v.UploadID
	fmt.Println(UploadID)
}

// 分片上传
func TestChunkUpload(t *testing.T) {
	u, _ := url.Parse("https://my-storage-1306331535.cos.ap-nanjing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv(define.CloudId),
			SecretKey: os.Getenv(define.CloudKey),
		},
	})
	key := "mystorage/go_book_1.pdf"
	UploadID := "16598800825404feb5f2b7156d4b858d44bbbdcf41ee932df9c1f9e207fd1145bec0a49e80"
	f, err := os.ReadFile("test_file/0.chunk")
	// opt可选
	resp, err := client.Object.UploadPart(
		context.Background(), key, UploadID, 1, bytes.NewBuffer(f), nil,
	)
	if err != nil {
		t.Fatal(err)
	}
	PartETag := resp.Header.Get("ETag")
	fmt.Println(PartETag)
}

// 分片上传的结束
func TestChunkUploadFinish(t *testing.T) {
	u, _ := url.Parse("https://my-storage-1306331535.cos.ap-nanjing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv(define.CloudId),
			SecretKey: os.Getenv(define.CloudKey),
		},
	})
	key := "mystorage/go_book_1.pdf"
	UploadID := "16598800825404feb5f2b7156d4b858d44bbbdcf41ee932df9c1f9e207fd1145bec0a49e80"
	PartETag := "51d4e8b136b90ee94fe9eb891a60a93d"

	opt := &cos.CompleteMultipartUploadOptions{}
	opt.Parts = append(opt.Parts, cos.Object{
		PartNumber: 1, ETag: PartETag},
	)
	_, _, err := client.Object.CompleteMultipartUpload(
		context.Background(), key, UploadID, opt,
	)
	if err != nil {
		t.Fatal(err)
	}
}
