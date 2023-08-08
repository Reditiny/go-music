package utils

import "github.com/jinzhu/copier"

// CopyOne T V 要求为指针类型 TODO 错误处理
func CopyOne[T any, V any](in T) (out V) {
	var tar V
	copier.Copy(&tar, &in)
	return tar
}

func CopyList[T any, V any](in []T) (out []V) {
	var tar []V
	copier.Copy(&tar, &in)
	return tar
}
