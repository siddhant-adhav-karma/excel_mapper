package main

type Transform interface {
	Transform()
}

type Sum struct {
	Columns []string
}

type Concat struct {
	Columns []string
	Separator string
}

type Average struct {
	Columns []string
}

type Unique struct {
	Column string
}

type Count struct {
}

func (s *Sum) Transform() {
}

func (s *Concat) Transform() {
}

func (s *Average) Transform(){
}

func (s *Unique) Transform() {
}

func (s *Count) Transform() {
}
