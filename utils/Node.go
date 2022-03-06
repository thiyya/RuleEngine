package utils

type Node struct {
	Id             string  `json="value"`
	ParentId       string  `json="-"`
	CriterionInfo  string  `json="display"`
	ChildCriterion []*Node `json="items"`
}

func (n *Node) AddModelTreeInfo(nodes ...*Node) bool {
	var size = n.Size()
	for _, node := range nodes {
		if node.ParentId == n.Id {
			n.ChildCriterion = append(n.ChildCriterion, node)
		} else {
			for _, c := range n.ChildCriterion {
				if c.AddModelTreeInfo(node) {
					break
				}
			}
		}
	}
	return n.Size() == size+len(nodes)
}

func (n *Node) Size() int {
	size := len(n.ChildCriterion)
	for _, c := range n.ChildCriterion {
		size += c.Size()
	}
	return size
}
