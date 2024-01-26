package e

import (
	"fmt"
	"golang.org/x/text/width"
	"strings"
	"unsafe"
)

// 定义一个名为E_取文本左边的函数，接收两个参数：欲取其部分的文本和欲取出字符的数目
func E_取文本左边(欲取其部分的文本 string, 欲取出字符的数目 int) string {
	// 如果欲取其部分的文本的长度小于等于欲取出字符的数目
	if len(欲取其部分的文本) <= 欲取出字符的数目 {
		// 返回欲取其部分的文本的全部内容
		return 欲取其部分的文本[:len(欲取其部分的文本)]
	} else {
		// 否则，返回欲取其部分的文本的前N个字符，其中N为欲取出字符的数目
		return 欲取其部分的文本[:欲取出字符的数目]
	}
}

// 定义一个函数E_取文本长度，接收一个字符串参数文本数据，返回该字符 串的长度
func E_取文本长度(文本数据 string) int {
	return len(文本数据)
}

// 定义一个函数E_取文本右边，接收两个参数：欲取其部分的文本和欲取出字符的数目中文3字节
func E_取文本右边(欲取其部分的文本 string, 欲取出字符的数目 int) string {
	// 获取欲取其部分的文本的长度
	长度 := len(欲取其部分的文本)
	// 如果长度小于等于欲取出字符的数目，则返回整个文本
	if 长度 <= 欲取出字符的数目 {
		return 欲取其部分的文本[0:]
	} else {
		return 欲取其部分的文本[长度-欲取出字符的数目:]
	}

}

// 定义一个函数E_取文本中间，接收三个参数：欲取其部分的文本、起始取出位置和欲取出字符的数目
func E_取文本中间(欲取其部分的文本 string, 起始取出位置 int, 欲取出字符的数目 int) string {
	// 返回从起始取出位置开始，长度为欲取出字符的数目的子字符串
	return 欲取其部分的文本[起始取出位置 : 起始取出位置+欲取出字符的数目]
}

// 定义一个名为E_多字符的函数，接收一个字节切片类型的参数欲取其字符的字符代码
func E_多字符(欲取其字符的字符代码 []byte) string {
	// 使用unsafe包中的Pointer函数将字节切片转换为指针类型，然后解引用得到字符串指针
	// 最后将字符串指针转换为字符串并返回
	return *(*string)(unsafe.Pointer(&欲取其字符的字符代码))
}

// 定义一个名为E_字符的函数，接收一个字节类型的参数欲取其字符的字符代码
func E_字符(欲取其字符的字符代码 byte) string {
	// 将字节类型的参数转换为字符串类型并返回
	return string(欲取其字符的字符代码)
}

// 定义一个名为E_取代码的函数，接收两个参数：欲取字符代码的文本和欲取其代码的字符位置
func E_取代码(欲取字符代码的文本 string, 欲取其代码的字符位置 int) uint8 {
	// 返回欲取字符代码的文本中指定位置的字符的ASCII码值
	return 欲取字符代码的文本[欲取其代码的字符位置]
}

// 定义一个名为E_寻找文本的函数，接收四个参数：被搜寻的文本、欲寻找的文本、起始位置和是否不区分大小写,起始位置从0开始起算
func E_寻找文本(被搜寻的文本 string, 欲寻找的文本 string, 起始位置 int, 是否不区分大小写 bool) int {
	if 是否不区分大小写 == true {
		被搜寻的文本 = E_到大写(被搜寻的文本)
		欲寻找的文本 = E_到大写(欲寻找的文本)
	}
	if 起始位置 == 0 {
		return strings.Index(被搜寻的文本, 欲寻找的文本)
	} else if strings.Index(被搜寻的文本[起始位置:], 欲寻找的文本) == -1 {
		return strings.Index(被搜寻的文本[起始位置:], 欲寻找的文本)
	} else {
		return strings.Index(被搜寻的文本[起始位置:], 欲寻找的文本) + 起始位置
	}
}

// 将大写字母转换为小写
func E_到小写(欲变换的文本 string) string {
	return strings.ToLower(欲变换的文本)
}

// 将小写字母转换为大写
func E_到大写(欲变换的文本 string) string {
	return strings.ToUpper(欲变换的文本)
}

// 返回一个整数值，指定一文本在另一文本中最后出现的位置，位置值从 1 开始。
func E_倒找文本(被搜寻的文本 string, 欲寻找的文本 string, 起始位置 int, 是否不区分大小写 bool) int {
	if 是否不区分大小写 == true {
		被搜寻的文本 = E_到大写(被搜寻的文本)
		欲寻找的文本 = E_到大写(欲寻找的文本)
	}
	if 起始位置 == 0 {
		return strings.LastIndex(被搜寻的文本, 欲寻找的文本)
	} else if strings.LastIndex(被搜寻的文本[起始位置:], 欲寻找的文本) == -1 {
		return strings.LastIndex(被搜寻的文本[起始位置:], 欲寻找的文本)
	} else {
		return strings.LastIndex(被搜寻的文本[起始位置:], 欲寻找的文本) + 起始位置
	}
}
func E_到全角(欲变换的文本 string) string {
	return width.Widen.String(欲变换的文本)
}

