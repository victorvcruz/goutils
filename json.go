package goutils

import (
	"encoding/json"
	"log"
)

func ToJson(structure any) string {
	marshal, err := json.Marshal(structure)
	if err != nil {
		log.Println(err)
	}
	return string(marshal)
}
