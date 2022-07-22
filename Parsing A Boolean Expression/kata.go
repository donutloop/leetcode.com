package kata

func parseBoolExpr(expression string) bool {
	v, _, _ := parseUnit(expression, 0)
	return v
}

func IsBool(c byte) bool {
	return c == 'f' || c == 't'
}

func parseBool(c byte) bool {
	if c == 'f' {
		return false
	}
	return true
}

func IsNot(c byte) bool {
	return c == '!'
}

func parseNot(expression string, i int) (bool, int) {
	var value bool
	var j int
	for ; i < len(expression); i++ {
		if IsClosed(expression[i]) {
			j = i
			break
		}

		v, k, match := parseUnit(expression, i)
		if k != -1 {
			i = k
		}
		if match {
			value = v
		}
	}
	return !value, j
}

func IsAnd(c byte) bool {
	return c == '&'
}

func parseAnd(expression string, i int) (bool, int) {
	var value *bool
	var j int
	for ; i < len(expression); i++ {
		if IsClosed(expression[i]) {
			j = i
			break
		}

		v, k, match := parseUnit(expression, i)
		if k != -1 {
			i = k
		}
		if match && value == nil {
			value = new(bool)
			*value = v
			match = false
		} else if match {
			*value = *value && v
			match = false
		}

	}
	return *value, j
}

func IsOr(c byte) bool {
	return c == '|'
}

func parseOr(expression string, i int) (bool, int) {
	var value *bool
	var j int
	for ; i < len(expression); i++ {
		if IsClosed(expression[i]) {
			j = i
			break
		}

		v, k, match := parseUnit(expression, i)
		if k != -1 {
			i = k
		}

		if match && value == nil {
			value = new(bool)
			*value = v
			match = false
		} else if match {
			*value = *value || v
			match = false
		}

	}
	return *value, j
}

func parseUnit(expression string, i int) (bool, int, bool) {
	var v bool
	k := -1
	var match bool
	if IsOr(expression[i]) {
		v, k = parseOr(expression, i+1)
		match = true
	} else if IsAnd(expression[i]) {
		v, k = parseAnd(expression, i+1)
		match = true
	} else if IsNot(expression[i]) {
		v, k = parseNot(expression, i+1)
		match = true
	} else if IsBool(expression[i]) {
		v = parseBool(expression[i])
		match = true
		k = -1
	}
	return v, k, match
}

func IsClosed(c byte) bool {
	return c == ')'
}
