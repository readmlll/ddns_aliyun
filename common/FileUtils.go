package common

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	pathPkg "path"
	"path/filepath"
)

var _LogUtils LogUtils = LogUtils{}

type FileUtils struct {
}

func (that FileUtils) PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}

	//isnotexist来判断，是不是不存在的错误
	if os.IsNotExist(err) { //如果返回的错误类型使用os.isNotExist()判断为true，说明文件或者文件夹不存在
		return false
	}
	return false //如果有错误了，但是不是不存在的错误，所以把这个错误原封不动的返回
}

func (that FileUtils) IsDir(path string) bool {
	exist := that.PathExists(path)
	if !exist {
		return false
	}
	fi, e := os.Stat(path)
	if e != nil {
		return false
	}
	return fi.IsDir()
}

func (that FileUtils) IsFile(path string) bool {
	return that.IsDir(path)
}

// GetFileNameByPath Base 函数返回路径的最后一个元素。在提取元素前会去掉末尾的斜杠。
// 如果路径是 ""，会返回 "."；如果路径是只有一个斜杆构成的，会返回 "/"
func (that FileUtils) GetFileNameByPath(path string) string {
	return pathPkg.Base(path)
}

// GetDirPathByPath Dir() 和 Base() 函数将一个路径名字符串分解成目录和文件名两部分。
// Dir返回路径除去最后一个路径元素的部分，即该路径最后一个元素所在的目录。
// 在使用Split去掉最后一个元素后，会简化路径并去掉末尾的斜杠。
// 如果路径是空字符串，会返回"."；如果路径由1到多个斜杠后跟0到多个非斜杠字符组成，会返回"/"；其他任何情况下都不会返回以斜杠结尾的路径。
func (that FileUtils) GetDirPathByPath(path string) string {
	return pathPkg.Dir(path)
}

// GetFileTypeByPath Ext 函数返回 path 文件扩展名。扩展名是路径中最后一个从 . 开始的部分，包括 .。如果该元素没有 . 会返回空字符串。
func (that FileUtils) GetFileTypeByPath(path string) string {
	return pathPkg.Ext(path)
}

func (that FileUtils) JoinPath(paths ...string) string {
	return pathPkg.Join(paths...)
}

func (that FileUtils) CreateDirs(dirPath string) bool {
	if !that.PathExists(dirPath) {
		err := os.MkdirAll(dirPath, os.ModePerm)
		return err != nil
	}
	return true
}

func (that FileUtils) OpenFileAppend(path string, append bool) *os.File {
	var f *os.File
	var err error
	if that.PathExists(path) {
		if append {
			f, err = os.OpenFile(path, os.O_RDWR|os.O_APPEND, os.ModePerm) //打开文件
		} else {
			f, err = os.OpenFile(path, os.O_RDWR, os.ModePerm) //打开文件
		}

		if err != nil {
			_LogUtils.PrintLn(fmt.Sprintf("打开文件出错，文件存在 %s %s", path, err.Error()))
			return nil
		}
		return f
		//如果存在那就追加
	}

	//如果不存在则先检测目录是否存在
	dirPath := that.GetDirPathByPath(path)
	if !that.PathExists(dirPath) {
		b := that.CreateDirs(dirPath)
		if !b {
			_LogUtils.PrintLn(fmt.Sprintf("创建文件失败 %s , 目录不存在尝试创建也失败", path))
			return nil
		}
	}

	if append {
		f, err = os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm) //打开文件
	} else {
		f, err = os.OpenFile(path, os.O_CREATE|os.O_RDWR, os.ModePerm) //打开文件
	}

	if err != nil {
		_LogUtils.PrintLn(fmt.Sprintf("打开文件出错，文件不存在 %s %s", path, err.Error()))
		return nil
	}
	return f
}

func (that FileUtils) WriteString(f *os.File, text string) (n int, err error) {
	n, err = io.WriteString(f, text) //写入文件(字符串)
	return n, err
}

func (that FileUtils) writeString2File(path string, text string, append bool) bool {
	var f = that.OpenFileAppend(path, append)
	if f == nil {
		return false
	}
	defer f.Close()
	var err error
	_, err = that.WriteString(f, text) //写入文件(字符串)
	if err != nil {
		_LogUtils.PrintLn(fmt.Sprintf("写入文件出错，写入出错 %s %s", path, err.Error()))
		return false
	}
	return true
}

//判断是否为绝对路径
func (that FileUtils) IsAbsPath(path string) bool {
	if filepath.IsAbs(path) {
		return true
	}
	return false
}

//绝对路径转相对路径
func (that FileUtils) AbsPath2RelativePath(path string, path_join_flag string) string {
	path, _ = filepath.Rel(path, path_join_flag)
	return path
}

//相对路径转绝对路径
func (that FileUtils) RelativePath2AbsPath(path string) string {
	path, _ = filepath.Abs(path)
	return path
}

func (that FileUtils) ReadTextFileAll(path string) string {
	fi, _ := os.Open(path)
	text, err := ioutil.ReadAll(fi)
	if err != nil {
		log.Fatal(err)
	}

	return string(text)
}
