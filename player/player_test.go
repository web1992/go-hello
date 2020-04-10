package player

import "testing"

func TestTryOut(t *testing.T) {

	var tapRecord = TapRecorder("TapRecord")
	var tapPlayer = TapPlayer("TapPlayer")

	TryOut(tapRecord)
	TryOut(tapPlayer)
}
