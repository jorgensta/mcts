package mcts

// NimGameAction - NimAction struct
type NimGameAction struct {
	pieces int8
	value  int8
}

// ApplyTo - NimGameAction implementation of ApplyTo method of Action interface
func (a NimGameAction) ApplyTo(s GameState) GameState {
	NimGameState := s.(NimGameState)
	NimGameState.board = copyNimBoard(NimGameState.board)
	if NimGameState.nextToMove != a.value {
		panic("*hands slapped*,  not your turn")
	}

	if NimGameState.board < 0 {
		panic("*hands slapped*,  action illegal - square already occupied")
	}

	NimGameState.board -= a.pieces
	NimGameState.nextToMove *= -1
	return NimGameState
}

// NimGameState struct
type NimGameState struct {
	nextToMove int8
	board      int8
	ended      bool
	result     GameResult
}

// CreateNimGameState - initializes Nim game state
func CreateNimGameState(pileSize int8, nextToMove int8) NimGameState {
	board := initNewNimBoard(pileSize)
	state := NimGameState{nextToMove: nextToMove, board: board}
	return state
}

// IsGameEnded - Nim implementation of IsGameEnded method of GameState interface
func (n NimGameState) IsGameEnded() bool {
	_, ended := n.EvaluateGame()
	return ended
}

// EvaluateGame - Nim implementation of EvaluateGame method of GameState interface
func (n NimGameState) EvaluateGame() (result GameResult, ended bool) {

	defer func() {
		n.result = result
		n.ended = ended
	}()

	if n.ended {
		return n.result, n.ended
	}

	if n.board == 0 && n.nextToMove == 1 {
		return GameResult(-1), true
	}

	if n.board == 0 && n.nextToMove == -1 {
		return GameResult(1), true
	}

	return GameResult(0), false
}

// GetLegalActions - NimGameState implementation of GetLegalActions method of GameState interface
func (n NimGameState) GetLegalActions() []Action {
	var actions []Action

	maxPiecesToRemove := 3

	for i := 1; i < int(n.board)+1; i++ {
		if i <= maxPiecesToRemove {
			actions = append(actions, NimGameAction{pieces: int8(i), value: n.nextToMove})
			continue
		}
		break
	}

	return actions
}

// NextToMove - NimGameState implementation of NextToMove method of GameState interface
func (n NimGameState) NextToMove() int8 {
	return n.nextToMove
}

// GetBoard - returns board for Nim
func (n NimGameState) GetBoard() int8 {
	return n.board
}
