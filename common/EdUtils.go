package common

import (
	"encoding/json"
)

type EdUtils struct {
}

func (that EdUtils) Str2jsonDict(jsonText string) (dat map[string]interface{}) {
	dat = make(map[string]interface{})
	if err := json.Unmarshal([]byte(jsonText), &dat); err == nil {
		return dat
	} else {
		println(err)
		return nil
	}
}
