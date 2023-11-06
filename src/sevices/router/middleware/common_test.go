package middleware

import (
	"strings"
	"testing"
)

// 提取大括号内的变量值
func TestReplaceVal(t *testing.T) {
	str := "{{1}}"
	Split(str, t)
	str = "1{{2}}3"
	Split(str, t)
	str = "{{1}}2{{3}}"
	Split(str, t)
	str = "1{{2}}3{{4}}"
	Split(str, t)
	str = "1{{2}}3{{4}}5"
	Split(str, t)
}

func Split(str string, t *testing.T) {
	t.Log("---------")
	strS := strings.Split(str, "{{")
	for _, i := range strS {
		if strings.Index(i, "}}") > -1 {
			TableStr := i[:strings.Index(i, "}}")]
			t.Log(TableStr)
		}
	}
}
