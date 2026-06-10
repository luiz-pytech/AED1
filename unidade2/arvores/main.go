package main

import . "arvores/bst"

func main() {
	bstRoot := CreateBSTNode(8)
	bstRoot.Add(1)
	bstRoot.Add(10)
	bstRoot.Add(23)
	bstRoot.Add(2)
	bstRoot.Add(6)
	bstRoot.Add(9)
	bstRoot.Add(4)
	bstRoot.Add(3)

	print("\n", bstRoot.Search(-10), "\n")
	print(bstRoot.Search(100), "\n")
	print(bstRoot.Search(9), "\n")

	print(bstRoot.Min(), "\n")
	print(bstRoot.Max(), "\n")

	// print("Pré-ordem: ")
	// bstRoot.PreOrderNav()
	// print("\nEm-ordem: ")
	// bstRoot.InOrderNav()
	//print("\nPós-ordem: ")
	//bstRoot.PostOrderNav()
	print("\nPor níveis: ")
	bstRoot.LevelOrderNav()

	print("\nAltura: ", bstRoot.Height())
	print("\nTamanho: ", bstRoot.Size())
	//print("\nNum pares: ", bstRoot.NumPar())

}
