package models

type DataFlow struct {
	Name       string
	TargetData string
	ExecGraph  ExecutionGraph
	// TODO: implement Transients
	// TODO: implement ResolutionSpec
}
