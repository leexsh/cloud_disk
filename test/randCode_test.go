package test

import (
	"leexsh/file_sys/core/helper"
	"testing"
)

func TestRandCode(t *testing.T) {
	println(helper.GenerateEmailCode())
}
