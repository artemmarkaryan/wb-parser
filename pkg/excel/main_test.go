package excel

import (
	"path/filepath"
	"runtime"
	"testing"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basedir   = filepath.Dir(b)
)

func TestConvertAndSave(t *testing.T) {
	err := ConvertAndSave(
		[]map[string]string{
			{"name": "ivan"},
			{"weight": "100"},
		},
		filepath.Join(basedir, "test.xlsx"),
	)
	if err != nil {
		t.Error(err.Error())
	}
	return
}
