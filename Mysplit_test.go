package Mysplit

import (
	"reflect"
	"testing"
)

func TestMySplit(t *testing.T) {
	t.Log("功能测试")
	got := Mysplit("a:b:c:d", ":")
	want := []string{"a", "b", "c", "d"}
	if ok := reflect.DeepEqual(got, want); !ok {
		t.Fatalf("功能测试失败, 期望:%v  运行结果:%v", want, got)
	}
}

func TestNoneSplit(t *testing.T) {
	got := Mysplit("a:b:c:d", "|")
	want := []string{"a:b:c:d"}
	if ok := reflect.DeepEqual(got, want); !ok {
		t.Fatalf("功能测试失败, 期望:%v  运行结果:%v", want, got)
	}
}

func TestNoneSplit1(t *testing.T) {
	got := Mysplit("a:b:c::d", ":")
	want := []string{"a", "b", "c", "d"}
	if ok := reflect.DeepEqual(got, want); !ok {
		t.Fatalf("功能测试失败, 期望:%v  运行结果:%v", want, got)
	}
}
