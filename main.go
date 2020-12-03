// This example simply captures data from your default microphone until you press Enter, after which it plays back the captured audio.
package main

import (
	"os"

	"github.com/gordonklaus/portaudio"
)

func main() {
	nSamples := 0
	sig := make(chan os.Signal, 1)
	portaudio.Initialize()
	defer portaudio.Terminate()
	in := make([]int32, 64)
	stream, err := portaudio.OpenDefaultStream(1, 0, 44100, len(in), in)
	chk(err)
	defer stream.Close()

	chk(stream.Start())
	for {
		chk(stream.Read())
		//chk(binary.Write(f, binary.BigEndian, in))
		nSamples += len(in)
		select {
		case <-sig:
			return
		default:
		}
	}
	chk(stream.Stop())
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
