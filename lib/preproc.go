package og

import (
	"strings"
)

func indentCount(str string) int {
	for i := range str {
		if str[i] != ' ' && str[i] != '\t' {
			return i
		}
	}
	return 0
}
func Preproc(str string) string {
	rawLines := strings.Split(str, "\n")
	lines := []string{}
	for _, v := range rawLines {
		if len(v) > 0 {
			lines = append(lines, v)
		}
	}
	res := []string{}
	lastIndent := 0
	indentSize := 0
	for i := range lines {
		indent := indentCount(lines[i])
		token := "=>"
		eqIf := "= if "
		if len(lines[i]) > indent+4 && lines[i][indent:indent+4] == "for " {
			lines[i] += ";"
		}
		if len(lines[i]) > indent+3 && lines[i][indent:indent+3] == "if " && !strings.Contains(lines[i], token) {
			lines[i] += ";"
		}
		if len(lines[i]) > indent+3 && strings.Contains(lines[i], eqIf) && !strings.Contains(lines[i], token) {
			lines[i] += ";"
		}
		if len(lines[i]) > indent+5 && lines[i][indent:indent+5] == "else " && !strings.Contains(lines[i], token) {
			lines[i] += ";"
		}
		if indentSize == 0 && indent != lastIndent {
			indentSize = indent
		}
		if indent > lastIndent {
			res[len(res)-1] = res[len(res)-1] + " {"
		} else if indent < lastIndent {
			indentBuff := lastIndent - indent
			for indentBuff > 0 {
				res = append(res, "}")
				indentBuff -= indentSize
			}
		}
		lastIndent = indent
		res = append(res, lines[i])
	}
	for lastIndent > 0 {
		res = append(res, "}")
		lastIndent -= indentSize
	}
	joined := strings.Join(res, "\n") + "\n"
	return joined
}
