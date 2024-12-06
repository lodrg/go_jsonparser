package test

import (
	"github.com/stretchr/testify/assert"
	"go_jsonparser/api"
	"testing"
)

func TestJsonParser(t *testing.T) {
	json := "{\"this is string\":\"good json\"}"
	parseJson := api.ParseJson(json)
	if m, ok := parseJson.(map[string]interface{}); ok {
		t.Logf("map值: %+v", m)
	}
	expect := map[string]interface{}{"this is string": "good json"}
	assert.Equal(t, expect, parseJson, "map not match")
}

func TestJsonParser_with_number(t *testing.T) {
	json := "{\"age\":22}"
	parseJson := api.ParseJson(json)
	if m, ok := parseJson.(map[string]interface{}); ok {
		t.Logf("map值: %+v", m)
	}
	expect := map[string]interface{}{"age": 22}
	assert.Equal(t, expect, parseJson, "map not match")
}

func TestJsonParser_with_numberAndString(t *testing.T) {
	json := "{\"this is string\":\"good json\",\"age\":22}"
	parseJson := api.ParseJson(json)
	if m, ok := parseJson.(map[string]interface{}); ok {
		t.Logf("map值: %+v", m)
	}
	expect := map[string]interface{}{"this is string": "good json", "age": 22}
	assert.Equal(t, expect, parseJson, "map not match")
}

func TestJsonParser_with_neg_number(t *testing.T) {
	json := "{\"age\":-22}"
	parseJson := api.ParseJson(json)
	if m, ok := parseJson.(map[string]interface{}); ok {
		t.Logf("map值: %+v", m)
	}
	expect := map[string]interface{}{"age": -22}
	assert.Equal(t, expect, parseJson, "map not match")
}

func TestJsonParser_with_array(t *testing.T) {
	json := "[{\"id\":1,\"value\":10},{\"id\":2,\"value\":20}]"
	result := api.ParseJson(json)

	// 1. 检查类型是否为数组
	arr, ok := result.([]interface{})
	assert.True(t, ok, "should be array type")

	// 2. 检查数组长度
	assert.Equal(t, 2, len(arr), "array length should be 2")

	// 3. 检查第一个对象
	obj1, ok := arr[0].(map[string]interface{})
	assert.True(t, ok, "first element should be object")
	assert.Equal(t, 1, obj1["id"])
	assert.Equal(t, 10, obj1["value"])

	// 4. 检查第二个对象
	obj2, ok := arr[1].(map[string]interface{})
	assert.True(t, ok, "second element should be object")
	assert.Equal(t, 2, obj2["id"])
	assert.Equal(t, 20, obj2["value"])
}
