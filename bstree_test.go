package bstree

import (
	"os"
	"testing"
)

/**
*@Author icepan
*@Date 2020/12/5 下午7:07
*@Describe
**/
type Item struct {
	Key int
	Val string
}

var bts []*BSTree

func byKey(a, b interface{}) int {
	k1, ok := a.(int)
	k2 := b.(Item)
	if !ok {
		k3 := a.(Item)
		return k3.Key - k2.Key
	}
	return k1 - k2.Key
}

func createItems1() []Item {
	items := []Item{}
	i1 := Item{5, "five"}
	i2 := Item{3, "three"}
	i3 := Item{7, "seven"}
	i4 := Item{4, "four"}
	i5 := Item{6, "six"}
	i6 := Item{10, "ten"}
	items = append(items, i1, i2, i3, i4, i5, i6)
	return items
}

func buildTree1(items []Item) *BSTree {
	bt := NewBSTree(byKey)
	for _, item := range createItems1() {
		bt.Set(item)
	}
	return bt
}

func setup() {
	bt1 := buildTree1(createItems1())
	bts = append(bts, bt1)
}

func TestBSTree_Min(t *testing.T) {
	for _, bt := range bts {
		min := bt.Min().(Item)
		if min.Key != 3 {
			t.Errorf("expected %d,got %d", 3, min.Key)
		}
	}
}

func TestBSTree_Max(t *testing.T) {
	for _, bt := range bts {
		max := bt.Max().(Item)
		if max.Key != 10 {
			t.Errorf("expected %d,got %d", 3, max.Key)
		}
	}
}

func TestBSTree_Len(t *testing.T) {
	items := createItems1()
	for _, bt := range bts {
		if bt.Len() != len(items) {
			t.Errorf("expected %d,got %d", len(items), bt.Len())
		}
	}
}

func TestBSTree_Exist(t *testing.T) {
	items := createItems1()
	for _, bt := range bts {
		for _, item := range items {
			if f := bt.Exist(item); !f {
				t.Errorf("expected %v,got %v", true, f)
			}
		}
	}

}

func TestBSTree_Get(t *testing.T) {
	items := createItems1()
	for _, bt := range bts {
		for _, item := range items {
			i, f := bt.Get(item)
			if f == false {
				t.Errorf("expected %d,got null", item.Key)
				continue
			}
			it := i.(Item)
			if it.Key != item.Key {
				t.Errorf("expected %s,got %s\n", item.Val, it.Val)
			}
		}
	}
}

func TestBSTree_Del(t *testing.T) {
	items := createItems1()
	for _, bt := range bts {
		for _, item := range items {
			if f := bt.Del(item); f == false {
				t.Errorf("%d,delete fail!", item.Key)
			}
			if _, f := bt.Get(item); f == true {
				t.Errorf("deleted %d,but still exists", item.Key)
				return
			}
		}
	}
}

func TestBSTree_Comp(t *testing.T) {
	for _, bt := range bts {
		if f := bt.Comp(Item{1, "one"}, Item{2, "two"}); f >= 0 {
			t.Errorf("expected %d,got %d", -1, f)
		}
	}
}

func TestBSTree_Scan(t *testing.T) {
	items := createItems1()
	ans := make(map[string]int)
	for _, item := range items {
		ans[item.Val] = item.Key
	}
	for _, bt := range bts {
		bt.Scan(func(item interface{}) bool {
			it := item.(Item)
			if v, ok := ans[it.Val]; !ok || v != it.Key {
				t.Errorf("expected %d,got %d", v, it.Key)
				return false
			}
			return true
		})
	}
}

func TestBSTree_Range(t *testing.T) {
	for _, bt := range bts {
		for start := 0; start < 10; start += 2 {
			for end := start + 1; end < 20; end += 2 {
				bt.Range(start, end, func(item interface{}) bool {
					it := item.(Item)
					if it.Key > end {
						t.Errorf("expected %d<val<%d,got %d>%d", start, end, it.Key, end)
					} else if it.Key < start {
						t.Errorf("expected %d<val<%d,got %d<%d", start, end, it.Key, start)
					}
					return true
				})
			}
		}
	}
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

//==================Benchmark================
func BenchmarkBSTree_Set(b *testing.B) {
	bt := NewBSTree(byKey)
	for i := 0; i < b.N; i++ {
		bt.Set(Item{i, ""})
	}
}

func BenchmarkBSTree_Get(b *testing.B) {
	bt := NewBSTree(byKey)
	for i := 0; i < b.N; i++ {
		bt.Set(Item{i, ""})
	}
	for i := 0; i < b.N; i++ {
		bt.Get(i)
	}
}

func BenchmarkBSTree_Scan(b *testing.B) {
	bt := NewBSTree(byKey)
	for i := 0; i < b.N; i++ {
		bt.Set(Item{i, ""})
	}
	bt.Scan(func(item interface{}) bool {
		return true
	})
}

func BenchmarkBSTree_Range(b *testing.B) {
	bt := NewBSTree(byKey)
	for i := 0; i < b.N; i++ {
		bt.Set(Item{i, ""})
	}
	n := b.N
	for i := 0; i < n; i++ {
		bt.Range(i, n, func(item interface{}) bool {
			return true
		})
	}
}
