package animation

import (
	"io"
	"log"
	"net/http"
)

type AnimationServer struct {
	Animation *Animation
	buffer    []byte
}

func NewAnimationServer(animation *Animation, bufferSize uint) *AnimationServer {
	return &AnimationServer{Animation: animation, buffer: make([]byte, bufferSize)}
}

func (as *AnimationServer) ServeHandler(w http.ResponseWriter, r *http.Request) {
	if n, err := as.Animation.Read(as.buffer); err != nil && err != io.EOF {
		log.Print("ERROR: ", err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Write(as.buffer[:n])
	}
}

func (as *AnimationServer) ClearHandler(w http.ResponseWriter, r *http.Request) {
	as.Animation.bd.Buffer.Reset()
	w.WriteHeader(http.StatusOK)
}
