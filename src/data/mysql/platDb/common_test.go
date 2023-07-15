package platDb

import (
	"encoding/json"
	"testing"
)

func TestCommAll(t *testing.T) {
	InitPlatFromDb()

	one, err := DictGroupTable.FindAll()
	// one, err := DictGroupTable.FindById(1)
	// one, err := DictGroupTable.FindByIds([]uint64{1, 2, 3})
	// one, err := DictGroupTable.FindByObject(&DictGroup{ID: 1})
	oneStr := []byte("")
	if err == nil {
		oneStr, _ = json.Marshal(one)
	}
	t.Logf("%s", string(oneStr))
	t.Logf("%v", err)
}
