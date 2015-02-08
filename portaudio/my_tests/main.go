package main

/*
import (
	"code.google.com/p/portaudio-go/portaudio"
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
	p := portaudio.LowLatencyParameters(h.DefaultInputDevice, h.DefaultOutputDevice)
	p.Input.Channels = 1
	p.Output.Channels = 1
	e := &echo{buffer: make([]float32, int(p.SampleRate*delay.Seconds()))}
	e.Stream, err = portaudio.OpenStream(p, e.processAudio)
	chk(err)
	return e
}

func (e *echo) processAudio(in, out []float32) {
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
*/

import (
	"code.google.com/p/portaudio-go/portaudio"
	"fmt"
	"math"
	"time"
)

const sampleRate = 16000 //44100

func main() {
	portaudio.Initialize()
	defer portaudio.Terminate()

	s, err := newStereoSine(500, 1500, sampleRate)
	if err != nil {
		fmt.Println("newStereoSine error")
		return
	}

	defer s.Close()
	chk(s.Start())
	time.Sleep(5 * time.Second)
	chk(s.Stop())
}

type stereoSine struct {
	*portaudio.Stream
	stepL, phaseL float64
	stepR, phaseR float64
}

func newStereoSine(freqL, freqR, sampleRate float64) (ss *stereoSine, err error) {

	ss = &stereoSine{nil, freqL / sampleRate, 0, freqR / sampleRate, 0}

	var di *portaudio.DeviceInfo

	di, err = portaudio.DefaultOutputDevice()
	if err != nil {
		return nil, err
	}

	//fmt.Println("DefaultHighOutputLatency= ", di.DefaultHighOutputLatency)
	//fmt.Println("DefaultLowOutputLatency= ", di.DefaultLowOutputLatency)

	var p = portaudio.StreamParameters{
		Output: portaudio.StreamDeviceParameters{
			Device:   di,
			Channels: 2,
			Latency:  di.DefaultHighOutputLatency,
		},
		SampleRate:      sampleRate,
		FramesPerBuffer: portaudio.FramesPerBufferUnspecified,
	}

	ss.Stream, err = portaudio.OpenStream(p, ss.processAudio)
	if err != nil {
		return nil, err
	}

	//ss.Stream, err = portaudio.OpenStream(portaudio.HighLatencyParameters(nil, odi), ss.processAudio)
	//s.Stream, err = portaudio.OpenDefaultStream(0, 2, sampleRate, 0, s.processAudio)

	return ss, nil
}

func (g *stereoSine) processAudio(out [][]float32) {
	for i := range out[0] {
		out[0][i] = float32(math.Sin(2 * math.Pi * g.phaseL))
		_, g.phaseL = math.Modf(g.phaseL + g.stepL)
		out[1][i] = float32(math.Sin(2 * math.Pi * g.phaseR))
		_, g.phaseR = math.Modf(g.phaseR + g.stepR)
	}
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}

/*

/*
import (
	"code.google.com/p/portaudio-go/portaudio"
	"encoding/binary"
	"fmt"
	"os"
	"os/signal"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("missing required argument:  output file name")
		return
	}
	fmt.Println("Recording.  Press Ctrl-C to stop.")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)

	fileName := os.Args[1]
	if !strings.HasSuffix(fileName, ".aiff") {
		fileName += ".aiff"
	}
	//fmt.Println(fileName)

	f, err := os.Create(fileName)
	chk(err)

	// form chunk
	_, err = f.WriteString("FORM")
	chk(err)
	chk(binary.Write(f, binary.BigEndian, int32(0))) //total bytes
	_, err = f.WriteString("AIFF")
	chk(err)

	// common chunk
	_, err = f.WriteString("COMM")
	chk(err)
	chk(binary.Write(f, binary.BigEndian, int32(18)))                  //size
	chk(binary.Write(f, binary.BigEndian, int16(1)))                   //channels
	chk(binary.Write(f, binary.BigEndian, int32(0)))                   //number of samples
	chk(binary.Write(f, binary.BigEndian, int16(32)))                  //bits per sample
	_, err = f.Write([]byte{0x40, 0x0e, 0xac, 0x44, 0, 0, 0, 0, 0, 0}) //80-bit sample rate 44100
	chk(err)

	// sound chunk
	_, err = f.WriteString("SSND")
	chk(err)
	chk(binary.Write(f, binary.BigEndian, int32(0))) //size
	chk(binary.Write(f, binary.BigEndian, int32(0))) //offset
	chk(binary.Write(f, binary.BigEndian, int32(0))) //block
	nSamples := 0
	defer func() {
		// fill in missing sizes
		totalBytes := 4 + 8 + 18 + 8 + 8 + 4*nSamples
		_, err = f.Seek(4, 0)
		chk(err)
		chk(binary.Write(f, binary.BigEndian, int32(totalBytes)))
		_, err = f.Seek(22, 0)
		chk(err)
		chk(binary.Write(f, binary.BigEndian, int32(nSamples)))
		_, err = f.Seek(42, 0)
		chk(err)
		chk(binary.Write(f, binary.BigEndian, int32(4*nSamples+8)))
		chk(f.Close())
	}()

	portaudio.Initialize()
	defer portaudio.Terminate()

	stream, err := portaudio.OpenDefaultStream(1, 0, 44100, len(in), in)
	chk(err)
	defer stream.Close()

	chk(stream.Start())

	for {
		chk(stream.Read())
		chk(binary.Write(f, binary.BigEndian, in))
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
*/
