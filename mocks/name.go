package mocks

import (
	"crypto/rand"
	"math/big"
)

// 姓氏列表
var surnames = []string{
	"赵", "钱", "孙", "李", "周", "吴", "郑", "王", "冯", "陈",
	"褚", "卫", "蒋", "沈", "韩", "杨", "朱", "秦", "尤", "许",
	"何", "吕", "施", "张", "孔", "曹", "严", "华", "金", "魏",
	"陶", "姜", "戚", "谢", "邹", "喻", "柏", "水", "窦", "章",
	"云", "苏", "潘", "葛", "奚", "范", "彭", "郎", "鲁", "韦",
	"昌", "俞", "任", "崔", "邬", "廖", "刁", "邱", "瞿", "殷",
	"梁", "罗", "毕", "郝", "骆", "贺", "项", "龚", "江", "童",
	"颜", "郭", "梅", "盛", "林", "刘", "田", "李", "黄", "温",
	"廉", "柳", "龙", "胡", "孟", "熊", "秋", "鹿", "史", "常",
}

func RandomChineseName() string {
	// 随机选择姓氏
	surname := surnames[randInt(0, len(surnames)-1)]

	// 随机生成名字的长度（2 或 3）
	nameLen := randInt(1, 2)

	// 随机生成名字的 Unicode 编码值
	nameCode := make([]rune, nameLen)
	for i := 0; i < nameLen; i++ {
		nameCode[i] = rune(randInt(19968, 20902))
	}

	// 将编码值转换为对应的中文字符
	name := string(nameCode)

	return surname + name
}

func randInt(min, max int) int {
	// 计算随机整数范围
	diff := big.NewInt(int64(max - min + 1))

	// 生成随机数
	n, err := rand.Int(rand.Reader, diff)
	if err != nil {
		panic(err)
	}

	// 转换为整数并返回
	return int(n.Int64()) + min
}
