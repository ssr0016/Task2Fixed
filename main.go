package main

import "fmt"

func flattenHelper(iterable interface{}, result []interface{}) []interface{} {
	switch v := iterable.(type) {
	case []interface{}:
		for _, elem := range v {
			result = flattenHelper(elem, result)
		}
	default:
		if v != nil {
			result = append(result, v)
		}
	}
	return result
}

func Exercise(iterable []interface{}) []interface{} {
	result := make([]interface{}, 0)
	for _, elem := range iterable {
		result = flattenHelper(elem, result)
	}
	return result
}

func main() {
	nestedSlice := []interface{}{1, []interface{}{2, 3, nil, 4}, []interface{}{nil}, 5}
	flattenedSlice := Exercise(nestedSlice)
	fmt.Println(flattenedSlice)
}
