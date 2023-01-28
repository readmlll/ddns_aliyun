package common

import "runtime"

type RunTime struct {
}

var _SysUtils SysUtils = SysUtils{}

func (that RunTime) GetPid() string {
	return _SysUtils.GetPid()
}

func (that RunTime) GetGid() string {
	return _SysUtils.GetGid()
}

func (that RunTime) GetRunFuncName(skip_call_level int) (string, string, int, bool) {
	skip_call_level += 1
	pc, file, line, ok := runtime.Caller(skip_call_level)
	f := runtime.FuncForPC(pc)
	return f.Name(), file, line, ok
}
