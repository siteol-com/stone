package platDb

import (
	"testing"
)

func Test_FindOne(t *testing.T) {
	InitPlatFromDb()
	//query := &Account{ID: 1}
	//res, err := query.FindOne()
	//t.Logf("Res => %v", res)
	//t.Logf("Err => %v", err)

	queryDictGroup := &DictGroup{}
	res1, err := queryDictGroup.FindAll()
	t.Logf("Res1 => %v", res1)
	t.Logf("Err => %v", err)
}
