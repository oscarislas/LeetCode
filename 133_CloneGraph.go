//https://leetcode.com/problems/clone-graph/
//133. Clone Graph

// Given a reference of a node in a connected undirected graph.
// Return a deep copy (clone) of the graph.
// Each node in the graph contains a val (int) and a list (List[Node]) of its neighbors.
// class Node {
//     public int val;
//     public List<Node> neighbors;
// }

package main

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	visited := make(map[int]*Node, 0)
	queue := []*Node{node}
	visited[node.Val] = &Node{Val: node.Val}
	for len(queue) > 0 {
		oldNode := queue[0]
		queue = queue[1:]
		for _, neighbor := range oldNode.Neighbors {
			if _, ok := visited[neighbor.Val]; !ok {
				visited[neighbor.Val] = &Node{Val: neighbor.Val}
				queue = append(queue, neighbor)
			}
			visited[oldNode.Val].Neighbors = append(visited[oldNode.Val].Neighbors, visited[neighbor.Val])
		}
	}
	return visited[node.Val]
}
