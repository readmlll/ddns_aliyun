package common

import (
	"fmt"
	"os"
)

var _RunTime = RunTime{}
var _DateUtils = DateUtils{}
var _StrUtils = StrUtils{}
var _FileUtils = FileUtils{}

func (that LogUtils) PrintLn(msg string) string {
	fun_name, _, line, _ := _RunTime.GetRunFuncName(1)
	log_str := fmt.Sprintf("%s %s %s --- [%s - %s] : %s",
		_DateUtils.GetCommonTimeStrAndMillisecond(), "INFO", _RunTime.GetPid(),
		fun_name, _StrUtils.Int2String(line), msg)
	fmt.Println(log_str)
	return log_str
}

func (that *LogUtils) commonWriteString2File(_type string, text string) (n int, err error) {

	that.checkDate()
	var pathf *os.File
	switch _type {
	case "Info":
		pathf = that.InfoFilePath
		break
	case "Error":
		pathf = that.ErrorFilePath
		break
	case "Debug":
		pathf = that.DebugFilePath
		break
	}
	n, err = _FileUtils.WriteString(pathf, text+"\n\r")
	return n, err
}

func (that *LogUtils) Info(msg string) string {
	if that.level == "NONE" {
		return ""
	}
	if that.level == "ERROR" {
		return ""
	}
	fun_name, _, line, _ := _RunTime.GetRunFuncName(1)
	log_str := fmt.Sprintf("%s %s %s --- [%s - %s] : %s",
		_DateUtils.GetCommonTimeStrAndMillisecond(), "INFO", _RunTime.GetPid(),
		fun_name, _StrUtils.Int2String(line), msg)
	fmt.Println(log_str)

	that.commonWriteString2File("Info", log_str)

	return log_str
}
func (that *LogUtils) Error(msg string) string {
	if that.level == "NONE" {
		return ""
	}

	fun_name, _, line, _ := _RunTime.GetRunFuncName(1)
	log_str := fmt.Sprintf("%s %s %s --- [%s - %s] : %s",
		_DateUtils.GetCommonTimeStrAndMillisecond(), "ERROR", _RunTime.GetPid(),
		fun_name, _StrUtils.Int2String(line), msg)
	fmt.Println(log_str)

	that.commonWriteString2File("Error", log_str)
	return log_str
}

func (that *LogUtils) Debug(msg string) string {
	if that.level == "NONE" {
		return ""
	}
	fun_name, _, line, _ := _RunTime.GetRunFuncName(1)
	log_str := fmt.Sprintf("%s %s %s --- [%s - %s] : %s",
		_DateUtils.GetCommonTimeStrAndMillisecond(), "DEBUG", _RunTime.GetPid(),
		fun_name, _StrUtils.Int2String(line), msg)
	fmt.Println(log_str)

	that.commonWriteString2File("Debug", log_str)
	return log_str
}

func (that *LogUtils) checkDate() {
	old_date := that.cur_date
	cur_date := _DateUtils.GetCommonDate()
	if cur_date == old_date {
		return
	}
	err_new_file := _StrUtils.Replace(that.ErrorFilePath.Name(), old_date, cur_date, -1)
	info_new_file := _StrUtils.Replace(that.InfoFilePath.Name(), old_date, cur_date, -1)
	debug_new_file := _StrUtils.Replace(that.DebugFilePath.Name(), old_date, cur_date, -1)

	that.ErrorFilePath.Close()
	that.InfoFilePath.Close()
	that.DebugFilePath.Close()

	that.ErrorFilePath = _FileUtils.OpenFileAppend(err_new_file, true)
	that.InfoFilePath = _FileUtils.OpenFileAppend(info_new_file, true)
	that.DebugFilePath = _FileUtils.OpenFileAppend(debug_new_file, true)
}

func (l LogUtils) NewLogUtils(infoFile string, errorFile string, debbugFile string, level string) (that *LogUtils) {
	that = new(LogUtils)
	that.cur_date = _DateUtils.GetCommonDate()
	if _StrUtils.IsEmpty(level) {
		level = "NONE"
	}
	level = _StrUtils.ToUpperCase(level)
	that.level = level
	if level == "NONE" {
		return that
	}
	if _StrUtils.IsEmpty(errorFile) {
		errorFile = "./" + that.cur_date + "/error_log.log"
		errorFile = _FileUtils.RelativePath2AbsPath(errorFile)
	}
	if _StrUtils.IsEmpty(debbugFile) {
		debbugFile = "./" + that.cur_date + "/debug_log.log"
		debbugFile = _FileUtils.RelativePath2AbsPath(debbugFile)
	}
	if _StrUtils.IsEmpty(infoFile) {
		infoFile = "./" + that.cur_date + "/info_log.log"
		infoFile = _FileUtils.RelativePath2AbsPath(infoFile)
	}

	that.ErrorFilePath = _FileUtils.OpenFileAppend(errorFile, true)
	that.InfoFilePath = _FileUtils.OpenFileAppend(infoFile, true)
	that.DebugFilePath = _FileUtils.OpenFileAppend(debbugFile, true)
	return that
}

type LogUtils struct {
	level         string
	cur_date      string
	InfoFilePath  *os.File
	ErrorFilePath *os.File
	DebugFilePath *os.File
}
