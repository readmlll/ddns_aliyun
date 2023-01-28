package common

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"runtime"
)

type SysUtils struct {
}

func (that SysUtils) GetPid() string {
	return _StrUtils.Int2String(os.Getpid())
}

func (that SysUtils) GetGid() string {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	return string(b)
}

func (that SysUtils) CmdExec(cmdStr string, arg ...string) string {

	cmd := exec.Command(cmdStr, arg...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
		return ""
	}
	return string(out)
}
