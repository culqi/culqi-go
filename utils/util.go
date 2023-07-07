package culqi

import (
	"encoding/json"
	"fmt"
)

func JsonToMap(data []byte) map[string]interface{} {
	var mapData map[string]interface{}
	errorJson := json.Unmarshal([]byte(data), &mapData)
	if errorJson != nil {
		fmt.Println("Error while decoding the data", errorJson.Error())
	}
	return mapData
}
