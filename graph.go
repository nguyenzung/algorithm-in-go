package main

import (
	"fmt"
	"math"
	// "github.com/Workiva/go-datastructures"
)

type Node struct {
	label       interface{}
	linkedNodes map[*Node]float64
}

func (node *Node) addLinkedNode(linkedNode *Node, weight float64) bool {
	if linkedNode != nil {
		node.linkedNodes[linkedNode] = weight
		return true
	}
	return false
}

func (node *Node) getLinkedNodes() []*Node {
	nodes := make([]*Node, 0, len(node.linkedNodes))
	for k := range node.linkedNodes {
		nodes = append(nodes, k)
	}
	return nodes
}

func (node *Node) getDistance(d *Node) float64 {
	return node.linkedNodes[d]
}

type Graph struct {
	nodes map[interface{}]*Node
}

func (graph *Graph) addNode(node *Node) bool {
	if graph.nodes[node.label] == nil {
		graph.nodes[node.label] = node
		return true
	}
	return false
}

func (graph *Graph) containNode(label interface{}) bool {
	return graph.nodes[label] != nil
}

func (graph *Graph) getNode(label interface{}) *Node {
	return graph.nodes[label]
}

func (graph *Graph) makeConnection(s interface{}, d interface{}, w float64) bool {
	sNode := graph.getNode(s)
	dNode := graph.getNode(d)
	if sNode != nil && dNode != nil {
		return sNode.addLinkedNode(dNode, w)
	}
	return false
}

func (graph *Graph) checkPathDFS(s interface{}, d interface{}) bool {
	sNode := graph.getNode(s)
	dNode := graph.getNode(d)
	path := ""
	result := false
	if sNode != nil && dNode != nil {
		visitedNodes := make(map[*Node]bool)
		result = graph.dfs(sNode, dNode, visitedNodes, &path)
	}
	fmt.Println("[DFS PATH]: {", path, "}")
	return result
}

func (graph *Graph) checkPathBFS(s interface{}, d interface{}) bool {
	sNode := graph.getNode(s)
	dNode := graph.getNode(d)
	pathMap := make(map[*Node]*Node)
	result := false
	if sNode != nil && dNode != nil {
		visitedNodes := make(map[*Node]bool)
		result = graph.bfs(sNode, dNode, visitedNodes, pathMap)
	}
	fmt.Println("[BFS PATH]: {", printBfsPath(sNode, dNode, pathMap), "}")
	return result
}

func (graph *Graph) shortestPath(s interface{}, d interface{}) bool {
	var num uint64
	num = math.MaxUint64
	fmt.Println(num)
	return false
}

func (graph *Graph) dfs(s *Node, d *Node, visitedNodes map[*Node]bool, path *string) bool {
	if s == d {
		return true
	}
	visitedNodes[s] = true
	linkedNodes := s.getLinkedNodes()
	for _, node := range linkedNodes {
		if !visitedNodes[node] && graph.dfs(node, d, visitedNodes, path) {
			*path = fmt.Sprintf("[%d -----> %d] ", s.label, node.label) + *path
			return true
		}
	}
	return false
}

func (graph *Graph) bfs(s *Node, d *Node, visitedNodes map[*Node]bool, pathMap map[*Node]*Node) bool {
	queue := make([]*Node, 0)
	queue = append(queue, s)
	visitedNodes[s] = true
	var node *Node
	for len(queue) > 0 {
		node, queue = popNodeOutOfQueue(queue)
		if node == d {
			return true
		} else {
			linkedNodes := node.getLinkedNodes()
			for _, adjNode := range linkedNodes {
				if !visitedNodes[adjNode] {
					visitedNodes[adjNode] = true
					pathMap[adjNode] = node
					queue = append(queue, adjNode)
				}
			}
		}
	}
	return false
}

func printBfsPath(s *Node, d *Node, pathMap map[*Node]*Node) string {
	node := d
	path := ""
	for node != s {
		prev := pathMap[node]
		if prev == nil {
			break
		}
		path = (fmt.Sprintf("[%d -----> %d] ", prev.label, node.label) + path)
		node = prev
	}
	return path
}

func popNodeOutOfQueue(queue []*Node) (*Node, []*Node) {
	if len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		return node, queue
	}
	return nil, queue
}

func (graph *Graph) describe() {
	fmt.Println("\nGraph description", graph.nodes)
}

func makeNode(label interface{}) *Node {
	node := Node{label: label, linkedNodes: make(map[*Node]float64)}
	return &node
}

func makeGraph() *Graph {
	graph := Graph{nodes: make(map[interface{}]*Node)}
	return &graph
}

func main() {
	graph := makeGraph()
	fmt.Println("[AddNode]", graph.addNode(makeNode(0)))
	fmt.Println("[AddNode]", graph.addNode(makeNode(1)))
	fmt.Println("[AddNode]", graph.addNode(makeNode(2)))
	fmt.Println("[AddNode]", graph.addNode(makeNode(3)))
	fmt.Println("[AddNode]", graph.addNode(makeNode(4)))
	fmt.Println("[AddNode]", graph.addNode(makeNode(5)))
	fmt.Println("[AddNode]", graph.addNode(makeNode(6)))
	fmt.Println("[AddNode]", graph.addNode(makeNode(7)))
	fmt.Println("[AddNode]", graph.addNode(makeNode(8)))

	fmt.Println("[MakeConnection]", graph.makeConnection(0, 1, 12))
	fmt.Println("[MakeConnection]", graph.makeConnection(1, 5, 12))
	fmt.Println("[MakeConnection]", graph.makeConnection(6, 2, 12))
	fmt.Println("[MakeConnection]", graph.makeConnection(6, 5, 12))
	fmt.Println("[MakeConnection]", graph.makeConnection(0, 4, 12))
	fmt.Println("[MakeConnection]", graph.makeConnection(4, 6, 12))
	fmt.Println("[MakeConnection]", graph.makeConnection(2, 5, 12))
	fmt.Println("[MakeConnection]", graph.makeConnection(5, 7, 12))
	fmt.Println("[MakeConnection]", graph.makeConnection(7, 8, 12))

	fmt.Println("[GetNode]", graph.getNode(0).getLinkedNodes())

	fmt.Println("[DFS]", graph.checkPathDFS(6, 8))
	fmt.Println("[BFS]", graph.checkPathBFS(6, 8))

	// graph.shortestPath(0, 7)
	graph.describe()
}
