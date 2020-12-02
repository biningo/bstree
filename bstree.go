// Copyright 2020 icepan. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package bstree

type node struct {
	item       interface{}
	leftChild  *node
	rightChild *node
}

type BSTree struct {
	root   *node
	length int                        //number of nodes
	comp   func(a, b interface{}) int //comparison function -1 0 1 < = >
}

func newNode(item interface{}) *node {
	n := &node{}
	n.item = item
	return n
}

//NewBSTree return a new BSTree
func NewBSTree(comp func(a, b interface{}) int) *BSTree {
	if comp == nil {
		panic("nil comp")
	}
	tree := &BSTree{}
	tree.comp = comp
	return tree
}

//Comp is a convenience function that performs a comparison of two items
//using the same "comp" function provided to New.
func (tree *BSTree) Comp(a, b interface{}) int {
	return tree.comp(a, b)
}

// Len returns the number of items in the tree
func (tree *BSTree) Len() int {
	return tree.length
}

func (tree *BSTree) find(item interface{}) (*node, bool) {
	cur := tree.root
	for cur != nil {
		t := tree.comp(item, cur.item)
		if t < 0 {
			cur = cur.leftChild
		} else if t > 0 {
			cur = cur.rightChild
		} else {
			return cur, true
		}
	}
	return nil, false
}
func (tree *BSTree) findNodeAndParent(item interface{}) (pre *node, cur *node) {
	cur = tree.root
	t := tree.comp(item, cur.item)
	for cur != nil && t != 0 {
		pre = cur
		if t < 0 {
			cur = cur.leftChild
		} else {
			cur = cur.rightChild
		}
		t = tree.comp(item, cur.item)
	}
	return pre, cur
}

//Set or replace a value for a key
func (tree *BSTree) Set(item interface{}) {
	if tree.root == nil {
		tree.root = newNode(item)
		return
	}
	cur := tree.root
	comp := tree.comp
	for cur != nil {
		t := comp(item, cur.item)
		if t < 0 {
			if cur.leftChild == nil {
				cur.leftChild = newNode(item)
			}
			cur = cur.leftChild
		} else if t > 0 {
			if cur.rightChild == nil {
				cur.rightChild = newNode(item)
			}
			cur = cur.rightChild
		} else {
			cur.item = item
			return
		}
	}
}

//Get value for key
func (tree *BSTree) Get(item interface{}) (interface{}, bool) {
	node, exist := tree.find(item)
	if !exist {
		return nil, false
	}
	return node.item, true
}

//Exist return true if item exist
func (tree *BSTree) Exist(item interface{}) bool {
	_, flag := tree.find(item)
	return flag
}

//Del node by key
func (tree *BSTree) Del(item interface{}) bool {
	pre, cur := tree.findNodeAndParent(item)
	if cur == nil { //not found
		return false
	}

	var right_min *node
	var right_min_pre *node

	//Converts to only right subtrees or no child nodes
	if cur.leftChild != nil && cur.rightChild != nil {
		right_min_pre = cur
		right_min = cur.rightChild
		for right_min.leftChild != nil {
			right_min_pre = right_min
			right_min = right_min.leftChild
		}
		pre = right_min_pre
		cur = right_min
	}
	var child *node
	if cur.leftChild != nil {
		child = cur.leftChild
	} else if cur.rightChild != nil {
		child = cur.rightChild
	}
	if pre == nil {
		tree.root = child
	} else if pre.leftChild == cur {
		pre.leftChild = child
	} else if pre.rightChild == cur {
		pre.rightChild = child
	}
	return true
}

//Max find max item
func (tree *BSTree) Max() interface{} {
	node := tree.root
	for node != nil {
		if node.rightChild == nil {
			return node.item
		}
		node = node.rightChild
	}
	return nil
}

//Min find min item
func (tree *BSTree) Min() interface{} {
	node := tree.root
	for node != nil {
		if node.leftChild == nil {
			return node.item
		}
		node = node.leftChild
	}
	return nil
}

// Scan the tree by order
// Return false to stop
func (tree *BSTree) Scan(handler func(item interface{}) bool) {
	if handler == nil {
		return
	}
	stack := []*node{}
	node := tree.root
	for node != nil {
		stack = append(stack, node)
		node = node.leftChild
	}

	for len(stack) > 0 {
		node, stack = stack[len(stack)-1], stack[:len(stack)-1]
		if flag := handler(node.item); flag == false {
			return
		}
		node = node.rightChild
		for node != nil {
			stack = append(stack, node)
			node = node.leftChild
		}
	}
}

//Range scan the tree within the range [start,end]
func (tree *BSTree) Range(start, end interface{}, handler func(item interface{}) bool) {
	if handler == nil {
		return
	}
	stack := []*node{}
	node := tree.root
	for node != nil {
		stack = append(stack, node)
		node = node.leftChild
	}

	for len(stack) > 0 {
		node, stack = stack[len(stack)-1], stack[:len(stack)-1]
		if tree.comp(start, node.item) <= 0 && tree.comp(end, node.item) >= 0 {
			if flag := handler(node.item); flag == false {
				return
			}
		}
		node = node.rightChild
		for node != nil {
			stack = append(stack, node)
			node = node.leftChild
		}
	}
}
