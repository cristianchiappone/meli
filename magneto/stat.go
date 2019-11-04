package main

type stat struct {
	CountMutantDna int32   `json:"count_mutant_dna,omitempty"`
	CountHumanDna  int32   `json:"count_human_dna,omitempty"`
	Ratio          float64 `json:"count_ratio_dna,omitempty"`
}