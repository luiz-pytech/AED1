package tree

type ITree interface {
  Add(value int)
  Search(value int) bool
  Min() int
  Max() int
  PreOrderNav()
  InOrderNav()
  PostOrderNav()
  LevelOrderNav()
  Remove(value int) error
  Size() int
}