package mocks

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

var prefix = []string{"130", "131", "132", "133", "134", "135", "136", "137", "138", "139", "150", "151", "152", "153", "154", "155", "156", "157", "158", "159", "170", "171", "172", "173", "174", "175", "176", "177", "178", "180", "181", "182", "183", "184", "185", "186", "187", "188", "189"}

// RandomPhone 随机生成手机号码
func RandomPhone() string {
	max := big.NewInt(int64(len(prefix)))
	index, err := rand.Int(rand.Reader, max)
	if err != nil {
		panic(err)
	}
	phone := prefix[index.Int64()]

	for i := 0; i < 8; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(10))
		if err != nil {
			panic(err)
		}
		phone += fmt.Sprintf("%d", num.Int64())
	}

	return phone
}
