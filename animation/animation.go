package animation

import (
	"github.com/enolgor/go-gfxino-codec/display"
)

type Animation struct {
	bd     *display.BufferedDisplay
	scenes []func(d display.Display)
	c      int
}

func NewAnimation(bd *display.BufferedDisplay) *Animation {
	return &Animation{bd: bd, scenes: []func(d display.Display){}, c: -1}
}

func (a *Animation) AddScene(scene func(d display.Display)) {
	a.scenes = append(a.scenes, scene)
}

func (a *Animation) SceneCount() int {
	return len(a.scenes)
}

func (a *Animation) Read(p []byte) (int, error) {
	if a.bd.Buffer.Len() == 0 {
		sc := a.nextScene()
		if sc != nil {
			(*sc)(a.bd)
		}
	}
	return a.bd.Read(p)
}

func (a *Animation) nextScene() *func(d display.Display) {
	if len(a.scenes) == 0 {
		return nil
	}
	a.c++
	if a.c >= len(a.scenes) {
		a.c = 0
	}
	return &a.scenes[a.c]
}
