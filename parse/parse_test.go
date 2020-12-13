package parse

import (
	"fmt"
	"testing"
)

func TestUrlencode(t *testing.T) {
	queryStr := `{"ratifyNo":"","projectName":"","personInCharge":"","dependUnit":"","code":"B01","projectType":"218","subPType":"","psPType":"","keywords":"","ratifyYear":"","conclusionYear":"2010","beginYear":"","endYear":"","checkDep":"","checkType":"","quickQueryInput":"","adminID":"","pageNum":0,"pageSize":5,"queryType":"input","complete":"true"}`
	urlParams := Urlencode(queryStr)
	fmt.Println(urlParams)
}
