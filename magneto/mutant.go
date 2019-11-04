package main

func isMutant(dna []string) bool {
    N:= len(dna)
    
    if(len(dna) < 4){
        return false
    }

	lsec := [2][2]int{} //ultimo elemento de secuencia mutante
    count_mutant := 0 //cantida de secuencias mutantes
    
	//VERTICAL
	secV := 0
	for i := 0; i < (N - 3); i++ {
		for j := 0; j < N; j++ {
			if dna[i][j] == dna[i+3][j] && dna[i][j] == dna[i+2][j] && dna[i][j] == dna[i+1][j] {
				if secV > 0 {
                    if  (i != lsec[secV-1][0] &&
                        i+1 != lsec[secV-1][0] &&
                        i+2 != lsec[secV-1][0] &&
                        i+3 != lsec[secV-1][0] ||
                        j != lsec[secV-1][1]) {
						    count_mutant++
					}
				} else {
					lsec[secV][0] = i + 3
					lsec[secV][1] = j
					count_mutant++
					secV++
				}
				if count_mutant >= 2 {
					return true
				}
			}
		}
	}

	//HORIZONTAL
	secH := 0
	for i := 0; i < N; i++ {
		for j := 0; j < (N - 3); j++ {
			if dna[i][j] == dna[i][j+3] && dna[i][j] == dna[i][j+2] && dna[i][j] == dna[i][j+1] {
				if secH > 0 {
                    if  ((i != lsec[secH-1][0]) ||
                        (j != lsec[secH-1][1]) &&
                        (j+1 != lsec[secH-1][1]) &&
                        (j+2 != lsec[secH-1][1]) &&
                        (j+3 != lsec[secH-1][1])) {
					    	count_mutant++
					}
				} else {
					lsec[secH][0] = i
					lsec[secH][1] = j + 3
					count_mutant++
					secH++
				}
				if count_mutant >= 2 {
					return true
				}
			}
		}
	}

	//OBLICUO DERECHA-IZQUIERDA
	secO := 0
	for i := 0; i < (N - 3); i++ {
		for j := 0; j < (N - 3); j++ {
			if dna[i][j] == dna[i+3][j+3] && dna[i][j] == dna[i+2][j+2] && dna[i][j] == dna[i+1][j+1] {
				if secO > 0 {
                    if  ((i != lsec[secO-1][0] || j != lsec[secO-1][1]) &&
                        (i+1 != lsec[secO-1][0] || j+1 != lsec[secO-1][1]) &&
                        (i+2 != lsec[secO-1][0] || j+2 != lsec[secO-1][1]) &&
                        (i+3 != lsec[secO-1][0] || j+3 != lsec[secO-1][1])) {
						    count_mutant++
					}
				} else {
					lsec[secO][0] = i + 3
					lsec[secO][1] = j + 3
					count_mutant++
					secO++
				}
				if count_mutant >= 2 {
					return true
				}
			}
		}
	}

	//OBLICUO IZQUIERDA-DERECHA
	secCO := 0
	for i := 0; i < (N - 3); i++ {
		for j := 0; j < (N - 3); j++ {
			if dna[j][N-1-i] == dna[j+3][N-4-i] && dna[j][N-1-i] == dna[j+2][N-3-i] && dna[j][N-1-i] == dna[j+1][N-2-i] {
				if secCO > 0 {
					if (j != lsec[secCO-1][0] || (N-1-i) != lsec[secCO-1][1]) &&
						(j+3 != lsec[secCO-1][0] || (N-4-i) != lsec[secCO-1][1]) &&
						(j+2 != lsec[secCO-1][0] || (N-3-i) != lsec[secCO-1][1]) &&
						(j+1 != lsec[secCO-1][0] || (N-2-i) != lsec[secCO-1][1]) {
						    count_mutant++
					}
				} else {
					lsec[secCO][0] = j + 3
					lsec[secCO][1] = (N - 4 - i)
					count_mutant++
					secCO++
				}
				if count_mutant >= 2 {
					return true
				}
			}
		}
	}
	return false
}
