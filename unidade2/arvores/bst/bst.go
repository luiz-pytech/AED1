package tree

import "errors"

type BSTNode struct {
  left *BSTNode
  val int
  right *BSTNode
}

func CreateBSTNode(val int) *BSTNode{
  return &BSTNode{val:val}
}

//impl rec
func (node *BSTNode) Add(val int){
	
}

//impl rec
func (node *BSTNode) Search(val int) bool{
	return false
}

//impl rec
func (node *BSTNode) Min() int{
	return 0
}

//impl rec
func (node *BSTNode) Max() int{
	return 0
}

func (node *BSTNode) PreOrderNav(){

}

func (node *BSTNode) InOrderNav(){
	if node != nil {
		node.left.InOrderNav()
		print(node.val," ")
		node.right.InOrderNav()
	}  

}

func (node *BSTNode) PostOrderNav(){

}

func (node *BSTNode) LevelOrderNav(){

}

func (node *BSTNode) Remove(val int) error{
	return errors.New("not implemented") 
}