package api

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	index int
	json  string
)

func ParseJson(jsonString string) interface{} {
	json = strings.TrimSpace(jsonString)
	index = 0
	return parseValue()
}

func parseValue() interface{} {
	if index >= len(json) {
		return nil
	}

	c := json[index]

	switch c {
	case '{':
		return parseObject()
	case '[':
		return parseArray()
	case '"':
		return parseString()
	default:
		if isDigit(json[index]) || json[index] == '-' {
			return parseNumber()
		}
		panic(fmt.Sprintf("意外的字符: %c", c))
	}
}

func parseNumber() interface{} {
	start := index

	// 处理负号
	if index < len(json) && json[index] == '-' {
		index++
	}

	// 确保负号后面有数字
	if !isDigit(json[index]) {
		panic("invalid number format: digit expected after '-'")
	}

	// 找到数字的结尾
	for index < len(json) && isDigit(json[index]) {
		index++
	}
	number := json[start:index]
	result, _ := strconv.Atoi(number) // 忽略错误处理以保持简洁
	return result
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func parseString() string {
	sb := strings.Builder{}
	index++
	for json[index] != '"' {
		sb.WriteByte(json[index])
		index++
	}
	index++
	return sb.String()
}

func parseArray() interface{} {
	list := make([]interface{}, 0)
	index++

	if json[index] == ']' {
		index++
		return list
	}

	// 找到结束的 ]
	for json[index] != ']' {
		// []之间的值是数组元素
		// 递归调用 parseValue 解析值
		value := parseValue()
		list = append(list, value)

		if json[index] == ',' {
			index++ // 跳过 ','
		}
	}
	index++ // 跳过 ']'
	return list
}

func parseObject() map[string]interface{} {
	m := make(map[string]interface{})
	index++

	// 处理空对象
	if index < len(json) && json[index] == '}' {
		index++
		return m
	}

	// 循环解析键值对，直到遇到结束符号
	for index < len(json) {
		// 解析一个键值对
		key := parseString()
		if index < len(json) && json[index] == ':' {
			index++
			value := parseValue()
			m[key] = value
		}

		// 检查结束符号或逗号
		if index >= len(json) {
			break
		}

		if json[index] == '}' {
			index++
			return m
		}

		if json[index] == ',' {
			index++
		}
	}
	panic("Invalid JSON object")
}
