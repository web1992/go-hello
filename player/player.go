package player

import "fmt"

type Player interface {
	Play(string)
	Stop()
}

func TryOut(player Player) {
	player.Play("Test Track")
	player.Stop()

	// type
	tr, ok := player.(TapRecorder)
	if ok {
		tr.Record()
	}
}

type TapRecorder string

func (t TapRecorder) Play(s string) {
	fmt.Println("HaHaHa...", t)
}

func (t TapRecorder) Stop() {
	fmt.Println("Stop...", t)
}

func (t TapRecorder) Record() {
	fmt.Println("Record...")
}

type TapPlayer string

func (t TapPlayer) Play(s string) {
	fmt.Println("HaHaHa...", t)
}

func (t TapPlayer) Stop() {
	fmt.Println("Stop...", t)
}
