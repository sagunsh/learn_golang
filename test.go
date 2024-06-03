package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	var jsonString = `{"name": "Jane Doe", "address": {"city": "New York"}}`
	var data map[string]interface{}

	err := json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		// Handle error
	}

	name := data["name"].(string)
	address := data["address"].(map[string]interface{})
	city := address["city"].(string)

	fmt.Println(name, city) // Output: Jane Doe New York
}
