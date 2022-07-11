package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var input string
	fmt.Scan(&input)

	var j map[string]interface{}
	err := json.Unmarshal([]byte(input), &j)
	if err != nil {
		panic(err)
	}

	fmt.Println(addMap(j))
}

func addMap(m map[string]interface{}) int {
	total := 0
	for _, v := range m {
		switch vTyped := v.(type) {
		case float64:
			total += int(vTyped)
		case []interface{}:
			total += addList(vTyped)
		case map[string]interface{}:
			total += addMap(vTyped)
		}
	}
	return total
}

func addList(l []interface{}) int {
	total := 0
	for _, v := range l {
		switch vTyped := v.(type) {
		case float64:
			total += int(vTyped)
		case []interface{}:
			total += addList(vTyped)
		case map[string]interface{}:
			total += addMap(vTyped)
		}
	}
	return total
}
