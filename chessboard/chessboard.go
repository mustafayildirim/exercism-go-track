package chessboard

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank = []bool

// Declare a type named Chessboard contains a map of eight Ranks, accessed with values from "A" to "H"
type Chessboard = map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
	rankItems := cb[rank]
	total := 0
	for _, val := range rankItems {
		if val {
			total++
		}
	}

	return total
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {

	if !(file <= 8 && file >= 1) {
		return 0
	}

	total := 0
	for _, rank := range cb {

		for index, val := range rank {
			if val && index+1 == file {
				total++
			}
		}

	}
	return total

}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	total := 0
	for _, rank := range cb {
		for range rank {
			total++
		}

	}
	return total
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	total := 0
	for _, rank := range cb {
		for _, val := range rank {
			if val {
				total++
			}
		}

	}
	return total

}
