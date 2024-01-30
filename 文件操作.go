package e

import "os"

/*
打开方式
1、#读入：从指定文件读入数据，如果该文件不存在则失败；
2、#写出：写出数据到指定文件，如果该文件不存在则失败；
3、#读写：从文件中读入数据或者写出数据到文件，如果该文件不存在则失败；
4、#重写：写出数据到指定文件。如果该文件不存在则先创建一个新文件，如果已经存在就先清除其中的所有数据；
5、#改写：写出数据到指定文件。如果该文件不存在则创建一个新文件，如果已经存在就直接打开；
6、#改读：从文件中读入数据或者写出数据到文件。如果该文件不存在则创建一个新文件，如果已经存在就直接打开
其他数字、#读入
共享方式
1、#无限制：允许其它进程任意读写此文件
2、#禁止读：禁止其它进程读此文件
3、#禁止写：禁止其它进程写此文件
4、#禁止读写：禁止其它进程读写此文件
*/
func E_打开文件(文件路径 string, 打开方式 int, 共享方式 int) (*os.File, error) {
	switch 打开方式 {
	case 1:
		打开方式 = os.O_RDONLY
	case 2:
		打开方式 = os.O_WRONLY
	case 3:
		打开方式 = os.O_RDWR
	case 4:
		打开方式 = os.O_CREATE | os.O_WRONLY | os.O_TRUNC
	case 5:
		打开方式 = os.O_CREATE | os.O_WRONLY
	case 6:
		打开方式 = os.O_CREATE | os.O_RDONLY
	default:
		打开方式 = 0
	}
	switch 共享方式 {
	case 1:
		return os.OpenFile(文件路径, 打开方式, 0777)
	case 2:
		return os.OpenFile(文件路径, 打开方式, 0777)
	default:
		return os.OpenFile(文件路径, 打开方式, 0777)
	}

}

/*
返回移动都的文件位置
起始起始移动位置
1、#文件首； 2、#文件尾； 3、#现行位置
返回新的偏移量（相对开头）和可能的错误
*/
func E_移动读写位置(欲进行操作的文件指针 *os.File, 起始移动位置 int, 移动距离 int64) (int64, error) {
	switch 起始移动位置 {
	case 1:
		return 欲进行操作的文件指针.Seek(移动距离, 0)
	case 2:
		return 欲进行操作的文件指针.Seek(移动距离, 1)
	case 3:
		return 欲进行操作的文件指针.Seek(移动距离, 2)
	default:
		return 欲进行操作的文件指针.Seek(移动距离, 0)
	}

}

// 释放文件数据，关闭后无法再操作
func E_关闭文件(欲进行操作的文件指针 *os.File) error {
	return 欲进行操作的文件指针.Close()
}
