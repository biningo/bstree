# btree

[![GoDoc](https://godoc.org/github.com/biningo/bstree?status.svg)](https://godoc.org/github.com/biningo/bstree)

An efficient and simple [Binary Search Tree](https://en.wikipedia.org/wiki/Binary_search_tree) implementation in Go. 

## Installing

To start using btree, install Go and run `go get`:

```sh
$ go get -u github.com/biningo/bstree
```

## Usage

```go
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

	log.Println(arrKey) // [3,4,5,6,7,10]
	log.Println(arrVal)

	f := tree.Del(Item{Key: 10})
	log.Println(f) //true
    
	if v, f := tree.Get(Item{Key: 3}); f == true {
		item := v.(Item)
		log.Println(item.Val) //"three"
	}
    
	tree.Scan(func(item interface{}) bool {
		i := item.(Item)
		log.Println(i.Key, i.Val)
		return true
	})
    // 1 one
    // 2 two
    // ....
    

	item := tree.Max().(Item)
	log.Println(item.Val) // "seven"
	item = tree.Min().(Item)
	log.Println(item.Val) // "three"

    
	arrKey = []int{}
	tree.Range(Item{Key: 4}, Item{Key: 6}, func(item interface{}) bool {
		i := item.(Item)
		arrKey = append(arrKey, i.Key)
		return true
	})
	log.Println(arrKey) // [4,5,6]



	//by key
	tree2:=bstree.NewBSTree(func(a, b interface{}) int {
		key,ok:=a.(int)
		item:=b.(Item)
		if !ok{
			key2 := a.(Item)
			return key2.Key-item.Key
		}
		return key-item.Key
	})


	tree2.Set(Item{5, "five"})
	tree2.Set(Item{3, "three"})
	tree2.Set(Item{7, "seven"})
	tree2.Set(Item{4, "four"})
	tree2.Set(Item{6, "six"})

	tree2.Range(4,6, func(item interface{}) bool {
		i:=item.(Item)
		log.Println(i.Key,i.Val)
		return true
	})
    
    tree2.Get(6)
    
	m:=tree2.Max().(Item)
	log.Println(m.Key)
}

```

## Operations

### Basic

```
Len()                   # return the number of items in the bstree
Set(item)               # insert or replace an existing item
Get(item)               # get an existing item
Del(item)            # delete an item
```

### Iteration

```
Scan(iter)     #Scan the tree by order
Range(start,end,item) #scan the tree within the range [start,end]
```

### Queues

```
Min()                   # return the first item in the bstree
Max()                   # return the last item in the bstree
TODO:PopMin()                # remove and return the first item in the TODO:PopMax()                # remove and return the last item in the bstree
```
## Benchmarks

```b
TODO
```

## License

Source code is available under the MIT [License](/LICENSE).
