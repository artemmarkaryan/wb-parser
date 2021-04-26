package mapToCSV

import (
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"testing"
	"time"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basedir   = filepath.Dir(b)
)

func TestConvertAndSave(t *testing.T) {
	m := []map[string]string{
		{"name": "b", "size": "m"},
		{"name": "b", "package": "box"},
	}
	filename := basedir + "/" + strconv.Itoa(int(time.Now().Unix())) + ".csv"

	err := ConvertAndSave(m, filename)
	if err != nil {
		t.Error(err.Error())
		return
	}
	err = os.Remove(filename)
	if err != nil {
		t.Error(err.Error())
	}
}
