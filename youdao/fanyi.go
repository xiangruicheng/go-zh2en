package youdao

import (
	"strings"
)

func Fanyi(str string) string {
	resp1, err := CallSimple(str)
	if err == nil && resp1.Result.Code == 200 && resp1.Data.Entries[0].Explain != "" {
		onelen := strings.Index(resp1.Data.Entries[0].Explain, ";")
		if onelen > 0 {
			return resp1.Data.Entries[0].Explain[0:onelen]
		} else {
			return resp1.Data.Entries[0].Explain
		}
	} else {
		resp2, err := CallComplex(str)
		if err == nil && resp2.Fanyi.Tran != "" {
			// User's replace User
			if strings.Index(resp2.Fanyi.Tran, "'s") > 0 {
				resp2.Fanyi.Tran = strings.ReplaceAll(resp2.Fanyi.Tran, "'s", "")
			}
			return resp2.Fanyi.Tran
		}
	}
	return ""
}
