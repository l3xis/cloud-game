// credit to https://github.com/fogleman/nes
package ui

import (
	"image"
	// "log"
	"runtime"

	// "github.com/gordonklaus/portaudio"
	"github.com/giongto35/game-online/webrtc"
)

const (
	width  = 256
	height = 240
	scale  = 3
	title  = "NES"
)

func init() {
	// we need to keep OpenGL calls on a single thread
	runtime.LockOSThread()
}

func Run(paths []string, imageChannel chan *image.RGBA, inputChannel chan int, webRTC *webrtc.WebRTC) {
	// TODO: stream audio
	// initialize audio
	// portaudio.Initialize()
	// defer portaudio.Terminate()

	// audio := NewAudio()
	// if err := audio.Start(); err != nil {
	// 	log.Fatalln(err)
	// }
	// defer audio.Stop()

	// run director
	director := NewDirector(imageChannel, inputChannel, webRTC)
	// director := NewDirector(audio, imageChannel, inputChannel)
	director.Start(paths)
}
