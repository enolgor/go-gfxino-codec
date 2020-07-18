package commands

const (
	UINT8 uint8 = iota
	UINT16
	UINT
	COLOR
	COLOR_SKIPPABLE
)

var CommandArgsMap map[uint8][]uint8 = map[uint8][]uint8{
	SETBITSIZE8:   {},
	SETROTATEON:   {},
	SETROTATEOFF:  {},
	SETFLIPON:     {},
	SETFLIPOFF:    {},
	SETBRIGHTNESS: {UINT8},
	DISPLAY:       {},
	DELAY:         {UINT8},
	CLEARDISPLAY:  {},
	SETCOLOR:      {COLOR},
	CLEARCOLOR:    {},
	FILLSCREEN:    {COLOR},
	DRAWPIXEL:     {UINT, UINT, COLOR_SKIPPABLE},
	DRAWLINE:      {UINT, UINT, UINT, UINT, COLOR_SKIPPABLE},
	DRAWFASTVLINE: {UINT, UINT, UINT, COLOR_SKIPPABLE},
	DRAWFASTHLINE: {UINT, UINT, UINT, COLOR_SKIPPABLE},
	DRAWRECT:      {UINT, UINT, UINT, UINT, COLOR_SKIPPABLE},
	FILLRECT:      {UINT, UINT, UINT, UINT, COLOR_SKIPPABLE},
	DRAWCIRCLE:    {UINT, UINT, UINT, COLOR_SKIPPABLE},
	FILLCIRCLE:    {UINT, UINT, UINT, COLOR_SKIPPABLE},
	DRAWROUNDRECT: {UINT, UINT, UINT, UINT, UINT, COLOR_SKIPPABLE},
	FILLROUNDRECT: {UINT, UINT, UINT, UINT, UINT, COLOR_SKIPPABLE},
	DRAWTRIANGLE:  {UINT, UINT, UINT, UINT, UINT, UINT, COLOR_SKIPPABLE},
	FILLTRIANGLE:  {UINT, UINT, UINT, UINT, UINT, UINT, COLOR_SKIPPABLE},
}

func GetArgSize(cmd uint8, mode8bit bool, skipColor bool) int {
	s := 0
	for _, a := range CommandArgsMap[cmd] {
		switch a {
		case UINT8:
			s++
		case UINT16:
			s += 2
		case UINT:
			if mode8bit {
				s++
			} else {
				s += 2
			}
		case COLOR:
			s += 2
		case COLOR_SKIPPABLE:
			if skipColor {
				s += 0
			} else {
				s += 2
			}

		}
	}
	return s
}
