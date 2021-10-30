package graph

import "fmt"

type Graph struct {
	vertices []*Vertex
}

type Vertex struct {
	key      int
	adjacent []*Vertex
}

func (g *Graph) AddVertex(k int) {
	if !contains(g.vertices, k) {
		g.vertices = append(g.vertices, &Vertex{key: k})
	} else {
		fmt.Printf("key %d already exists", k)
	}
}

func (g *Graph) AddEdge(from, to int) {
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	//check error
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("Invalid edge (%v-->%v", fromVertex, toVertex)
		fmt.Println(err.Error())
	} else if contains(fromVertex.adjacent, to) {
		err := fmt.Errorf("Existing edge (%v-->%v", fromVertex, toVertex)
		fmt.Println(err.Error())
	} else {
		//add edge
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	}
}
func (g *Graph) getVertex(k int) *Vertex {
	for i, v := range g.vertices {
		if v.key == k {
			return g.vertices[i]
		}
	}
	return nil
}
func contains(s []*Vertex, k int) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}
	}
	return false
}
func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\n Vertex %v : ", v.key)
		for _, v := range v.adjacent {
			fmt.Printf("%v", v.key)
		}
	}
	fmt.Println()
}
