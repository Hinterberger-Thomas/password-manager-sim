package database

import (
	"encoding/json"
	"fmt"
)

func JSONToList(jsonOb string) []Account {

	var acc []Account

	json.Unmarshal([]byte(jsonOb), &acc)

	return acc
}
func ListToJSON(jsonOb []Account) []byte {
	appsJSON, err := json.Marshal(jsonOb)

	if err != nil {
		fmt.Println(err)
	}

	return appsJSON
}
