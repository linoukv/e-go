package e

import (
	"bytes"
	"fmt"
)

type E_字节集 []byte

func E_取字节集长度(字节集数据 []byte) int {
	return len(字节集数据)
}
func E_到字节集(欲转换为字节集的数据 interface{}) []byte {
	return []byte(fmt.Sprintf("%v", 欲转换为字节集的数据))
}
func E_取字节集左边(欲取其部分的字节集 []byte, 欲取出字节的数目 int) []byte {
	return 欲取其部分的字节集[:欲取出字节的数目]
}
func E_取字节集右边(欲取其部分的字节集 []byte, 欲取出字节的数目 int) []byte {
	return 欲取其部分的字节集[len(欲取其部分的字节集)-欲取出字节的数目:]
}

/*
0为首位置，1为第2个位置，如此类推
*/
func E_取字节集中间(欲取其部分的字节集 []byte, 起始取出位置 int, 欲取出字节的数目 int) []byte {
	return 欲取其部分的字节集[起始取出位置 : 起始取出位置+欲取出字节的数目]
}

/*
0为首位置，1为第2个位置，如此类推
*/
func E_寻找字节集(被搜寻的字节集 []byte, 欲寻找的字节集 []byte, 起始搜寻位置 int) int {
	return bytes.Index(被搜寻的字节集[起始搜寻位置:], 欲寻找的字节集)
}
