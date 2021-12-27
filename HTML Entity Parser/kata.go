package kata

var token = map[string]string{"&quot;": "\"", "&apos;": "'", "&amp;": "&", "&gt;": ">", "&lt;": "<", "&frasl;": "/"}

func entityParser(text string) string {
	output := make([]byte, 0)
	tag := make([]byte, 0)
	for i := 0; i < len(text); i++ {
		if text[i] == ';' {
			tag = append(tag, text[i])
			v, ok := token[string(tag)]
			if ok {
				output = append(output, v...)
			} else {
				output = append(output, string(tag)...)
			}
			tag = make([]byte, 0)
			continue
		}
		if text[i] == '&' {
			tag = append(tag, text[i])
			continue
		}
		if len(tag) != 0 {
			tag = append(tag, text[i])
		} else {
			output = append(output, text[i])
		}
	}

	if len(tag) != 0 {
		output = append(output, tag...)
	}

	return string(output)
}
