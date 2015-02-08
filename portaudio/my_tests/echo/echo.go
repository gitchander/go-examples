package main

import (
	"code.google.com/p/portaudio-go/portaudio"
	"fmt"
	"time"
)

func main() {
	portaudio.Initialize()
	defer portaudio.Terminate()
	e := newEcho(time.Second / 3)
	defer e.Close()
	chk(e.Start())
	time.Sleep(60 * time.Second)
	chk(e.Stop())
}

type echo struct {
	*portaudio.Stream
	buffer []float32
	i      int
}

func newEcho(delay time.Duration) *echo {
	h, err := portaudio.DefaultHostApi()
	chk(err)
	p := portaudio.HighLatencyParameters(h.DefaultInputDevice, h.DefaultOutputDevice)
	p.Input.Channels = 1
	p.Output.Channels = 1

	n := int(p.SampleRate * delay.Seconds())

	//fmt.Println("n= ", n)

	e := &echo{buffer: make([]float32, n)}
	e.Stream, err = portaudio.OpenStream(p, e.processAudio)
	chk(err)
	return e
}

func (e *echo) processAudio(in, out []float32) {

	if len(in) != len(out) {
		fmt.Printf("len(in)= %d, len(out)= %d\n", len(in), len(out))
	}

	for i := range out {
		out[i] = .7 * e.buffer[e.i]
		e.buffer[e.i] = in[i]
		e.i = (e.i + 1) % len(e.buffer)
	}
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}

//----------------------------------------------------------
/*
type SampleBuffer struct {
	data []float32
	head int
	tail int
}

func NewSampleBuffer(size int) *SampleBuffer {

	if size <= 0 {
		return nil
	}

	var sb = &SampleBuffer{
		data: make([]float32, size+1),
		head: 0,
		tail: 0,
	}

	return sb
}

func (this *SampleBuffer) Read(data []float32) (n int, err error) {

	return
}

func (this *SampleBuffer) Write(data []float32) (n int, err error) {

	return
}
*/
//----------------------------------------------------------
