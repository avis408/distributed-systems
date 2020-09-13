package dal

import (
	"testing"
)

func TestGetFileLines(t *testing.T) {

	file, err := GetFileLines("shard_1", "input_1.txt")
	if err != nil {
		t.Fail()
	}
	for _, f := range file {
		println(f)
	}
}
