package tree

import (
	"fmt"
)

type BSTNode struct {
	left  *BSTNode
	val   int
	right *BSTNode
}

func CreateBSTNode(val int) *BSTNode {
	return &BSTNode{val: val}
}

func (node *BSTNode) Add(val int) {
	if val <= node.val { //esquerda
		if node.left != nil {
			node.left.Add(val)
		} else {
			node.left = CreateBSTNode(val)
		}
	} else { //direita
		if node.right != nil {
			node.right.Add(val)
		} else {
			node.right = CreateBSTNode(val)
		}
	}
}

func (node *BSTNode) Search(val int) bool {

	if node == nil {
		return false
	}
	if node.val == val {
		return true
	}

	if val <= node.val { //esquerda
		return node.left.Search(val)
	}

	return node.right.Search(val)

}

func (node *BSTNode) Min() int {
	if node.left == nil {
		return node.val
	} else {
		return node.left.Min()
	}
}

func (node *BSTNode) Max() int {
	if node.right == nil {
		return node.val
	} else {
		return node.right.Max()
	}
}

func (node *BSTNode) PreOrderNav() {

	if node == nil {
		return
	}

	fmt.Print(node.val, " ")
	node.left.PreOrderNav()
	node.right.PreOrderNav()

}

func (node *BSTNode) InOrderNav() {
	if node == nil {
		return
	}

	node.left.InOrderNav()
	fmt.Print(node.val, " ")
	node.right.InOrderNav()
}

func (node *BSTNode) PostOrderNav() {
	if node == nil {
		return
	}

	node.left.PostOrderNav()
	node.right.PostOrderNav()
	fmt.Print(node.val, " ")
}

func (node *BSTNode) LevelOrderNav() {
	if node == nil {
		return
	}

	queue := []*BSTNode{node}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		fmt.Print(current.val, " ")

		if current.left != nil {
			queue = append(queue, current.left)
		}
		if current.right != nil {
			queue = append(queue, current.right)
		}
	}
}

func (node *BSTNode) Height() int {
	if node == nil {
		return -1
	}

	left_height := node.left.Height()
	right_height := node.right.Height()

	if left_height > right_height {
		return left_height + 1
	}
	return right_height + 1
}

func (node *BSTNode) Remove(val int) *BSTNode {

	if val > node.val {
		node.right = node.right.Remove(val)
	} else if val < node.val {
		node.left = node.left.Remove(val)
	} else { // ENCONTREI
		if node.left == nil && node.right == nil {
			return nil
		} else if node.left != nil && node.right == nil {
			return node.left
		} else if node.left == nil && node.right != nil {
			return node.right
		}
		min := node.right.Min()
		node.val = min
		node.right.Remove(min)
	}

	return node

}

func (node *BSTNode) Size() int {
	if node == nil {
		return 0
	}
	return 1 + node.left.Size() + node.right.Size()
}
