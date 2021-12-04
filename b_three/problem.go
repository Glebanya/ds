package b_tree

func StringToBTree(str string) *BinaryTree {
	bt := &(BinaryTree{})
	runeArray := []rune(str)
	addToTree(bt.root, runeArray)
	return bt
}

func addToTree(node *binaryNode, runeArray []rune) {

	runeArrayLen := len(runeArray)
	if runeArrayLen == 0 {
		return
	}
	if runeArrayLen == 1 {
		if node == nil {
			node = &binaryNode{
				data: runeArray[0],
			}
		}
	}
	var (
		brackets    = 0
		symbolIndex = 1
	)
	for i := 0; i < runeArrayLen; i++ {
		if runeArray[i] == '(' {
			brackets++
		} else if runeArray[i] == ')' {
			brackets--
		} else {
			symbolIndex = i
		}
		addToTree(node, runeArray[symbolIndex:symbolIndex+1])
		addToTree(node.left, runeArray[2:symbolIndex-3])
		addToTree(node.right, runeArray[symbolIndex+2:runeArrayLen-3])
	}

}
