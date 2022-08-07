package test

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"testing"
)

var chunksize int = 1 * 1024 * 1024 // 1M

// 文件分片

func TestChunkFile(t *testing.T) {
	fileInfo, err := os.Stat("test_file/go_book.pdf")
	if err != nil {
		t.Fatal(err)
	}
	// 分片个数
	chunkNum := math.Ceil(float64(float64(fileInfo.Size()) / float64(chunksize)))
	myfile, err := os.OpenFile("test_file/go_book.pdf", os.O_RDONLY, 0666)
	if err != nil {
		t.Fatal(err)
	}
	b := make([]byte, chunksize)
	for i := 0; i < int(chunkNum); i++ {
		myfile.Seek(int64(i*chunksize), 0)
		if int64(chunksize) > fileInfo.Size()-int64(i*chunksize) {
			b = make([]byte, fileInfo.Size()-int64(i*chunksize))
		}
		myfile.Read(b)
		f, err := os.OpenFile("test_file/"+strconv.Itoa(i)+".chunk", os.O_CREATE|os.O_WRONLY, os.ModePerm)
		if err != nil {
			t.Fatal(err)
		}
		f.Write(b)
		f.Close()
	}
	myfile.Close()
}

func TestMergeFile(t *testing.T) {
	myfile, err := os.OpenFile("test_file/go_book_1.pdf", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
	fileInfo, err := os.Stat("test_file/go_book.pdf")
	if err != nil {
		t.Fatal(err)
	}

	chunkNum := math.Ceil(float64(float64(fileInfo.Size()) / float64(chunksize)))
	for i := 0; i < int(chunkNum); i++ {
		f, err := os.OpenFile("test_file/"+strconv.Itoa(i)+".chunk", os.O_RDONLY, os.ModePerm)
		if err != nil {
			t.Fatal(err)
		}
		b, err := ioutil.ReadAll(f)
		if err != nil {
			t.Fatal(err)
		}
		myfile.Write(b)
		f.Close()
	}
	myfile.Close()
}

func TestCheck(t *testing.T) {
	file1, err := os.OpenFile("test_file/go_book.pdf", os.O_RDONLY, 0666)
	if err != nil {
		t.Fatal(err)
	}
	b1, err := ioutil.ReadAll(file1)
	if err != nil {
		t.Fatal(err)
	}

	file2, err := os.OpenFile("test_file/go_book_1.pdf", os.O_RDONLY, 0666)
	if err != nil {
		t.Fatal(err)
	}
	b2, err := ioutil.ReadAll(file2)
	if err != nil {
		t.Fatal(err)
	}

	s1md5 := fmt.Sprintf("%x", md5.Sum(b1))
	s2md5 := fmt.Sprintf("%x", md5.Sum(b2))
	fmt.Println(s1md5)
	fmt.Println(s2md5)
}
