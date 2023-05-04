// 随机邮箱生成

package mocks

import "math/rand"

func RandomEmail() string {
	return RandomString(10) + "@" + RandomString(5) + ".com"
}

func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
