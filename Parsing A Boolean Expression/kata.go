package kata

func parseBoolExpr(expression string) bool {
	return parse(expression, 0)
}

func parse(expression string, i int) bool {

	if IsNot(expression[i]) {
		v, _ := parseNot(expression, i+1)
		return v
	} else if IsAnd(expression[i]) {
		v, _ := parseAnd(expression, i+1)
		return v
	} else if IsOr(expression[i]) {
		v, _ := parseOr(expression, i+1)
		return v
	} else if IsBool(expression[i]) {
		return parseBool(expression[i])
	}

	return false
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
		if IsOr(expression[i]) {
			var k int
			value, k = parseOr(expression, i+1)
			i = k
		} else if IsAnd(expression[i]) {
			var k int
			value, k = parseAnd(expression, i+1)
			i = k
		} else if IsNot(expression[i]) {
			var k int
			value, k = parseNot(expression, i+1)
			i = k
		} else if IsBool(expression[i]) {
			value = parseBool(expression[i])
		} else if expression[i] == ')' {
			j = i
			break
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
		if expression[i] == ')' {
			j = i
			break
		}

		var v bool
		var k int
		var match bool
		if IsOr(expression[i]) {
			v, k = parseOr(expression, i+1)
			i = k
			match = true
		} else if IsAnd(expression[i]) {
			v, k = parseAnd(expression, i+1)
			i = k
			match = true
		} else if IsNot(expression[i]) {
			v, k = parseNot(expression, i+1)
			i = k
			match = true
		} else if IsBool(expression[i]) {
			v = parseBool(expression[i])
			match = true
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
		if expression[i] == ')' {
			j = i
			break
		}

		var v bool
		var k int
		var match bool
		if IsOr(expression[i]) {
			v, k = parseOr(expression, i+1)
			i = k
			match = true
		} else if IsAnd(expression[i]) {
			v, k = parseAnd(expression, i+1)
			i = k
			match = true
		} else if IsNot(expression[i]) {
			v, k = parseNot(expression, i+1)
			i = k
			match = true
		} else if IsBool(expression[i]) {
			v = parseBool(expression[i])
			match = true
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
