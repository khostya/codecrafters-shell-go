package split

import "strings"

func Split(s string) []string {
	res := make([]string, 0)

	const (
		opch   = "ch"
		zero   = "zero"
		single = "'"
		double = "\""
	)

	openCH := zero
	startIdx := 0
	isOpen := false
	current := strings.Builder{}

	for i, ch := range s {
		_ = string(ch)

		needWrite := false

		if !isOpen {
			if ch == ' ' {
				continue
			}
			isOpen = true
			if ch == '\'' {
				openCH = single
			} else if ch == '"' {
				openCH = double
			} else {
				openCH = opch
				needWrite = true
			}
			startIdx = i
		} else {
			if ch == ' ' && openCH == opch {
				isOpen = false
				res = append(res, current.String())
			} else if ch == '\'' && openCH == single {
				if startIdx != 0 && s[startIdx-1] == '\'' {

				} else if i < len(s)-1 && s[i+1] == '\'' {
					startIdx = i + 1
				} else {
					isOpen = false
					res = append(res, current.String())
				}
			} else if ch == '"' && openCH == double {
				if startIdx != 0 && s[startIdx-1] == '"' {

				} else if i < len(s)-1 && s[i+1] == '"' {
					startIdx = i + 1
				} else {
					isOpen = false
					res = append(res, current.String())
				}
			} else {
				needWrite = true
			}
		}

		if needWrite {
			current.WriteRune(ch)
		}

		if !isOpen {
			openCH = zero
			current = strings.Builder{}
		}
	}

	if current.Len() != 0 {
		res = append(res, current.String())
	}
	return res
}
