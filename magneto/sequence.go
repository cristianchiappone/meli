package main

type sequence struct {
	DNA    []string `json:"dna,omitempty"`
	RESULT string   `json:"is_mutant,omitempty"`
}