package models

type ExecutionGraph struct {
	DependencyHierarchy [][]BuilderMeta
}
