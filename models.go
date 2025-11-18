package main

type Mapping struct {
	Source    []string
	Target    string
	Transform string
	Params    map[string]any
}

type File struct {
}
