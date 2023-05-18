package test_demo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	got := Split(map[string]interface{}{"a": 1, "b": 2, "c": 3}, "|")
	want := []interface{}{"a", "|", "b", "|", "c", "|"}
	fmt.Println(got)
	fmt.Println(want)
	if !reflect.DeepEqual(got, want) {
		t.Error("excepted:", want, "got:", got)
	}

}
