package split

import "strings"

func Split(s string) []string {
	command, argstr, _ := strings.Cut(s, " ")

	var singleQuote bool
	var doubleQuote bool
	var backslash bool
	var arg string
	var args []string
	for _, r := range argstr {
		switch r {
		case '\'':
			if backslash && doubleQuote {
				arg += "\\"
			}
			if backslash || doubleQuote {
				arg += string(r)
			} else {
				singleQuote = !singleQuote
			}
			backslash = false
		case '"':
			if backslash || singleQuote {
				arg += string(r)
			} else {
				doubleQuote = !doubleQuote
			}
			backslash = false
		case '\\':
			if backslash || singleQuote {
				arg += string(r)
				backslash = false
			} else {
				backslash = true
			}
		case ' ':
			if backslash && doubleQuote {
				arg += "\\"
			}
			if backslash || singleQuote || doubleQuote {
				arg += string(r)
			} else if arg != "" {
				args = append(args, arg)
				arg = ""
			}
			backslash = false
		default:
			if doubleQuote && backslash {
				arg += "\\"
			}
			arg += string(r)
			backslash = false
		}
	}

	if arg != "" {
		args = append(args, arg)
	}

	return append([]string{command}, args...)
}
