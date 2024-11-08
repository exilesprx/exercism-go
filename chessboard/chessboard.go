package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	f, exists := cb[file]
	if !exists {
		return 0
	}

	sum := 0
	for rank := range f {
		if occupied := f[rank]; occupied {
			sum += 1
		}
	}

	return sum
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	if rank < 1 || rank > 8 {
		return 0
	}

	sum := 0
	for _, file := range cb {
		if occupied := file[rank-1]; occupied {
			sum += 1
		}
	}

	return sum
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	squares := 0
	for _, file := range cb {
		squares += len(file)
	}

	return squares
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	sum := 0
	for _, file := range cb {
		for _, occupied := range file {
			if occupied {
				sum += 1
			}
		}
	}

	return sum
}
