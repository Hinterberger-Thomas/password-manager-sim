package database

import (
	"encoding/json"
	"fmt"
)

func jsonToList(jsonOb string) []Account {

	var acc []Account

	json.Unmarshal([]byte(jsonOb), &acc)

	return acc
}
func ListToJson(jsonOb []Account) []byte {
	appsJSON, err := json.Marshal(jsonOb)

	if err != nil {
		fmt.Println(err)
	}

	return appsJSON
}
