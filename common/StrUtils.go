package common

import (
	"strconv"
	"strings"
)

type StrUtils struct {
}

func (that StrUtils) Str2Int(str string) (int, error) {
	i, err := strconv.Atoi(str)
	return i, err
}
func (that StrUtils) Str2Int64(str string) (int64, error) {
	i, err := strconv.ParseInt(str, 10, 64)
	return i, err
}
func (that StrUtils) Int2String(n int) string {
	str := strconv.Itoa(n)
	return str
}
func (that StrUtils) Int64ToString(n int64) string {
	str := strconv.FormatInt(n, 10)
	return str
}

func (that StrUtils) Trim(_str string) string {
	return strings.TrimSpace(_str)
}

func (that StrUtils) IsEmpty(_str string) bool {
	_str = that.Trim(_str)
	if _str == "" && len(_str) == 0 {
		return true
	}
	return false
}
func (that StrUtils) Replace(_str string, old string, new string, count int) string {

	return strings.Replace(_str, old, new, count)
}

func (that StrUtils) ToUpperCase(_str string) string {
	return strings.ToUpper(_str)
}
func (that StrUtils) ToLowerCase(_str string) string {
	return strings.ToLower(_str)
}
