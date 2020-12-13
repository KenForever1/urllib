package parse

import (
	"encoding/json"
	"fmt"
	"strings"
	"urllib/utils"
)

func Urlencode(query string) string {

	data := make(map[string]interface{})
	err := json.Unmarshal([]byte(query), &data)
	if err != nil {
		fmt.Println("Umarshal failed:", err)
	}
	var build strings.Builder
	for k, v := range data {
		build.WriteString(k)
		build.WriteString("=")
		build.WriteString(utils.Strval(v))
		build.WriteString("&")
	}
	return build.String()[:build.Len()-1]
}