// 可以把全角转换成半角
func E_到半角(欲变换的文本 string) string {
	return width.Narrow.String(欲变换的文本)
}

// 返回一个文本，代表指定数值、逻辑值或日期时间被转换后的结果。如果为文本数据，将被直接返回
func E_到文本(待转换的数据 interface{}) string {
	return fmt.Sprint(待转换的数据)
}

// 返回一个删除前空文本
func E_删首空(欲删除空格的文本 string) string {
	return strings.TrimLeft(欲删除空格的文本, " ")
}

// 返回一个删除后空文本
func E_删尾空(欲删除空格的文本 string) string {
	return strings.TrimRight(欲删除空格的文本, " ")
}

// 返回一个文本，其中包含被删除了首部及尾部全角或半角空格的指定文本
func E_删首尾空(欲删除空格的文本 string) string {
	return strings.TrimSpace(欲删除空格的文本)
}
func E_删全部空(欲删除空格的文本 string) string {
	return strings.Replace(欲删除空格的文本, " ", "", -1)
}

// 替换的起始位置，0为首位置，1为第2个位置，如此类推
func E_文本替换(欲被替换的文本 string, 起始替换位置 int, 替换长度 int, 用作替换的文本 string) string {
	return 欲被替换的文本[:起始替换位置] + 用作替换的文本 + 欲被替换的文本[起始替换位置+替换长度:]
}

/*
进行替换的起始位置从0开始，0为第一个，1为第二个；
替换进行的次数-1为全部替换
进行替换的起始位置参数值指定被替换子文本的起始搜索位置0为第一个位置
*/
func E_子文本替换(欲被替换的文本 string, 欲被替换的子文本 string, 用作替换的子文本 string, 进行替换的起始位置 int, 替换进行的次数 int, 是否区分大小写 bool) string {
	if 替换进行的次数 == 0 {
		return 欲被替换的文本
	}
	if 是否区分大小写 == true {
		return strings.Replace(欲被替换的文本[进行替换的起始位置:], 欲被替换的子文本, 用作替换的子文本, 替换进行的次数)
	}
	i := strings.Count(E_到大写(欲被替换的文本), E_到大写(欲被替换的子文本))
	临时文本 := ""
	j := 0
	k := 0
	for n := 1; n <= i; n++ {
		k = strings.Index(E_到大写(欲被替换的文本[k:]), E_到大写(欲被替换的子文本))
		j += k + len(用作替换的子文本)
		if k == -1 && n == 1 {
			return 欲被替换的文本
		} else {
			if n == 1 {
				临时文本 += E_文本替换(欲被替换的文本[0:], k, len(欲被替换的子文本), 用作替换的子文本)
			} else {
				临时文本 += E_文本替换(欲被替换的文本[j:], k, len(欲被替换的子文本), 用作替换的子文本)
			}

		}
		if 进行替换的起始位置 != -1 && 替换进行的次数 == n {
			return 临时文本
		}
	}
	return 临时文本
}

// 返回具有指定数目半角空格的文本。
func E_取空白文本(重复次数 int) string {
	c := ""
	for i := 1; i <= 重复次数; i++ {
		c += " "
	}
	return c
}

// 返回一个文本，其中包含指定次数的文本重复结果
func E_取重复文本(重复次数 int, 待重复文本 string) string {
	c := ""
	for i := 1; i <= 重复次数; i++ {
		c += 待重复文本
	}
	return c
}

// 如果返回值小于0，表示文本一小于文本二；如果等于0，表示文本一等于文本二；如果大于0，表示文本一大于文本二。
func E_文本比较(待比较文本一 string, 待比较文本二 string, 是否区分大小写 bool) int {
	if 是否区分大小写 == true {
		待比较文本一 = E_到大写(待比较文本一)
		待比较文本二 = E_到大写(待比较文本二)
	}
	return strings.Compare(待比较文本一, 待比较文本二)
}

// 将指定文本进行分割，返回分割后的一维文本数组
func E_分割文本(待分割文本 string, 用作分割的文本 string) []string {
	return strings.Split(待分割文本, 用作分割的文本)
}

// 定义一个名为E_指针到文本的函数，接收一个指向字符串的指针作为参数
func E_指针到文本(内存文本指针 *string) string {
	// 返回该指针所指向的字符串的值
	return *内存文本指针
}
