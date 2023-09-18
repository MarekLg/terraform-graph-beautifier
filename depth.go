package main

import "github.com/pcasteran/terraform-graph-beautifier/tfgraph"

func ApplyDepth(graph *tfgraph.Graph, depth *uint) {
	if depth == nil || *depth == 0 {
		return
	}

	traverse(graph, graph.Root, *depth-1)
}

func traverse(graph *tfgraph.Graph, module *tfgraph.Module, depth uint) {
	culls := []string{}

	for name, child := range module.Children {
		if childModule, ok := child.(*tfgraph.Module); ok {
			if depth > 0 {
				traverse(graph, childModule, depth-1)
			} else {
				culls = append(culls, name)
			}
		}
	}

	for _, cull := range culls {
		module.Children[cull] = tfgraph.NewBaseConfigElement(module, cull, "module")

		cullDependencies(graph, cull)
	}
}

func cullDependencies(graph *tfgraph.Graph, module string) {
	for i := 0; i < len(graph.Dependencies); i++ {
		dependency := graph.Dependencies[i]

		if isChildOf(dependency.Source, module) || isChildOf(dependency.Destination, module) {
			graph.Dependencies = append(graph.Dependencies[:i], graph.Dependencies[i+1:]...)
			i--
		}
	}
}

func isChildOf(child tfgraph.ConfigElement, parentName string) bool {
	parent := child.GetParent()

	if parent == nil {
		return false
	}

	return parent.GetName() == parentName || isChildOf(parent, parentName)
}
