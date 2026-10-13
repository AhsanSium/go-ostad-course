package main

import "fmt"

type TapePlayer struct {}

func (t TapePlayer) Play (song string) {
	fmt.Println("Playing Song: ", song)
}
func (t TapePlayer) Stop () {
	fmt.Println("Stopping Tape Player ~~~~")
}

func PlayList (device Playable, songs []string) {
	for _, song := range songs {
		device.Play(song)
		// device.record(song) // cannot use device (variable of type Playable) as TapeRecorder value in argument to device.record
		// device.(TapeRecorder).Record(song) // type assertion to access Record method
		if recorder, ok := device.(TapeRecorder); ok { // type assertion to access Record method
			recorder.Record(song)
		}
	}
	device.Stop()
}

type TapeRecorder struct {}

func (t TapeRecorder) Play (song string) {
	fmt.Println("Playing Song: ", song)
}

func (t TapeRecorder) Stop() {
	fmt.Println("Stopping Tape Recorder ~~~~")
}

func (t TapeRecorder) Record (song string) {
	fmt.Println("Recording Song: ", song)
}

type Playable interface {
	Play(song string)
	Stop()
}

func main () {
	fmt.Println("Hello Go")
	songs := []string{"Song 1", "Song 2", "Song 3"}
	fmt.Println(songs)

	tapePlayer := TapePlayer{}

	PlayList(tapePlayer, songs)

	recoreder := TapeRecorder{}

	PlayList(recoreder, songs) // cannot use recoreder (variable of struct type TapeRecorder) as TapePlayer value in argument to PlayList
	/*
	Need to create new func PlayListRecorder and all the Play Stop func
	*/

}