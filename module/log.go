package module

import (
	"io"
	"os"
	"time"
)

//日志保存路径和格式

const (
	LOGPATH  = "./log/"
	FORMAT   = "20060102"
	LineFeed = "\r\n"
)

//日志内容写入函数
func WriteLog(fileName, msg string) error {
	path := LOGPATH + time.Now().Format(FORMAT) + "/"
	if !IsExist(path) {
		return CreateDir(path)
	}
	var (
		err error
		f   *os.File
	)
	f, err = os.OpenFile(path+fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	_, err = io.WriteString(f, msg+LineFeed)
	defer f.Close()
	return err
}

//日志文件路径创建
func CreateDir(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	_ = os.Chmod(path, os.ModePerm)
	return nil
}

//日志文件是否存在
func IsExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}
