package main

import (
	"encoding/json"
	"fmt"
)

type rule struct {
	RuleName string `json:"rule_name"`
	RuleType int32  `json:"rule_type"`
}

func main() {
	jsonStr := `{
		"rule_name": "55",
		"rule_type": 4
	}`
	var u rule
	json.Unmarshal([]byte(jsonStr), &u)
	fmt.Println(u)
}
