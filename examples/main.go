package main

import "log"

/**
*@Author icepan
*@Date 2020/12/2 下午4:04
*@Describe
**/
import "github.com/biningo/bstree"

type Item struct {
	Key int
	Val string
}

func byKey(a, b interface{}) int {
	ia, ib := a.(Item), b.(Item)
	return ia.Key - ib.Key
}

func main() {
	tree := bstree.NewBSTree(byKey)
	tree.Set(Item{5, "five"})
	tree.Set(Item{3, "three"})
	tree.Set(Item{7, "seven"})
	tree.Set(Item{4, "four"})
	tree.Set(Item{6, "six"})
	tree.Set(Item{10, "ten"})

	arrKey := []int{}
	arrVal := []string{}
	tree.Scan(func(item interface{}) bool {
		i := item.(Item)
		arrKey = append(arrKey, i.Key)
		arrVal = append(arrVal, i.Val)
		return true
	})

	log.Println(arrKey)
	log.Println(arrVal)

	f := tree.Del(Item{Key: 10})
	log.Println(f)
	if v, f := tree.Get(Item{Key: 3}); f == true {
		item := v.(Item)
		log.Println(item.Val)
	}
	tree.Set(Item{31, "三十一"})
	tree.Scan(func(item interface{}) bool {
		i := item.(Item)
		log.Println(i.Key, i.Val)
		return true
	})

	item := tree.Max().(Item)
	log.Println(item.Val)
	item = tree.Min().(Item)
	log.Println(item.Val)

	arrKey = []int{}

	tree.Range(Item{Key: 4}, Item{Key: 6}, func(item interface{}) bool {
		i := item.(Item)
		arrKey = append(arrKey, i.Key)
		return true
	})
	log.Println(arrKey)
}
