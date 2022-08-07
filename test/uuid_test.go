package test

import (
	"fmt"
	uuid2 "github.com/hashicorp/go-uuid"
	"testing"
)

func TestUUID(t *testing.T) {
	uuid, err := uuid2.GenerateUUID()
	fmt.Println(uuid)
	fmt.Println(err)
}
