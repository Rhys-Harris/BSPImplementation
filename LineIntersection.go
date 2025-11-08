package main

type LineIntersection byte

const (
	LI_NONE LineIntersection = iota
	LI_INTERSECT
	LI_COINCIDENT
)
