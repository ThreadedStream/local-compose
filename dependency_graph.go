package main

type vertex struct {
	name string
	deps []vertex
}

var (
	vertices []vertex
)

//Adds dependency to the vertex
func addDependency(from vertex, to vertex) {
	from.deps = append(from.deps, to)
	if !find(vertices, from.name) {
		vertices = append(vertices, from)
	}
	if !find(vertices, to.name) {
		vertices = append(vertices, to)
	}
}

//Deletes specified dependency
func deleteDependency(from vertex, dependency string) {
	for idx, vtx := range from.deps {
		if vtx.name == dependency {
			from.deps = removeElement(from.deps, idx)
		}
	}
}

//finds the first vertex with dependencies present
func findVertexWithDeps() vertex {
	for _, vtx := range vertices {
		if vtx.deps != nil {
			return vtx
		}
	}
	return vertex{}
}

//determine the order of execution
func determineOrder(node vertex, resolved *[]vertex) {
	for _, dep := range node.deps {
		if !find(*resolved, dep.name) {
			determineOrder(dep, resolved)
		}
	}
	*resolved = append(*resolved, node)
}

func removeElement(from []vertex, idx int) []vertex {
	from[idx] = from[len(from)-1]

	return from[:len(from)-1]
}
